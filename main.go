package main

import (
	jsonparse "encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/gorilla/mux"

	"github.com/0xKitsune/go-web3/jsonrpc"

	"birdsofspace.com/new-nft-api/pkg/channels"
	"birdsofspace.com/new-nft-api/pkg/constants"
	"birdsofspace.com/new-nft-api/pkg/models"
)

func main() {

	var GetRPC = func(chid int) string {
		chainID := strconv.Itoa(chid)
		url := fmt.Sprintf("https://raw.githubusercontent.com/ethereum-lists/chains/master/_data/chains/eip155-%s.json", chainID)

		resp, _ := http.Get(url)
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		var chain models.Chain
		_ = jsonparse.Unmarshal(body, &chain)
		var fastestRpc string
		var fastestTime time.Duration
		fastestRpc = chain.RPC[0]
		for _, rpc := range chain.RPC {
			if strings.Contains(rpc, "https") {
				start := time.Now()

				testCon, err := ethclient.Dial(rpc)
				if err != nil {
					continue
				}
				testCon.Close()
				defer resp.Body.Close()
				_, _ = io.ReadAll(resp.Body)
				elapsed := time.Since(start)
				if fastestTime == 0 || elapsed < fastestTime {
					fastestTime = elapsed
					fastestRpc = rpc
				}
			}

		}
		testCon, err := ethclient.Dial(fastestRpc)
		if err != nil {
			fastestRpc = chain.RPC[0]
		}
		testCon.Close()
		return fastestRpc
	}

	for _, v := range constants.ChainRunning {
		client, _ := jsonrpc.NewClient(GetRPC(v))
		channels.ClientRegistry[v] = client
	}

	s := rpc.NewServer() // Register the type of data requested as JSON

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(models.JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":9119", r)

}
