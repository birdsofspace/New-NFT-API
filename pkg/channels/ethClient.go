package channels

import "github.com/0xKitsune/go-web3/jsonrpc"

var ClientRegistry = make(map[int]*jsonrpc.Client)
