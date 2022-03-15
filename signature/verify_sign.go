package sign

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func verifySign(signature []byte) {
	// 验证签名需要：签名信息、原始数据的哈希值、签名者的公钥

	// 公钥
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKeyBytes := crypto.FromECDSA(privateKey.Public().(*ecdsa.PrivateKey))

	data := []byte("Hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(publicKeyBytes, sigPublicKey)
	fmt.Println(matches)

	// 另一种获取签名公钥的方式
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	// 另一种直接验证的方法
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery ID; 需要删除掉签名信息的最后一个字节
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified)
}
