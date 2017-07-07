package main

import (
	"context"
	"fmt"

	"github.com/ipfs/go-ipfs/core"
	config "github.com/ipfs/go-ipfs/repo/config"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
)

func main() {
	// We don't want the node to discover nor connecting
	// other peers more than the one we are giving
	config.DefaultBootstrapAddresses = []string{}

	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	nd, err := core.NewNode(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// DEBUG
	fmt.Printf("%v\n", nd)
	// DEBUG

	// DEBUG
	select {}
	// DEBUG
}
