# libp2p Requirements: Roadmap to Feature Complete

## Table of Contents

- [Description](#description)
- [Cross Language Milestones](#cross-language-milestones)
- [libp2p Modules Implementations](#libp2p-modules-implementations)
- [Performance Milestones](#performance-milestones)
- [Polish Level](#polish-level)

## Description

In order for libp2p to achieve its goals, the project will need to achieve the items described in this document.

This document is in the process of being updated.  Currently known deficiencies:
 * Rust is incomplete (some tomatoes should be apples)
 * Should we divide Node.js and Browser, or just a have a single JS column?
 * Needs a proofread from implementors to make sure the fruits are correct
 * [Performance Milestones](#performance-milestones) is incomplete - need to agree on what these are

As you improve this document, please remove items from the list above.

### libp2p Modules Implementations 

The libp2p protocols are the protocol stack for the modular libp2p protocols library. These form a p2p-first network stack, with no assumptions, self description, and modular interfaces. More at https://github.com/libp2p/specs/

> Legend: :green_apple: Done &nbsp; :lemon: In Progress &nbsp; :tomato: Missing &nbsp; :chestnut: Not planned

| libp2p Node                                  | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`libp2p`**                                 | :green_apple: | :green_apple:   | :green_apple:   | :green_apple: |

| Identify Protocol                            | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`Identify`**                               | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |


| Transport Protocols                          | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`TCP`**                                    | :green_apple: | :green_apple:   | :green_apple:   | :green_apple: |
| **`UTP`**                                    | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`UDP`**                                    | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`WebSockets`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`WebRTC`**                                 | :tomato:      | :green_apple:   | :green_apple:   | :tomato:      |
| **`SCTP`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`Tor`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`i2p`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`cjdns`**                                  | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`Bluetooth LE`**                           | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`Audio TP`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`Zerotier`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`QUIC`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      |


| Stream Muxers                                | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`benchmarks`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`muxado`**                                 | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`spdystream`**                             | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`multiplex`**                              | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`spdy`**                                   | :tomato:      | :green_apple:   | :green_apple:   | :tomato:      |
| **`http2`**                                  | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`QUIC`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      |


| Switch                                       | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`Dialer stack`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`Switch`**                                 | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |



| NAT Traversal                                | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`nat-pmp`**                                | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`upnp`**                                   | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`ext addr discovery`**                     | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`STUN-like`**                              | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`line-switch relay`**                      | :green_apple: | :tomato:        | :tomato:        | :tomato:      |
| **`pkt-switch relay`**                       | :tomato:      | :tomato:        | :tomato:        | :tomato:      |


| Peer Discovery                               | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`mDNS`**                                   | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`bootstrap list`**                         | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`DHT query`**                              | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`PEX`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`DNS`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |



| Content Routing                              | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`Kademlia DHT`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`pub/sub`**                                | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`PHT`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |



| Peer Routing                                 | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`Kademlia DHT`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`pub/sub`**                                | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`PHT`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      |


| Exchange                                     | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`HTTP`**                                   | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`Bitswap`**                                | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |
| **`Bittorrent`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      |


| Consensus                                    | Go            | JS - Node.js    |  JS - Browser   | Rust          |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: |
| **`Paxos`**                                  | :chestnut:    | :chestnut:      | :chestnut:      | :chestnut:    |
| **`Raft`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`PBTF`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      |
| **`Nakamoto`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      |


### Cross Language Milestones

- [ ] Standardized interfaces (using a neutral IDL) to define conformance in any language implementation
- [ ] Interoperability testing framework

### Performance Milestones

`TODO` needs performance milestones

### Polish level

The polish level achieved defined how much effort has gone into polishing all the corners of the tools, APIs, documentation, and so on. Early on, things will be much more rough as we need to focus on implementation speed over final polish. As we reach completion of the implementations, we turn our focus towards scalability, stability, and security. After that, we focus on end user polish.	

- [ ] Polish level: alpha - individuals innovators
- [ ] Polish level: alpha - production for early adopters
- [ ] Polish level: beta - production for small orgs (in progress)
- [ ] Polish level: beta - production for large orgs
- [ ] Polish level: 1.0 - production for all

