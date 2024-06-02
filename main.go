package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/0xKitsune/go-web3"
	"github.com/0xKitsune/go-web3/abi"
	"github.com/0xKitsune/go-web3/contract"
	"github.com/0xKitsune/go-web3/jsonrpc"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	Name           string   `json:"name"`
	Chain          string   `json:"chain"`
	ChainId        string   `json:"chainId"`
	Network        string   `json:"network"`
	RPC            []string `json:"rpc"`
	Faucets        []string `json:"faucets"`
	InfoURL        string   `json:"infoURL"`
	ShortName      string   `json:"shortName"`
	ChainName      string   `json:"chainName"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"nativeCurrency"`
}

func MapToStruct(data map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	var getRPC = func(chid int) string {
		chainID := strconv.Itoa(chid)
		url := fmt.Sprintf("https://raw.githubusercontent.com/ethereum-lists/chains/master/_data/chains/eip155-%s.json", chainID)

		resp, _ := http.Get(url)
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		var chain Chain
		_ = json.Unmarshal(body, &chain)
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

	var getABI = func(chid int) *abi.ABI {
		id := strconv.Itoa(chid)
		t := time.Now().Unix()
		resp, _ := http.Get("https://raw.githubusercontent.com/birdsofspace/global-config/main/" + id + "/ERC-721/ABI.json?time=" + strconv.FormatInt(t, 10))
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		contractAbi, _ := abi.NewABIFromReader(strings.NewReader(string(body)))
		return contractAbi
	}

	var getNFTAddress = func(chid int) string {
		id := strconv.Itoa(chid)
		t := time.Now().Unix()
		resp, _ := http.Get("https://raw.githubusercontent.com/birdsofspace/global-config/main/" + id + "/ERC-721/CONTRACT_ADDRESS?time=" + strconv.FormatInt(t, 10))
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return strings.TrimSpace(string(body))
	}

	var chainID int
	var nftidInt int
	var nftid int64

	flag.IntVar(&chainID, "c", 137, "Chain ID (default: 137)")
	flag.IntVar(&nftidInt, "n", 18, "NFT ID (default: 18)")
	flag.Parse()

	var err error
	nftid, err = strconv.ParseInt(strconv.Itoa(nftidInt), 10, 64)
	if err != nil {
		log.Fatalf("Invalid value for nftid: %v", err)
	}
	client, _ := jsonrpc.NewClient(getRPC(chainID))

	nftContract := contract.NewContract(web3.HexToAddress(getNFTAddress(chainID)), getABI(chainID), client)
	resultgetAttributes, _ := nftContract.Call("_tokenIdToAttributes", web3.Latest, big.NewInt(nftid))
	matureBirdCost, errMature := nftContract.Call("matureBirdCost", web3.Latest, big.NewInt(nftid))
	maxMatureBirdCost, errMaxMature := nftContract.Call("maxMatureBirdCost", web3.Latest, big.NewInt(nftid))

	var resultmatureBirdCost string
	var resultmaxMatureBirdCost string
	if matureBirdCost["0"] == nil && strings.Contains(errMature.Error(), "not baby bird") {
		resultmatureBirdCost = "not baby bird"
	}
	if maxMatureBirdCost["0"] != nil && strings.Contains(errMaxMature.Error(), "not mature bird") {
		resultmaxMatureBirdCost = "not mature bird"
	}

	var resultImage string
	var resultAnim string
	var resultVideo string

	switch resultgetAttributes["uniqueAttribute"] {
	case "Melodic Mirage":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_melodic_mirage"
			resultImage = "https://nfthost-a5679.web.app/swan-b.png"
			resultVideo = "https://nfthost-a5679.web.app/swan-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_melodic_mirage"
			resultImage = "https://nfthost-a5679.web.app/swan-m.png"
			resultVideo = "https://nfthost-a5679.web.app/swan-m.webm"
		}
	case "Thunderclap Talons":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_thunderclap_talons"
			resultImage = "https://nfthost-a5679.web.app/golden-eagle-b.png"
			resultVideo = "https://nfthost-a5679.web.app/golden-eagle-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_thunderclap_talons"
			resultImage = "https://nfthost-a5679.web.app/golden-eagle-m.png"
			resultVideo = "https://nfthost-a5679.web.app/golden-eagle-m.webm"
		}
	case "Ethereal Glide":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_ethereal_glide"
			resultImage = "https://nfthost-a5679.web.app/sparrow-b.png"
			resultVideo = "https://nfthost-a5679.web.app/sparrow-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_ethereal_glide"
			resultImage = "https://nfthost-a5679.web.app/sparrow-m.png"
			resultVideo = "https://nfthost-a5679.web.app/sparrow-m.webm"
		}
	case "Sonic Dash":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_sonic_dash"
			resultImage = "https://nfthost-a5679.web.app/cardinal-b.png"
			resultVideo = "https://nfthost-a5679.web.app/cardinal-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_sonic_dash"
			resultImage = "https://nfthost-a5679.web.app/cardinal-m.png"
			resultVideo = "https://nfthost-a5679.web.app/cardinal-m.webm"
		}
	case "Whirlwind Waltz":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_whirlwind_waltz"
			resultImage = "https://nfthost-a5679.web.app/cockatiel-b.png"
			resultVideo = "https://nfthost-a5679.web.app/cockatiel-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_whirlwind_waltz"
			resultImage = "https://nfthost-a5679.web.app/cockatiel-m.png"
			resultVideo = "https://nfthost-a5679.web.app/cockatiel-m.webm"
		}
	case "Toxic Vortex":
		if resultmatureBirdCost == "not baby bird" && resultmaxMatureBirdCost == "not mature bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if resultmatureBirdCost != "not baby bird" {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_toxic_vortex"
			resultImage = "https://nfthost-a5679.web.app/vulture-b.png"
			resultVideo = "https://nfthost-a5679.web.app/vulture-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_toxic_vortex"
			resultImage = "https://nfthost-a5679.web.app/vulture-m.png"
			resultVideo = "https://nfthost-a5679.web.app/vulture-m.webm"
		}
	default:
		resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
		resultImage = "https://nfthost-a5679.web.app/egg.png"
		resultVideo = "https://nfthost-a5679.web.app/egg.webm"
	}
	mergedResult := map[string]interface{}{
		"rpc":        getRPC(chainID),
		"attributes": resultgetAttributes,
		"animation":  resultAnim,
		"image":      resultImage,
		"video":      resultVideo,
	}

	mergedResultJSON, _ := json.Marshal(mergedResult)
	fmt.Println(string(mergedResultJSON))

}
