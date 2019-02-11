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

| libp2p Node                                  | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`libp2p`**                                 | :green_apple: | :green_apple:   | :green_apple:   | :green_apple: | :green_apple: |

| Identify Protocol                            | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`Identify`**                               | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |


| Transport Protocols                          | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`TCP`**                                    | :green_apple: | :green_apple:   | :green_apple:   | :green_apple: | :lemon:       |
| **`UTP`**                                    | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`UDP`**                                    | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :tomato:      |
| **`WebSockets`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`WebRTC`**                                 | :tomato:      | :green_apple:   | :green_apple:   | :tomato:      | :tomato:    |
| **`SCTP`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`Tor`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`i2p`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`cjdns`**                                  | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`Bluetooth LE`**                           | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`Audio TP`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`Zerotier`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`QUIC`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |


| Stream Muxers                                | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`benchmarks`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |
| **`muxado`**                                 | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`spdystream`**                             | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`multiplex`**                              | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :lemon:       |
| **`spdy`**                                   | :tomato:      | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |
| **`http2`**                                  | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`QUIC`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`yamux`**                                  | :chestnut:    | :chestnut:      | :chestnut:      | :chestnut:    | :tomato:      |


| Switch                                       | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`Dialer stack`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |
| **`Switch`**                                 | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :lemon:       |



| NAT Traversal                                | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`nat-pmp`**                                | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :tomato:    |
| **`upnp`**                                   | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :tomato:    |
| **`ext addr discovery`**                     | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`STUN-like`**                              | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`line-switch relay`**                      | :green_apple: | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`pkt-switch relay`**                       | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |


| Peer Discovery                               | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`mDNS`**                                   | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`bootstrap list`**                         | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :green_apple: |
| **`DHT query`**                              | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`PEX`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`DNS`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |



| Content Routing                              | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`Kademlia DHT`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`pub/sub`**                                | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :tomato:      |
| **`PHT`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |



| Peer Routing                                 | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`Kademlia DHT`**                           | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :tomato:      |
| **`pub/sub`**                                | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :tomato:      |
| **`PHT`**                                    | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |


| Exchange                                     | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`HTTP`**                                   | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |
| **`Bitswap`**                                | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |
| **`Bittorrent`**                             | :green_apple: | :green_apple:   | :green_apple:   | :tomato:      | :chestnut:    |


| Consensus                                    | Go            | JS - Node.js    |  JS - Browser   | Rust          | Python        |
| -------------------------------------------- | :-----------: | :-------------: | :-------------: | :-----------: | :-----------: |
| **`Paxos`**                                  | :chestnut:    | :chestnut:      | :chestnut:      | :chestnut:    | :chestnut:    |
| **`Raft`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`PBTF`**                                   | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |
| **`Nakamoto`**                               | :tomato:      | :tomato:        | :tomato:        | :tomato:      | :chestnut:    |


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

