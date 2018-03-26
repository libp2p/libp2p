WIP WIP WIP

# LIBP2P Overview

Libp2p is an extensible, modular peer-to-peer framework designed to make p2p
communication easier than IPC.

* Extensible: It's easy to implement new underlying transport protocols
  (protocols for connecting two peers).
* Modular: You can opt out of the pieces you don't need (e.g., the DHT, local
  peer discovery, etc.).
  
## Communicating

There are two concepts one needs to understand to start using libp2p:

1. Peer IDs
2. Protocol-negotiated streams.

Libp2p "nodes" (processes) identify each other by peer ID (PID), a
[multihash][multihash] of the peer's public key. This allows any two nodes to
identify each other, regardless of location, IP address, or transport protocol.

To communicate, a libp2p node "dials" (connects) to another peer by peer ID then
negotiates a protocol using [multistream select][mss]. If the receiving peer has
registered a protocol "handler" for the protocol requested, it will be invoked
and passed the inbound connection.

These are the two primary user-facing concepts; the rest of this document will
focus more on the structure of libp2p internally.

## Multiaddr

To speak to a remote peer, a libp2p node needs to know how to *address* this
remote peer. For this, we use [multiaddrs][multiaddr].

Multiaddrs are recipes for transport channels. They describe both the protocols
and the addresses within those protocols' namespaces needed to find and
communicate with the remote peer.

Usually, one shouldn't need to manually keep track of which multiaddrs each peer
uses, each peer tracks known mulitaddrs in an address book. However, you'll need
to configure peer discovery for this to work properly.

## Peer discovery

While one can use libp2p without peer discovery, doing so requires manually
managing the mapping between peer IDs and multiaddrs. Peer discovery services
perform this task automatically.

Currently, we have two primary peer discovery mechanisms:

1. A [DHT](dht). Nodes using the DHT for peer discovery will periodically record
   a mapping of their peer ID to their current multiaddrs in the DHT. Libp2p
   nodes wishing to dial this peer, will first lookup the peer's ID in the DHT
   to find their current address and then use the information contained therein
   to dial the peer.
2. Local peer discovery. Local peer discovery uses mDNS (local network multicast
   DNS) to discover peers on the local network. This allows, e.g., your phone
   and laptop to find each other if they're on the same WiFi (usually).

## Transport

In libp2p, transports are the underlying mechanism nodes use to communicate with
each other. When using libp2p, one will almost never deal with transports
directly after configuring them.

All transports must provide the following features:

* Reliability: Connections must be reliable streams (although we plan on adding
  packet transports, both reliable and unreliable).
* Encryption: Connections must be end-to-end encrypted.
* Authentication: The endpoints, RemotePeer and LocalPeer, must be authenticated.
* Multiplexing: It must be possible to multiplex multiple reliable streams over
  a single transport connection.

Encryption and authentication are pretty self-explanatory. When two peers communicate, each must be assured that:

1. It is communicating with the correct peer.
2. It's communication with that peer is private.

Multiplexing allows us to open multiple virtual connections over some set of
underlying "real" connections (usually just one). Conceptually, these streams
may as well be independent connections but, practically, they make opening and
maintaining multiple "connections" to the same peer relatively cheep.

As many transports only provide some of these features, we provide libraries for
negotiating stream multiplexing, encryption, and authentication transports for
less featureful protocols.

[multihash]: https://github.com/multiformats/multihash
[multiaddr]: https://github.com/multiformats/multiaddr
[mss]: https://github.com/multiformats/multistream-select/
[dht]: https://en.wikipedia.org/wiki/Distributed_hash_table
