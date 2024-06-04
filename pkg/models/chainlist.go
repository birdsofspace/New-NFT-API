package models

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
