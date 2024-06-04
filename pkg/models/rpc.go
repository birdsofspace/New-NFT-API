package models

import (
	"net/http"
	"strconv"

	"math/big"

	jsonparse "encoding/json"

	"birdsofspace.com/new-nft-api/pkg/channels"
	external2 "birdsofspace.com/new-nft-api/pkg/external"
	"github.com/0xKitsune/go-web3"
	"github.com/0xKitsune/go-web3/contract"
)

type Args struct {
	ChainID int `json:"chain_id"`
	NFTID   int `json:"nft_id"`
}

type JSONServer struct{}

func (T *JSONServer) NFT(r *http.Request, args *Args, reply *string) error {

	if args.ChainID == 0 && args.NFTID == 0 {
		*reply = ""
		return nil
	}

	var nftid int64
	nftid, _ = strconv.ParseInt(strconv.Itoa(args.NFTID), 10, 64)
	client, _ := channels.ClientRegistry[args.ChainID]
	nftContract := contract.NewContract(web3.HexToAddress(external2.GetNFTAddress(args.ChainID)), external2.GetABI(args.ChainID), client)
	resultgetAttributes, _ := nftContract.Call("_tokenIdToAttributes", web3.Latest, big.NewInt(nftid))
	level, _ := nftContract.Call("level", web3.Latest, big.NewInt(nftid))

	var resultImage string
	var resultAnim string
	var resultVideo string

	switch resultgetAttributes["uniqueAttribute"] {
	case "Melodic Mirage":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_melodic_mirage"
			resultImage = "https://nfthost-a5679.web.app/swan-b.png"
			resultVideo = "https://nfthost-a5679.web.app/swan-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_melodic_mirage"
			resultImage = "https://nfthost-a5679.web.app/swan-m.png"
			resultVideo = "https://nfthost-a5679.web.app/swan-m.webm"
		}
	case "Thunderclap Talons":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_thunderclap_talons"
			resultImage = "https://nfthost-a5679.web.app/golden-eagle-b.png"
			resultVideo = "https://nfthost-a5679.web.app/golden-eagle-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_thunderclap_talons"
			resultImage = "https://nfthost-a5679.web.app/golden-eagle-m.png"
			resultVideo = "https://nfthost-a5679.web.app/golden-eagle-m.webm"
		}
	case "Ethereal Glide":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_ethereal_glide"
			resultImage = "https://nfthost-a5679.web.app/sparrow-b.png"
			resultVideo = "https://nfthost-a5679.web.app/sparrow-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_ethereal_glide"
			resultImage = "https://nfthost-a5679.web.app/sparrow-m.png"
			resultVideo = "https://nfthost-a5679.web.app/sparrow-m.webm"
		}
	case "Sonic Dash":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_sonic_dash"
			resultImage = "https://nfthost-a5679.web.app/cardinal-b.png"
			resultVideo = "https://nfthost-a5679.web.app/cardinal-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_sonic_dash"
			resultImage = "https://nfthost-a5679.web.app/cardinal-m.png"
			resultVideo = "https://nfthost-a5679.web.app/cardinal-m.webm"
		}
	case "Whirlwind Waltz":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=baby_whirlwind_waltz"
			resultImage = "https://nfthost-a5679.web.app/cockatiel-b.png"
			resultVideo = "https://nfthost-a5679.web.app/cockatiel-b.webm"
		} else {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=mature_whirlwind_waltz"
			resultImage = "https://nfthost-a5679.web.app/cockatiel-m.png"
			resultVideo = "https://nfthost-a5679.web.app/cockatiel-m.webm"
		}
	case "Toxic Vortex":
		if int(level["0"].(uint8)) == 0 {
			resultAnim = "https://prod-api-central.birdsofspace.com/my.nft/?bird_name=egg"
			resultImage = "https://nfthost-a5679.web.app/egg.png"
			resultVideo = "https://nfthost-a5679.web.app/egg.webm"
		} else if int(level["0"].(uint8)) == 1 {
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
		"attributes": resultgetAttributes,
		"animation":  resultAnim,
		"image":      resultImage,
		"video":      resultVideo,
	}

	mergedResultJSON, _ := jsonparse.Marshal(mergedResult)
	*reply = string(mergedResultJSON)
	return nil
}
