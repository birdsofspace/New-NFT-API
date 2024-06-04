package external

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/0xKitsune/go-web3/abi"
)

var GetABI = func(chid int) *abi.ABI {
	id := strconv.Itoa(chid)
	t := time.Now().Unix()
	resp, _ := http.Get("https://raw.githubusercontent.com/birdsofspace/global-config/main/" + id + "/ERC-721/ABI.json?time=" + strconv.FormatInt(t, 10))
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	contractAbi, _ := abi.NewABIFromReader(strings.NewReader(string(body)))
	return contractAbi
}
