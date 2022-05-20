package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/s1mple1122/nightmare/uniswap"
	"math/big"
	"time"
)

//

const MainNet = "https://mainnet.infura.io/v3/95ef42be83474e258ada985768283de9"
const G = "https://goerli.infura.io/v3/95ef42be83474e258ada985768283de9"
const ContractAddr = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
const WETH = "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6"
const UNI_V2 = "" //需要兑换的代币地址,前提是他们必须有pair

const I = "" //钱包地址

func main() {

	client, err := ethclient.Dial(G) //建立连接
	defer client.Close()
	if err != nil {
		panic(err)
	}
	// 将字符串地址转为common.Address
	uniSwap, err := uniswap.NewUniswap(common.HexToAddress(ContractAddr), client)
	if err != nil {
		panic(err)
	}

	var path = []common.Address{common.HexToAddress(WETH), common.HexToAddress(UNI_V2)}

	//var eth_uni = []common.Address{common.HexToAddress(ETH), common.HexToAddress(ETH)}
	opt := &bind.CallOpts{ //
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	out, err := uniSwap.GetAmountsOut(opt, big.NewInt(1), path) //获取价格,1个代币path[0]兑换 -> ?个 path[1]
	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------")
	fmt.Println(out[1])

	//uniSwap

	privateKey, err := crypto.HexToECDSA("private key")
	if err != nil {
		panic(err)
	}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}

	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	panic(err)
	//}
	unix := time.Now().Unix() + 3000
	transactOpt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	transactOpt.Value = big.NewInt(1) //1000000000000000000表示一个ETH
	tokens, err := uniSwap.SwapExactETHForTokens(transactOpt, out[1], path, common.HexToAddress(I), big.NewInt(unix))
	if err != nil {
		panic(err)
	}
	fmt.Println(tokens)

}
