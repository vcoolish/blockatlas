package ethereum

type Page struct {
	Total uint  `json:"total"`
	Docs  []Doc `json:"docs"`
}

type Doc struct {
	Ops         []Op   `json:"operations"`
	Contract    string `json:"contract"`
	ID          string `json:"id"`
	BlockNumber uint64 `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
	Nonce       uint64 `json:"nonce"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	GasUsed     string `json:"gasUsed"`
	Input       string `json:"input"`
	Error       string `json:"error"`
	Coin        uint   `json:"coin"`
}

type Op struct {
	TxID     string    `json:"transactionId"`
	Contract *Contract `json:"contract"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	Coin     uint      `json:"coin"`
}

type Contract struct {
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	Decimals    uint   `json:"decimals"`
	TotalSupply string `json:"totalSupply"`
	Name        string `json:"name"`
}

type NodeInfo struct {
	LatestBlock int64 `json:"latest_block"`
}

type Collection struct {
	Name        string                 `json:"name"`
	ImageUrl    string                 `json:"image_url"`
	ExternalUrl string                 `json:"external_url"`
	Total       int                    `json:"owned_asset_count"`
	Contract    []PrimaryAssetContract `json:"primary_asset_contracts"`
}

type PrimaryAssetContract struct {
	Address     string `json:"address"`
	NftVersion  string `json:"nft_version"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type CollectiblePage struct {
	Collectibles []Collectible `json:"assets"`
}

type Collectible struct {
	TokenId         string        `json:"token_id"`
	AssetContract   AssetContract `json:"asset_contract"`
	ImageUrl        string        `json:"image_url"`
	ImagePreviewUrl string        `json:"image_preview_url"`
	Name            string        `json:"name"`
	ExternalLink    string        `json:"external_link"`
	Description     string        `json:"description"`
}

type AssetContract struct {
	Address  string `json:"address"`
	Category string `json:"name"`
}
