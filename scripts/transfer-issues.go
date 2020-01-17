package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	// names
	owner        = "libp2p"
	dhtRepo      = "go-libp2p-kad-dht"
	goLibp2pRepo = "go-libp2p"

	// id for go-libp2p
	goLibp2pRepoId = "MDEwOlJlcG9zaXRvcnk0MzQ2NTY3NQ=="

	// place holders
	cursorPlaceHoler = "{LABEL_CURSOR}"
)

var repoNames = []string{dhtRepo}

var client = graphql.NewClient("https://api.github.com/graphql")

// labels we need to apply to an issue belonging to a particular repo
var labelsMap = map[string][]string{
	dhtRepo: []string{"area/content-routing", "comp/kad-dht", "hint/needs-label-validation"},
}

func main() {
	// fetch all label Ids for the main repo
	labelIdMap, err := fetchLabelIds()
	if err != nil {
		fmt.Printf("failed to fetch label Ids, err=%s", err)
		os.Exit(1)
	}
	fmt.Printf("\n finished fetching labelIds for go-libp2p, total labels=%d", len(labelIdMap))

	// for each repo
	for _, repoName := range repoNames {
		// fetch all issueIds
		issueIds, err := fetchAllIssues(repoName)
		if err != nil {
			fmt.Printf("failed to fetch issueIds for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}
		fmt.Printf("\n finished processing issueIds for repoName %s, total issueIds in repoName=%d", repoName, len(issueIds))

		// resolve label names to label Ids
		labelIds, err := resolveLabelNamesToIds(repoName, labelIdMap)
		if err != nil {
			fmt.Printf("failed to resolve labelIds for repo %s, err=%s", repoName, err)
		}

		// transfer all issueIds & apply labels
		if err := transferIssuesAndApplyLabels(repoName, issueIds, labelIds); err != nil {
			fmt.Printf("failed to transfer issue/apply label for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}
		fmt.Printf("\n finished transferring issues & applying labels for repo %s", repoName)

		// create re-direct issue & pin it
	}

}

func resolveLabelNamesToIds(repoName string, labelIdMap map[string]string) ([]string, error) {
	labelNames, ok := labelsMap[repoName]
	if !ok {
		return nil, errors.New("labels not configured")
	}

	var labelIds []string
	for _, l := range labelNames {
		labelId, ok := labelIdMap[l]
		if !ok {
			return nil, errors.New(fmt.Sprintf("labelId absent for label %s", l))
		}
		labelIds = append(labelIds, labelId)
	}
	return labelIds, nil
}

func executeQuery(query string, resp interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", "Bearer TOKEN HERE")
	return client.Run(context.Background(), req, resp)
}

func fetchLabelIds() (map[string]string, error) {
	labelsToIds := make(map[string]string)

	queryString := fmt.Sprintf(`
	    {
	  repository(owner: "%s", name: "%s") {
	    labels(first: 100%s) {
	      edges {
	        cursor
	        node{
	          name
	          id
	        }
	      }

	    }
	  }
	}`, owner, goLibp2pRepo, cursorPlaceHoler)

	type LabelData struct {
		Repository struct {
			Labels struct {
				Edges []struct {
					Cursor string
					Node   struct {
						Name string
						Id   string
					}
				}
			}
		}
	}
	var labels LabelData
	cursor := ""
	for {
		// set cursor for pagination
		q := strings.Replace(queryString, cursorPlaceHoler, cursor, -1)
		if err := executeQuery(q, &labels); err != nil {
			return nil, err
		}

		if len(labels.Repository.Labels.Edges) == 0 {
			return labelsToIds, nil
		}

		for _, l := range labels.Repository.Labels.Edges {
			labelsToIds[l.Node.Name] = l.Node.Id
		}

		// set cursor to fetch the next page
		cursor = fmt.Sprintf(", after: \"%s\"", labels.Repository.Labels.Edges[len(labels.Repository.Labels.Edges)-1].Cursor)
	}
}

func fetchRepoIds() (map[string]string, error) {
	repoQuery := `{
  repository(owner:"%s", name:"%s") {
    id
  }
}`
	type Data struct {
		Repository struct {
			Id string
		}
	}

	repoNamesToIds := make(map[string]string)
	for _, r := range repoNames {
		var resp Data
		if err := executeQuery(fmt.Sprintf(repoQuery, owner, r), &resp); err != nil {
			return nil, err
		}
		repoNamesToIds[r] = resp.Repository.Id
	}

	fmt.Printf("\n finished fetching repoIds")

	return repoNamesToIds, nil
}

func fetchAllIssues(repoName string) ([]string, error) {
	var issueIds []string

	type IssueData struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Cursor string
					Node   struct {
						Id string
					}
				}
			}
		}
	}

	queryString := fmt.Sprintf(`
	    {
	  repository(owner: "%s", name: "%s") {
	    issues(first: 100%s,states:OPEN) {
	      edges {
	        cursor
	        node{
	          id
	        }
	      }

	    }
	  }
	}`, owner, repoName, cursorPlaceHoler)

	cursor := ""
	var issues IssueData
	for {
		// set cursor for pagination & create the req
		q := strings.Replace(queryString, cursorPlaceHoler, cursor, -1)
		if err := executeQuery(q, &issues); err != nil {
			return nil, err
		}

		if len(issues.Repository.Issues.Edges) == 0 {
			return issueIds, nil
		}

		for _, l := range issues.Repository.Issues.Edges {
			issueIds = append(issueIds, l.Node.Id)
		}

		// set cursor to fetch the next page
		cursor = fmt.Sprintf(", after: \"%s\"", issues.Repository.Issues.Edges[len(issues.Repository.Issues.Edges)-1].Cursor)
	}
}

func transferIssuesAndApplyLabels(repoName string, issues []string, labelIds []string) error {
	transferQueryStr := `mutation {
		  transferIssue(input:{issueId:"%s",repositoryId:"%s"}) {
		    clientMutationId
		  }
		}`

	for _, i := range issues {
		// transfer issue
		newStr := fmt.Sprintf(transferQueryStr, i, goLibp2pRepoId)

		if err := executeQuery(newStr, nil); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to transfer issue %s", i))
		}

		if err := applyLabels(i, labelIds); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to apply label to issue %s", i))
		}
	}

	return nil
}

func applyLabels(issueId string, labelIds []string) error {
	queryString := `mutation {
		  addLabelsToLabelable(input:{labelableId:"%s",labelIds:[%s]}) {
		   clientMutationId
		  }
		}`

	labelStr := "\"" + labelIds[0] + "\""
	for _, l := range labelIds[1:] {
		labelStr = labelStr + ",\"" + l + "\""
	}

	str := fmt.Sprintf(queryString, issueId, labelStr)
	return executeQuery(str, nil)
}