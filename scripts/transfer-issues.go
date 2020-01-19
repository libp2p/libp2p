package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Edge struct {
	Cursor string
	Node   struct {
		Name string
		Id   string
	}
}

type Issue struct {
	Id string
}

type Config struct {
	Owner              string
	RedirectIssueTitle string
	RedirectIssueBody  string
	GoLibp2pRepo       string

	RepoToLabelsMap map[string][]string
}

const (
	// place holders
	cursorPlaceHoler  = "{LABEL_CURSOR}"
	bearerTokenEnvVar = "GITHUB_BEARER_TOKEN"
)

var (
	cfg         Config
	bearerToken string
	client      = graphql.NewClient("https://api.github.com/graphql")
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("\n failed to read config file; err=%s", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Print("\n failed to unmarshal config, err=%s", err)
		os.Exit(1)
	}

	bearerToken = os.Getenv(bearerTokenEnvVar)
	if len(bearerToken) == 0 {
		fmt.Printf("\n failed to read github bearer token from Environment variable %s", bearerTokenEnvVar)
		os.Exit(1)
	}
}

func main() {
	// fetch mainRepoId
	mainRepoId, err := fetchRepoId(cfg.GoLibp2pRepo)
	if err != nil {
		fmt.Printf("\n failed to fetch repoId for main repo %s", cfg.GoLibp2pRepo)
		os.Exit(1)
	}

	// fetch all repo Ids
	repoIdsMap, err := fetchRepoIds()
	if err != nil {
		fmt.Printf("\n failed to fetch repoIds, err=%s", err)
		os.Exit(1)
	}
	fmt.Printf("\n finished fetching repoIds, total Ids=%d", len(repoIdsMap))

	// fetch all label Ids for the main repo
	labelIdMap, err := fetchLabelIds(cfg.GoLibp2pRepo)
	if err != nil {
		fmt.Printf("\n failed to fetch label Ids for repo %s, err=%s", cfg.GoLibp2pRepo, err)
		os.Exit(1)
	}
	fmt.Printf("\n finished fetching labelIds for repo %s, total labels=%d", cfg.GoLibp2pRepo, len(labelIdMap))

	// for each repo
	for repoName, _ := range cfg.RepoToLabelsMap {
		fmt.Printf("\n\n ------------------------ Will start processing repo %s now -------------", repoName)

		// resolve label names to label Ids
		labelIds, err := resolveLabelNamesToIds(repoName, labelIdMap)
		if err != nil {
			fmt.Printf("\n failed to resolve label names to labelIds for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}

		// resolve reponame to repoId
		repoId, ok := repoIdsMap[repoName]
		if !ok {
			fmt.Printf("\n failed to resolve repoId for repo %s", repoName)
			os.Exit(1)
		}

		// fetch all issueIds
		issueIds, err := fetchAllIssues(repoName)
		if err != nil {
			fmt.Printf("\n failed to fetch issueIds for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}
		fmt.Printf("\n finished fetching issueIds for repoName %s, total issueIds=%d", repoName, len(issueIds))

		// transfer all issueIds to the main repo & apply labels
		if err := transferIssuesAndApplyLabels(mainRepoId, issueIds, labelIds); err != nil {
			fmt.Printf("\n failed to transfer issue/apply label for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}
		fmt.Printf("\n finished transferring issues & applying labels for repo %s", repoName)

		// create re-direct issue & pin it
		if err := createAndPinRedirectIssue(repoId); err != nil {
			fmt.Printf("failed to create & pin redirect issue for repo %s, err=%s", repoName, err)
			os.Exit(1)
		}
		fmt.Printf("\n finished creating & pinning re-direct issue on repo %s", repoName)

		fmt.Printf("\n ------------------------ Finished processing repo %s, Total Issues Transferred=%d -------------", repoName, len(issueIds))
	}
}

// resolves the 'textual' label names to the corresponding label Ids required by the Github Api
func resolveLabelNamesToIds(repoName string, labelIdMap map[string]string) ([]string, error) {
	labelNames, ok := cfg.RepoToLabelsMap[repoName]
	if !ok {
		return nil, errors.New("labels not configured for the repo")
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	// https://developer.github.com/v4/previews/#pinned-issues-preview
	req.Header.Set("Accept", "application/vnd.github.elektra-preview+json")
	return client.Run(context.Background(), req, resp)
}

// fetches all the label Ids for the labels created on the given repo
func fetchLabelIds(repoName string) (map[string]string, error) {
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
	}`, cfg.Owner, repoName, cursorPlaceHoler)

	type LabelData struct {
		Repository struct {
			Labels struct {
				Edges []Edge
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

// fetches the issue Ids for existing OPEN issues on the given repo
func fetchAllIssues(repoName string) ([]string, error) {
	var issueIds []string

	type IssueData struct {
		Repository struct {
			Issues struct {
				Edges []Edge
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
	}`, cfg.Owner, repoName, cursorPlaceHoler)

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

// transfer the given issues to the destination repo & apply the labels with the given labelIds to them
func transferIssuesAndApplyLabels(destinationRepoId string, issues []string, labelIds []string) error {
	transferQueryStr := `mutation {
		  transferIssue(input:{issueId:"%s",repositoryId:"%s"}) {
            issue {
				id
			}
		  }
		}`

	// Github changes the issueId after transferring it
	// So, we need to read the new label Id from the response before we can label it
	type TransferIssueResponse struct {
		TransferIssue struct {
			Issue Issue
		}
	}

	var resp TransferIssueResponse
	for _, i := range issues {
		// transfer issue
		newStr := fmt.Sprintf(transferQueryStr, i, destinationRepoId)
		if err := executeQuery(newStr, &resp); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to transfer issue %s", i))
		}

		// apply label
		if err := applyLabels(resp.TransferIssue.Issue.Id, labelIds); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to apply label to issue %s", i))
		}
	}

	return nil
}

// creates & pins the redirect issues on the sub repos to redirect users to the transferred issues on the main repo
func createAndPinRedirectIssue(repoId string) error {
	// create issue
	createIssueQuery := fmt.Sprintf(`mutation {
       createIssue(input:{repositoryId:,"%s", title:"%s", body:"%s"}) {
        issue {
		  id
        }
       }
    }
`, repoId, cfg.RedirectIssueTitle, cfg.RedirectIssueBody)

	type CreateIssueResponse struct {
		CreateIssue struct {
			Issue Issue
		}
	}
	var createResp CreateIssueResponse
	if err := executeQuery(createIssueQuery, &createResp); err != nil {
		return errors.Wrap(err, "failed to create re-direct issue")
	}

	// pin issue
	pinIssueQuery := fmt.Sprintf(`mutation{
 		 pinIssue(input:{issueId:"%s"}) {
    		clientMutationId
  		}
	}`, createResp.CreateIssue.Issue.Id)

	if err := executeQuery(pinIssueQuery, nil); err != nil {
		return errors.Wrap(err, "failed to pin issue")
	}
	return nil
}

// apply labels with the given label Ids to the issue
func applyLabels(issueId string, labelIds []string) error {
	queryString := `mutation {
		  addLabelsToLabelable(input:{labelableId:"%s",labelIds:[%s]}) {
		   clientMutationId
		  }
		}`

	// need to convert the label Ids to csv strings
	labelStr := "\"" + labelIds[0] + "\""
	for _, l := range labelIds[1:] {
		labelStr = labelStr + ",\"" + l + "\""
	}

	str := fmt.Sprintf(queryString, issueId, labelStr)
	return executeQuery(str, nil)
}

// fetches the Ids for all repositories so we can create issues in the repos using their Ids
func fetchRepoIds() (map[string]string, error) {
	repoNamesToIds := make(map[string]string)
	for r, _ := range cfg.RepoToLabelsMap {
		id, err := fetchRepoId(r)
		if err != nil {
			return nil, err
		}
		repoNamesToIds[r] = id
	}

	return repoNamesToIds, nil
}

func fetchRepoId(repoName string) (string, error) {
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

	var resp Data
	if err := executeQuery(fmt.Sprintf(repoQuery, cfg.Owner, repoName), &resp); err != nil {
		return "", err
	}
	return resp.Repository.Id, nil
}
