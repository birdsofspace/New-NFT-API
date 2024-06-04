package external

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var GetNFTAddress = func(chid int) string {
	id := strconv.Itoa(chid)
	t := time.Now().Unix()
	resp, _ := http.Get("https://raw.githubusercontent.com/birdsofspace/global-config/main/" + id + "/ERC-721/CONTRACT_ADDRESS?time=" + strconv.FormatInt(t, 10))
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return strings.TrimSpace(string(body))
}
