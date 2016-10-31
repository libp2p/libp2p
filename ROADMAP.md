# libp2p Roadmap

> We track the development of the libp2p project through Github issues and [Waffle.io](https://waffle.io/libp2p/libp2p). See our waffle board at: [https://waffle.io/libp2p/libp2p](https://waffle.io/libp2p/libp2p)

--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------

# 2016 Q3

## Milestone - JS and Go libp2p interop

> Summary: Achieve interoperability between the JS and Go implementations, specificially with secio.

- Tasks
  - [x] Migration to pull-streams (track: https://github.com/ipfs/js-ipfs/issues/403). Also part of js-ipfs Milestone 2
  - [x] js-secio - https://github.com/libp2p/js-libp2p-secio
  - [x] Interop with Go
    - [x] Verify that it interops
    - [ ] Automated tests(track: https://github.com/ipfs/js-libp2p-ipfs/issues/15)
  - [ ] Integrate the `ping protocol` in js-libp2p (track: https://github.com/ipfs/js-libp2p-ipfs/pull/13)
- Dependencies:
  - n/a
- Requirements by other projects
  - js-ipfs and go-ipfs interop

### Notes `NA`
### Expected date of completion: `Q3`

--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------

# 2016 Q4

## Milestone - libp2p interfaces

> Summary: Review and converge both go and js interfaces

### Leads

- David
- Juan

### Tasks

### Notes

### Expected date of completion: `Week 2 - October 24`

----------------------------------------------------------------



## Milestone - interface-xtp

> Summary: ship xTP interface

### Leads

- David
- Juan

### Tasks

### Notes

### Expected date of completion: `Week 4 - November 7`

----------------------------------------------------------------



## Milestone - Documentation of the libp2p modules

> Summary: Create easy to write documentation for libp2p modules that is searchable.

### Leads

- Friedel

### Tasks

### Notes

### Expected date of completion: `Week 5 - November 14`

----------------------------------------------------------------



## Milestone - Standardize Organization and libp2p/libp2p org clean up

> Summary:

### Leads

- Richard

### Tasks

- [ ] Standardize all READMEs and repositories in the libp2p organization
  - [ ] Use standard-readme for all READMEs
  - [ ] Look at project-repos.ipfs.io to ensure green across the board.
  - [ ] Point to IPFS community guidelines for the libp2p project

### Notes

Having the community guidelines pointed to, and creating custom ones as needed, should be the priority.

### Expected date of completion: `Week 5 - November 14`

----------------------------------------------------------------



## Milestone - libp2p.io

> Summary: We need to ship a first version of the libp2p.io website

### Leads

- @VictorBjelkholm

### Tasks

- [ ] Agree and create tasks about what's missing for first version
- [ ] Agree on where to take inspiration
- [ ] Come up with site structure (sitemap)
- [ ] Have content written and integrated
- [ ] Setup Libp2p examples
- [ ] Deploy and setup CD for shipping website

### Notes


The website should be launched, accessible, and thoughtful.

libp2p.io has it's repository here: https://github.com/libp2p/libp2p-website

### Expected date of completion: `Week 6 - November 21`

----------------------------------------------------------------



## Milestone - ship js-libp2p-dht

> Summary:

### Leads

- David

### Tasks

### Notes

### Expected date of completion: `Week 9 - December 12`

----------------------------------------------------------------



## Milestone - Write libp2p-relay spec

> Summary:

### Leads

- David
- Juan

### Tasks

### Notes

### Expected date of completion: `Week 12 - January 2`

----------------------------------------------------------------


## Milestone - Ship libp2p-relay

> Summary:

### Leads

- David
- Jeromy


### Tasks

### Notes

### Expected date of completion: `Week 13 - January 9`

----------------------------------------------------------------

## Milestone - Ship Private Networks

> Summary: have pre-shared key based private networks in go-libp2p and go-ipfs

### Leads

- Kuba

### Tasks
 - [x] Choose crypto primitive - XSalsa20 -  https://github.com/ipfs/notes/issues/177#issuecomment-255588927
 - [ ] Get review on the crypto
 - [x] Decide on format for key file in ipfs repo - https://github.com/ipfs/notes/issues/177#issue-183845331
 - [x] Write a spec/manual
 - [ ] Publish the spec/manual
 - [ ] Implementation:
   - [x] Create private networks interface - https://github.com/libp2p/go-libp2p-interface-pnet/
   - [ ] Finish private netwroks interface
   - [ ] Finish private networks implementation - https://github.com/libp2p/go-libp2p-pnet
   - [ ] Integrate private networks interface with go-libp2p-swarm and go-ipfs
 - [ ] Tests for private networks

### Notes

### Expected date of completion: `Week 7 - November 28`

---------


--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------

# CACHED MILESTONES (for later re-evaluation)

## Milestone 5 - Developer Experience

> Summary: Get libp2p into a state that is super easy to use and understand.

##@ Tasks

- [ ] Examples
  - [ ] How to use js-libp2p (dial, multi transport, stream muxing, identify, protocol multiplexing)
  - [ ] How to use go-libp2p (everything js-libp2p + DHT)
- [ ] Spec
  - [ ] Move spec to the libp2p org
  - [ ] Revisit and improve
- [ ] interface-libp2p
  - [ ] Make js-libp2p-ipfs and js-libp2p-ipfs-browser use it
  - [ ] Create js-libp2p, a base class that other libp2p builds can reuse
- [ ] make go-libp2p have a top level, easy to use package
- [ ] add examples for each of the go-libp2p- repos

### Notes `NA`

### Expected date of completion: `NA`

----------------------------------------------------------------

## Milestone 6 - Move DHT to use IPRS

> Summary: Currently our records are simple protobufs that get stored in the datastore. With this milestone, any record (provider or IPNS) will be a full IPLD object that follows the IPRS spec.

### Tasks

- [ ] Technical Discussion: Decide if whether we should support two types of records or simply create another universe of the DHT
- [ ] Migrate go-ipfs to use IPRS
- [ ] js-libp2p-dht (also known as dht legacy)

### Notes `NA`

### Expected date of completion: `NA`

----------------------------------------------------------------

## Milestone 7 - Break DHT into Peer Routing and Content Routing

> Summary: `NA`

### Tasks

- [ ] Technical Discussion: Revisit proposed spec

### Notes

This milestone is still very green, however something we know for sure we want to do.

### Expected date of completion: `NA`

----------------------------------------------------------------

## Milestone 8 - Packet Switching / Overlay Network

> Summary: The packet switched overlay network

### Tasks

- [ ] Packet switch implemented and integrated into go-ipfs
  - https://github.com/ipfs/notes/issues/143

### Dependencies

- [ ] Multigram specced and implemented in Golang
  - https://github.com/multiformats/multigram

### Notes

The overlay network will not have any multi-hop routing for now.
We only introduce the packet switch and continue being a direct single-hop network.

### Expected date of completion: ``

## Milestone - Improve Connectivity of go-libp2p/go-ipfs

> Summary: Improve the connectivity of go-libp2p. Have the technical discussions that will enable us to follow a plan to achieve a better connected graph.

### Tasks

- [ ] Technical Discussion: Define a strategy for connection closing
  - https://github.com/libp2p/go-libp2p/issues/45
  - https://github.com/ipfs/go-ipfs/issues/3065
  - https://github.com/ipfs/notes/issues/110
- [ ] Implement the strategy for connection closing in go-libp2p
- [ ] Technical Discussion: Definte a strategy for Relay (or multiple Relay)
  - https://github.com/ipfs/notes/issues/110
- [ ] Implement Relay in go-libp2p
- [ ] Technical Discussion: Improve NAT Traversal
  - https://github.com/ipfs/notes/issues/110
  - https://github.com/ipfs/go-ipfs/issues/2509
- [ ] Improve NAT traversal in go-libp2p

### Notes `NA`

### Expected date of completion: `NA`
