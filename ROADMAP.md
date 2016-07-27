libp2p Roadmap
==============

> We track the development of the libp2p project through Github issues and [Waffle.io](https://waffle.io/libp2p/libp2p). See our waffle board at: [https://waffle.io/libp2p/libp2p](https://waffle.io/libp2p/libp2p)

# Milestone 1 - JS and Go libp2p interop

> Summary: Achieve interoperability between the JS and Go implementations, specificially with secio.

### Tasks:

- [ ] Migration to pull-streams (track: https://github.com/ipfs/js-ipfs/issues/403). Also part of js-ipfs Milestone 2
- [x] js-secio - https://github.com/libp2p/js-libp2p-secio
- [ ] Interop tests (track: https://github.com/ipfs/js-libp2p-ipfs/issues/15)
- [ ] Integrate the `ping protocol` in js-libp2p (track: https://github.com/ipfs/js-libp2p-ipfs/pull/13)

### Dependencies:

NA

### Requirements by other projects:

- js-ipfs and go-ipfs interop

### Notes:

NA

##### Expected date of completion: `Q3`

----------------------------------------------------------------

# Milestone 2 - Improve Connectivity of go-libp2p/go-ipfs

> Summary: Improve the connectivity of go-libp2p. Have the technical discussions that will enable us to follow a plan to achieve a better connected graph.

### Tasks:

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

### Dependencies:

### Requirements by other projects:

### Notes:

##### Expected date of completion: `Q3`

----------------------------------------------------------------

# Milestone 3 - The libp2p.io

> Summary: The main goal of this milestone is to get more developers excited about libp2p and learn how to use it.

### Tasks:

- [ ] Replicate(point to) IPFS community guidelines for the libp2p project
- [ ] libp2p.io - https://github.com/libp2p/libp2p-website

### Dependencies:

### Requirements by other projects:

### Notes:

The Website and replicating the IPFS community guidelines should be the priority

##### Expected date of completion: `Q3`

----------------------------------------------------------------

# Milestone 4 - Developer Experience

> Summary: Get libp2p into a state that is super easy to use and understand.

### Tasks:

- [ ] Examples
  - [ ] How to use js-libp2p (dial, multi transport, stream muxing, identify, protocol multiplexing)
  - [ ] How to use go-libp2p (everything js-libp2p + DHT)
- [ ] Spec
  - [ ] Move spec to the libp2p org
  - [ ] Revisit and improve
- [ ] interface-libp2p
  - [ ] Make js-libp2p-ipfs and js-libp2p-ipfs-browser use it
  - [ ] Create js-libp2p, a base class that other libp2p builds can reuse

### Dependencies:

### Requirements by other projects:

### Notes:

##### Expected date of completion: `Q4`

----------------------------------------------------------------

# Milestone 5 - Move DHT to use IPRS

> Summary: Currently our records are simple protobufs that get stored in the datastore. With this milestone, any record (provider or IPNS) will be a full IPLD object that follows the IPRS spec.

### Tasks:

- [ ] Technical Discussion: Decide if whether we should support two types of records or simply create another universe of the DHT
- [ ] Migrate go-ipfs to use IPRS
- [ ] js-libp2p-dht (also known as dht legacy)

### Dependencies:

### Requirements by other projects:

### Notes:

##### Expected date of completion: `Q4`

----------------------------------------------------------------

# Milestone 6 - Break DHT into Peer Routing and Content Routing

> Summary:

### Tasks:

- [ ] Technical Discussion: Revisit proposed spec

### Dependencies:

### Requirements by other projects:

### Notes:

This milestone is still very green, however something we know for sure we want to do.


##### Expected date of completion: `TBA`

----------------------------------------------------------------

# Milestone 7 - Packet Switching / Overlay Network

> Summary: The packet switched overlay network

### Tasks:

- [ ] Packet switch implemented and integrated into go-ipfs
  - https://github.com/ipfs/notes/issues/143

### Dependencies:

- [ ] Multigram specced and implemented in Golang
  - https://github.com/multiformats/multigram

### Requirements by other projects:

### Notes:

The overlay network will not have any multi-hop routing for now.
We only introduce the packet switch and continue being a direct single-hop network.

##### Expected date of completion: Q3
