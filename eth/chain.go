package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"os"
)

type Explorer struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Standard string `json:"standard"`
}

type Chain struct {
	Name           string        `json:"name"`
	Chain          string        `json:"chain"`
	Icon           string        `json:"icon"`
	Rpc            []string      `json:"rpc"`
	Faucets        []interface{} `json:"faucets"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"nativeCurrency"`
	InfoURL   string `json:"infoURL"`
	ShortName string `json:"shortName"`
	ChainId   int    `json:"chainId"`
	NetworkId int    `json:"networkId"`
	Slip44    int    `json:"slip44"`
	Ens       struct {
		Registry string `json:"registry"`
	} `json:"ens"`
	Explorers []Explorer `json:"explorers"`
}

func hexToChainId(hex string) int64 {
	decodeBig, err := hexutil.DecodeBig(hex)
	if err != nil {
		return 0
	}
	return decodeBig.Int64()
}

func GetChainInfoByHexId(hexId string) (*Chain, error) {
	dat, err := os.ReadFile(fmt.Sprintf("./chains/eip155-%d.json", hexToChainId(hexId)))
	if err != nil {
		return nil, err
	}
	var chain Chain

	if err := json.Unmarshal(dat, &chain); err != nil {
		return nil, err
	}
	return &chain, nil
}

func GetChainInfoByChainId(chainId string) (*Chain, error) {
	dat, err := os.ReadFile(fmt.Sprintf("./chains/eip155-%s.json", chainId))
	if err != nil {
		return nil, err
	}
	var chain Chain

	if err := json.Unmarshal(dat, &chain); err != nil {
		return nil, err
	}
	return &chain, nil
}

func IsContract(rpcUrl string, address common.Address) bool {
	log.Infof("Rpc url:", rpcUrl)
	dial, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Errorf("ethclient.Dial err: %v \n", err)
		return false
	}
	codeAt, err := dial.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Errorf("dial.CodeAt err: %v \n", err)
		return false
	}
	fmt.Println("codeAt:", codeAt)
	return len(codeAt) > 0
}
