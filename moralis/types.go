package moralis

type Block struct {
	Number    string `json:"number"`
	Hash      string `json:"hash"`
	Timestamp string `json:"timestamp"`
}

type Log struct {
	LogIndex        string      `json:"logIndex"`
	TransactionHash string      `json:"transactionHash"`
	Address         string      `json:"address"`
	Data            string      `json:"data"`
	Topic0          string      `json:"topic0"`
	Topic1          string      `json:"topic1"`
	Topic2          *string     `json:"topic2"`
	Topic3          interface{} `json:"topic3"`
}

type Tx struct {
	Hash                     string      `json:"hash"`
	Gas                      string      `json:"gas"`
	GasPrice                 string      `json:"gasPrice"`
	Nonce                    string      `json:"nonce"`
	Input                    string      `json:"input"`
	TransactionIndex         string      `json:"transactionIndex"`
	FromAddress              string      `json:"fromAddress"`
	ToAddress                string      `json:"toAddress"`
	Value                    string      `json:"value"`
	Type                     string      `json:"type"`
	V                        string      `json:"v"`
	R                        string      `json:"r"`
	S                        string      `json:"s"`
	ReceiptCumulativeGasUsed string      `json:"receiptCumulativeGasUsed"`
	ReceiptGasUsed           string      `json:"receiptGasUsed"`
	ReceiptContractAddress   interface{} `json:"receiptContractAddress"`
	ReceiptRoot              interface{} `json:"receiptRoot"`
	ReceiptStatus            string      `json:"receiptStatus"`
}

type NftTransfer struct {
	Operator          interface{} `json:"operator"`
	From              string      `json:"from"`
	To                string      `json:"to"`
	TokenId           string      `json:"tokenId"`
	Amount            string      `json:"amount"`
	TransactionHash   string      `json:"transactionHash"`
	LogIndex          string      `json:"logIndex"`
	Contract          string      `json:"contract"`
	TokenName         string      `json:"tokenName"`
	TokenSymbol       string      `json:"tokenSymbol"`
	TokenContractType string      `json:"tokenContractType"`
}

type Erc20Transfer struct {
	TransactionHash   string `json:"transactionHash"`
	Contract          string `json:"contract"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	LogIndex          string `json:"logIndex"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimals     string `json:"tokenDecimals"`
	ValueWithDecimals string `json:"valueWithDecimals"`
}

type Erc20Approval struct {
	TransactionHash   string `json:"transactionHash"`
	LogIndex          string `json:"logIndex"`
	Contract          string `json:"contract"`
	Spender           string `json:"spender"`
	Owner             string `json:"owner"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimals     string `json:"tokenDecimals"`
	ValueWithDecimals string `json:"valueWithDecimals"`
}

type NftApproval struct {
	TransactionHash   string `json:"transactionHash"`
	LogIndex          string `json:"logIndex"`
	Contract          string `json:"contract"`
	Owner             string `json:"owner"`
	Approved          string `json:"approved"`
	TokenId           string `json:"tokenId"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenContractType string `json:"tokenContractType"`
}

type Payload struct {
	Confirmed      bool            `json:"confirmed"`
	ChainId        string          `json:"chainId"`
	Abi            []interface{}   `json:"abi"`
	StreamId       string          `json:"streamId"`
	Tag            string          `json:"tag"`
	Retries        int             `json:"retries"`
	Block          Block           `json:"block"`
	Logs           []Log           `json:"logs"`
	Txs            []Tx            `json:"txs"`
	TxsInternal    []interface{}   `json:"txsInternal"`
	Erc20Transfers []Erc20Transfer `json:"erc20Transfers"`
	Erc20Approvals []Erc20Approval `json:"erc20Approvals"`
	NftApprovals   struct {
		ERC1155 []NftApproval `json:"ERC1155"`
		ERC721  []NftApproval `json:"ERC721"`
	} `json:"nftApprovals"`
	NftTransfers []NftTransfer `json:"nftTransfers"`
}
