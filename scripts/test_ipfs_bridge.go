package main

import (
	"context"
	"fmt"

	core "github.com/ipfs/go-ipfs/core"
	commands "github.com/ipfs/go-ipfs/core/commands"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
)

func main() {

	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	// Remove bootstrap nodes
	repoConfig, err := r.Config()
	if err != nil {
		panic(err)
	}
	repoConfig.Bootstrap = []string{}

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
	addr, peer, err := commands.ParsePeerParam("/dns4/tiger.musteka.la/tcp/443/wss/QmXFdPj3FuVpkgmNHNTFitkp4DSmVuF6HxNX6tCZr4LFz9")
	if err != nil {
		panic(err)
	}
	_ = nd
	fmt.Printf("---> %v\n-----> %v\n", addr, peer)
	//nd.P2P.Dial()
	// DEBUG

	// DEBUG
	select {}
	// DEBUG
}
