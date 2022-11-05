package template

type AppInfo struct {
	Contract    string `json:"contract"`
	App         string `json:"app"`
	User        string `json:"user"`
	Network     string `json:"network"`
	Date        string `json:"date"`
	ExplorerUrl string `json:"explorerUrl"`
}

type Erc20Transfer struct {
	Tx          string `json:"tx"`
	From        string `json:"from"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	TokenName   string `json:"tokenName"`
	TokenSymbol string `json:"tokenSymbol"`
}

type ContractCall struct {
	Tx   string `json:"tx"`
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data"`
}
