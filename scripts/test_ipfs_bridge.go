package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"strings"

	pstore "gx/ipfs/QmXZSd1qR5BxZkPyuwfT5jpqQFScZccoZvDneXsKzCNHWX/go-libp2p-peerstore"
	ma "gx/ipfs/QmcyqRMCAXVtYPS4DiBrA7sezL9rRGfW8Ctx7cywL4TXJj/go-multiaddr"

	core "github.com/ipfs/go-ipfs/core"
	config "github.com/ipfs/go-ipfs/repo/config"
	iaddr "github.com/ipfs/go-ipfs/thirdparty/ipfsaddr"
)

func main() {
	// Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := &core.BuildCfg{
		NilRepo: true,
		Online:  true,
		ExtraOpts: map[string]bool{
			"mplex": true,
		},
	}

	// Remove the hardcoded bootstrap nodes
	config.DefaultBootstrapAddresses = []string{}

	// The new IPFS Node
	node, err := core.NewNode(ctx, cfg)
	if err != nil {
		panic(err)
	}

	a, err := iaddr.ParseString("/ip4/52.176.89.220/tcp/4001/ipfs/Qmc7etyUd9tEa3ZBD3LCTMDL96qcMi8cKfHEiLt5nhVdVC")
	if err != nil {
		panic(err)
	}
	pi := pstore.PeerInfo{
		ID:    a.ID(),
		Addrs: []ma.Multiaddr{a.Transport()},
	}

	output := "connect " + pi.ID.Pretty()
	err = node.PeerHost.Connect(ctx, pi)
	if err != nil {
		fmt.Printf("%s failure: %s\n", output, err)
	} else {
		fmt.Printf("success")
	}

	// DEBUG
	// Keep it running!
	select {}
	// DEBUG
}

func random8ByteString() string {
	a := make([]byte, 8)
	if _, err := rand.Read(a); err != nil {
		panic(err)
	}
	return strings.ToLower(fmt.Sprintf("%X", a))
}
