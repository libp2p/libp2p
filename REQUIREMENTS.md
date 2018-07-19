# libp2p Requirements Document aka Roadmap to libp2p feature complete

## Table of Contents

- [Description](#description)
- [Cross Language Milestones](#cross-language-milestones)
- [libp2p Protocols](#libp2p-protocols)
- [Polish Level](#polish-level)

## Description

In order for libp2p to achieve its goals, the project will need to achieve the items described in this document.

### Cross Language Milestones

- [ ] Standardized interfaces (using a neutral IDL) to define conformance in any language implementation
- [ ] Interoperability testing framework

### libp2p Modules Implementations 

The libp2p protocols are the protocol stack for the modular libp2p protocols library. These form a p2p-first network stack, with no assumptions, self description, and modular interfaces. More at https://github.com/ipfs/specs/tree/master/libp2p

> Legend: :green_apple: Done &nbsp; :lemon: In Progress &nbsp; :tomato: Missing &nbsp; :chestnut: Not planned

| Transports                                   | Go            | JS - Node.js  |  JS - Browser | Rust          |
| -------------------------------------------- | :-----------: | :-----------: | :-----------: | :-----------: |
| **`TCP`**                                    | :green_apple: | :green_apple: | :chestnut:    | :green_apple: |
| **`WebRTC`**                                 | :tomato:      | :green_apple: | :green_apple: | :tomato:      |


`TODO` transform the below into tables

- [x] libp2p Node
  - [x] libp2p impl in Go
  - [x] libp2p impl in JS
  - [x] libp2p impl in Rust
- [x] Identify Protocol interfaces
  - [x] Identify: impl in Go
  - [x] Identify: impl in JS
  - [ ] Identify: impl in Rust
- [x] Transport Protocol interfaces
  - [x] Transport protocol: TCP in Go
  - [x] Transport protocol: TCP in JS node
  - [ ] Transport protocol: TCP in Rust
  - [x] Transport protocol: UTP in Go
  - [x] Transport protocol: UTP in JS node
  - [ ] Transport protocol: UTP in Rust
  - [x] Transport protocol: UDT in Go
  - [ ] Transport protocol: UDT in JS node
  - [ ] Transport protocol: UDT in Rust
  - [x] Transport protocol: WebSockets in Go
  - [x] Transport protocol: WebSockets in JS node
  - [x] Transport protocol: WebSockets in JS browser
  - [ ] Transport protocol: WebSockets in Rust
  - [ ] Transport protocol: WebRTC in Go
  - [x] Transport protocol: WebRTC in JS node
  - [x] Transport protocol: WebRTC in JS browser
  - [ ] Transport protocol: WebRTC in Rust
  - [ ] Transport protocol: SCTP in Go
  - [ ] Transport protocol: SCTP in JS node
  - [ ] Transport protocol: SCTP in Rust
  - [ ] Transport protocol: tor in Go
  - [ ] Transport protocol: tor in JS node
  - [ ] Transport protocol: tor in Rust
  - [ ] Transport protocol: i2p in Go
  - [ ] Transport protocol: i2p in JS node
  - [ ] Transport protocol: i2p in Rust
  - [ ] Transport protocol: cjdns in Go
  - [ ] Transport protocol: cjdns in JS node
  - [ ] Transport protocol: cjdns in Rust
  - [ ] Transport protocol: Bluetooth LE in Go
  - [ ] Transport protocol: Bluetooth LE in JS node
  - [ ] Transport protocol: Bluetooth LE in Rust
  - [ ] Transport protocol: Audio TP in Go
  - [ ] Transport protocol: Audio TP in JS node
  - [ ] Transport protocol: Audio TP in Rust
  - [ ] Transport protocol: Zerotier in Go
  - [ ] Transport protocol: Zerotier in JS node
  - [ ] Transport protocol: Zerotier in Rust
  - [ ] Transport protocol: QUIC in Go
  - [ ] Transport protocol: QUIC in JS node
  - [ ] Transport protocol: QUIC in Rust
- [x] Stream Muxer: interfaces
  - [x] Stream Muxer: benchmarks in Go
  - [x] Stream Muxer: benchmarks in JS
  - [ ] Stream Muxer: benchmarks in Rust
  - [x] Stream Muxer: muxado in Go
  - [ ] Stream Muxer: muxado in Rust
  - [x] Stream Muxer: spdystream in Go
  - [x] Stream Muxer: multiplex in Go
  - [x] Stream Muxer: multiplex in JS
  - [ ] Stream Muxer: multiplex in Rust
  - [x] Stream Muxer: spdy in JS
  - [ ] Stream Muxer: http2 in Go
  - [ ] Stream Muxer: http2 in JS
  - [ ] Stream Muxer: http2 in Rust
  - [ ] Stream Muxer: QUIC in Go
  - [ ] Stream Muxer: QUIC in JS
  - [ ] Stream Muxer: QUIC in Rust
- [x] Switch: interfaces
  - [x] Dialer stack in Go
  - [x] Dialer stack in JS
  - [ ] Dialer stack in Rust
  - [x] Switch implementation in Go
  - [x] Switch implementation in JS
  - [ ] Switch implementation in Rust
- [x] NAT Traversal: interfaces
  - [x] nat-pmp implementation in Go
  - [ ] nat-pmp implementation in JS
  - [ ] nat-pmp implementation in Rust
  - [x] upnp implementation in Go
  - [ ] upnp implementation in JS
  - [ ] upnp implementation in Rust
  - [x] ext addr discovery in Go
  - [ ] ext addr discovery in JS
  - [ ] ext addr discovery in Rust
  - [ ] STUN-like implementation in Go
  - [ ] STUN-like implementation in JS
  - [ ] STUN-like implementation in Rust
  - [x] line-switch relay implementation in Go
  - [ ] line-switch relay implementation in JS
  - [ ] line-switch relay implementation in Rust
  - [ ] pkt-switch relay implementation in Go
  - [ ] pkt-switch relay implementation in JS
  - [ ] pkt-switch relay implementation in Rust
- [x] Peer Discovery: interfaces
  - [x] Peer Discovery: mDNS in Go
  - [x] Peer Discovery: mDNS in JS
  - [ ] Peer Discovery: mDNS in Rust
  - [x] Peer Discovery: bootstrap list in JS
  - [x] Peer Discovery: bootstrap list in Go
  - [ ] Peer Discovery: bootstrap list in Rust
  - [x] Peer Discovery: DHT query in JS
  - [x] Peer Discovery: DHT query in Go
  - [ ] Peer Discovery: DHT query in Rust
  - [ ] Peer Discovery: PEX in JS
  - [ ] Peer Discovery: PEX in Go
  - [ ] Peer Discovery: PEX in Rust
  - [ ] Peer Discovery: DNS in JS
  - [ ] Peer Discovery: DNS in Go
  - [ ] Peer Discovery: DNS in Rust
- [x] Content Routing: protocol interfaces
  - [x] Content Routing: Kademlia DHT impl in Go
  - [x] Content Routing: Kademlia DHT impl in JS
  - [ ] Content Routing: Kademlia DHT impl in Rust
  - [ ] Content Routing: pub/sub impl in Go
  - [ ] Content Routing: pub/sub impl in JS
  - [ ] Content Routing: pub/sub impl in Rust
  - [ ] Content Routing: PHT in Go
  - [ ] Content Routing: PHT in JS
  - [ ] Content Routing: PHT in Rust
- [x] Peer Routing: protocol interfaces
  - [x] Peer Routing: Kademlia DHT impl in Go
  - [x] Peer Routing: Kademlia DHT impl in JS
  - [ ] Peer Routing: Kademlia DHT impl in Rust
  - [ ] Peer Routing: pub/sub impl in Go
  - [ ] Peer Routing: pub/sub impl in JS
  - [ ] Peer Routing: pub/sub impl in Rust
  - [ ] Peer Routing: PHT in Go
  - [ ] Peer Routing: PHT in JS
  - [ ] Peer Routing: PHT in Rust
- [x] Exchange: protocol interfaces
  - [x] Exchange: HTTP in Go
  - [x] Exchange: HTTP in JS
  - [ ] Exchange: HTTP in Rust
  - [x] Exchange: Bitswap in Go
  - [x] Exchange: Bitswap in JS
  - [ ] Exchange: Bitswap in Rust
  - [x] Exchange: Bittorrent in Go
  - [x] Exchnage: Bittorrent in JS
  - [ ] Exchnage: Bittorrent in Rust
- [ ] Consensus: protocol interfaces
  - [ ] Consensus: Paxos
  - [ ] Consensus: Raft
  - [ ] Consensus: PBFT
  - [ ] Consensus: Nakamoto

### Specifications

`TODO` needs a spec requirements/goals as well

### Performance Milestones

`TODO` needs performance milestones

### Polish level

The polish level achieved defined how much effort has gone into polishing all the corners of the tools, APIs, documentation, and so on. Early on, things will be much more rough as we need to focus on implementation speed over final polish. As we reach completion of the implementations, we turn our focus towards scalability, stability, and security. After that, we focus on end user polish.	

- [ ] Polish level: alpha - individuals innovators
- [ ] Polish level: alpha - production for early adopters
- [ ] Polish level: beta - production for small orgs (in progress)
- [ ] Polish level: beta - production for large orgs
- [ ] Polish level: 1.0 - production for all

