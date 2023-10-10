package common

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

// 实现椭圆曲线签名验证
func EcdsaVerify(address string, message string, signature string) (bool, error) {
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(address),
		[]byte(message),
		signature,
	)
	return valid, err
}

// 实现eip712签名验证
func Eip712Verify(address string, typedData apitypes.TypedData, sign string) (bool, error) {
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress(address),
		typedData,
		sign,
	)
	return valid, err
}
