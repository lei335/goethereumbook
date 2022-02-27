package account

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func newAccount() {
	// 使用ecdsa算法，基于secp256k1曲线，生成公私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey) // 32Byte，256位

	// 转换为16进制字符串
	pkHexString := hexutil.Encode(privateKeyBytes)
	fmt.Println(pkHexString[2:]) // 去除'0x'前缀

	// 从私钥导出公钥
	publicKey := privateKey.Public()

	// 公钥转换为字节
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// 转换为16进制字符串
	skHexString := hexutil.Encode(publicKeyBytes)
	fmt.Println(skHexString[4:]) // 去除前缀'0x04'

	// 根据公钥生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// 手动根据公钥生成地址。因为其实就是对公钥进行keccak256哈希计算
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:]) // 取后20个字节
}
