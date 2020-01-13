# libp2p issue labelling taxonomy

> r0, 2020-01-04; author: @raulk

**Table of contents**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Motivation](#motivation)
- [Descriptive labels](#descriptive-labels)
  - [Functional area (`area/`)](#functional-area-area)
  - [Component (`comp/`)](#component-comp)
  - [Difficulty (`dif/`)](#difficulty-dif)
  - [Size (`size/`)](#size-size)
  - [Kind (`kind/`)](#kind-kind)
  - [Impact (`impact/`)](#impact-impact)
  - [Topic (`topic/`)](#topic-topic)
- [Execution labels](#execution-labels)
  - [Priority (`Pn`)](#priority-pn)
  - [Hints/advisories (`hint/`)](#hintsadvisories-hint)
  - [Status (`status/`)](#status-status)
- [Implementation plan](#implementation-plan)
- [Inspiration](#inspiration)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Motivation

This document describes the issue/ticket labelling taxonomy that is being
adopted in the libp2p project.

This taxonomy will be initially rolled out within the domains of go-libp2p,
where it'll be trialled, adjusted and iterated upon. We hope it'll enhance
navigability, searchability, discoverability and accessibility of issues,
features, demands, requests, discussions, etc. on a path towards better
organisation, structure, and, ultimately, accountability within, between and
towards the community and implementation teams.

_"Why is this important?"_, you may ask. A taxonomy is the basic building
block that makes it possible to create workflows for triaging, discussing,
development, and ultimately, delivery. In fact, issue classification is only the
start of a set of methods and tools we desire to ship to help facilitate
ecosystem alignment, foster collaboration, bolster transparency, streamline
planning, and ultimately, reduce friction to keep the libp2p community engaged,
vibrant, and excited about what's to come.

We are fans of iterative experimentation. In the next weeks/months, as we roll
this out and adjust it, we will gather feedback and input from stakeholders
across the ecosystem.

Our endgame is to converge on a system that pleases all language implementers,
so that we can normalise this taxonomy across all language repos, and plan
features and changes in total inter-alignment. We are convinced that by creating
a cohesive and seamless experience across libp2p implementations, we'll unlock
tremendous value for our beloved downstream users and developers.

## Descriptive labels

These labels describe some aspect of the issue itself. They can generally be set
when the issue is first filed, or as discoveries are made. We will adopt tools
that allow submitters to set these labels.

### Functional area (`area/`)

> Defines the functional area of the libp2p stack an
> issue, report, question, etc. addresses. These labels refer to abstraction
> layers; to contextualise the issue to a specific implementation, also add a
> "component" label.

* area/conn-manager
* area/connection-bootstrapping
* area/nat-traversal
* area/peer-routing
* area/content-routing
* area/discovery
* area/identity
* area/multiplexing
* area/negotiation
* area/peerstore
* area/pubsub
* area/relay
* area/security-channel
* area/transport
* area/universal

### Component (`comp/`)

> Specifies the actual component (implementation) an issue refers to. Used when
> it's clear that an issue affects one or many specific components.

* comp/floodsub
* comp/gossipsub
* comp/kad-dht
* comp/noise
* comp/quic
* comp/secio
* comp/tcp
* comp/tls
* comp/webrtc
* comp/wss
* comp/yamux

### Difficulty (`dif/`)

> Subjectively estimates the difficulty of resolution of an issue, based on
> prior expertise about libp2p and its constructs, p2p networking, systems
> programming, etc.

Values and calibration guidance:

* dif/trivial: Can be confidently tackled by newcomers, who are widely
  unfamiliar with libp2p.
* dif/easy: A prexisting libp2p user should be able to pick this up.
* dif/medium: Prior experience in having developed libp2p modules is likely
  helpful.
* dif/hard: Suggests that having worked on the specific component affected
  by this issue is important.
* dif/expert: Requires extensive knowledge of the history, implications,
  ramifications of the issue, as well as deep understanding of the libp2p stack
  and experience with systems and network programming.

### Size (`size/`)

> Estimates the amount of work. This axis is distinct to the complexity, e.g.
> something can be easy but require a lot of work, or vice versa. 

Values and calibration guidance:

* size/XS: minutes.
* size/S: several hours.
* size/M: one day.
* size/L: 2-3 days.
* size/XL: 1 week.
* size/XXL: > 1 week.

### Kind (`kind/`)

> Describes the nature of the issue.

* kind/bug
* kind/improvement
* kind/tracking-issue
* kind/question
* kind/enhancement
* kind/architecture
* kind/discussion

### Impact (`impact/`)

> Contextualises the impact, and therefore the perceived/derived severity, of
> the issue.

* impact/regression
* impact/api-breakage
* impact/quality
* impact/dx
* impact/test-flakiness

### Topic (`topic/`)

> Classifies issues pertaining to particular themes. Topics can be created to
> refer to epics, larger bodies of work, or ad-hoc endeavours.

* topic/interoperability
* topic/docs
* topic/infra

## Execution labels

These labels describe some aspect of the execution on the issue. These labels
are dynamic; they change as a result of planning, analysis, scoping, and
generally applying the development workflow. They are usually set by
maintainers.

### Priority (`Pn`)

> Indicates the current implementation priority; used for planning and
> roadmapping purposes.

* P0: Critical.
* P1: High.
* P2: Medium.
* P3: Low.

### Hints/advisories (`hint/`)

> Surfaces observations and facts about issues, facilitating progress to be
> made.

* hint/good-first-issue
* hint/needs-contributor
* hint/needs-participation
* hint/needs-decision
* hint/needs-triage
* hint/needs-analysis
* hint/needs-author-input
* hint/needs-team-input
* hint/needs-community-input
* hint/needs-review
* hint/needs-help

### Status (`status/`)

> Marks the current status of the issue. 

* status/done
* status/deferred
* status/in-progress
* status/blocked
* status/inactive
* status/waiting
* status/rotten

## Implementation plan

1. Create labels on go-libp2p
   [declaratively](https://github.com/micnncim/action-label-syncer).
2. Write script to bulk transfer issues from inner repos to top-level go-libp2p
   repo via the [`transferIssue` GraphQL mutation](https://developer.github.com/v4/mutation/transferissue/), attaching the appropriate area/component labels during the process.
3. Disable issue tracker in inner repos.
4. Triage all issues to assign the remaining labels.
5. Profit! We can finally start implementing workflows and using tools like
   ZenHub.

## Inspiration

* https://github.com/facebook/react/labels
* https://github.com/rust/rust-lang/labels
* https://github.com/kubernetes/kubernetes/labels
* https://bugs.python.org/
* https://github.com/tensorflow/tensorflow/labels
* https://github.com/rust-lang/triagebot/wiki/Labeling
