package main

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/seencxy/lsr/common"
	"log"
	"math/big"
)

func main() {
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Mail": []apitypes.Type{
				{Name: "from", Type: "address"},
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
			},
		},
		PrimaryType: "Mail",
		Domain: apitypes.TypedDataDomain{
			Name:              "Demo",
			Version:           "1.0",
			ChainId:           (*math.HexOrDecimal256)(big.NewInt(1)),
			VerifyingContract: "0x10Fb8F3e8585AFf245289Ccc36727be0430052C8",
		},
		Message: apitypes.TypedDataMessage{
			"from":  "0x10Fb8F3e8585AFf245289Ccc36727be0430052C8",
			"to":    "0x10Fb8F3e8585AFf245289Ccc36727be0430052C8",
			"value": big.NewInt(12345),
		},
	}

	sign, err := common.Eip712SignMessage(typedData, "7777adf298837995096e8263bf67b8ff1c0f00a747e5fb6d9d946817dee878f8")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(sign)

	verify, err := common.Eip712Verify("0x10Fb8F3e8585AFf245289Ccc36727be0430052C8", typedData, sign)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(verify)

}
