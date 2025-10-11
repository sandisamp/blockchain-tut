package main

import (
	"github.com/sandisamp/blockchain-tut/network"
	"time"
)

// Server: Needed to host the blockchain and communicate with other nodes in the network.
// Transport Layer ==> tcp/ udp: Needed to define how the nodes will communicate with each other (e.g. TCP for reliable connections, UDP for fast but potentially unreliable connections).
// Block: Needed to store and link together transactions in a secure and tamper-proof manner.
// Tx: Needed to represent and validate transactions between nodes.
// Keypairs: Needed to securely sign and verify transactions and ensure that only authorized nodes can make changes to the blockchain.

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello from remote"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
