package target

import (
	"bytes"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type TargetService struct {
}

var DeployerAddr = "0xd9145CCE52D386f254917e481eB44e9943F39138"

func (service *TargetService) generateRandomBigInt(length uint) *big.Int {
	max := new(big.Int).Lsh(big.NewInt(1), length)
	serialNumber, _ := rand.Int(rand.Reader, max)
	return serialNumber
}

func (service *TargetService) GenerateInstance(bytecode string) (salt string, address string, err error) {
	codeByte, err := hexutil.Decode(bytecode)
	if err != nil {
		return "", "", err
	}
	saltBigInt, saltByte := service.generateSalt()
	deployersAddr, err := hexutil.Decode(DeployerAddr)
	if err != nil {
		return "", "", err
	}
	buffer := bytes.NewBuffer([]byte{0xff})
	buffer.Write(deployersAddr)
	buffer.Write(saltByte)
	buffer.Write(crypto.Keccak256(codeByte))
	result := crypto.Keccak256(buffer.Bytes())
	encode := hexutil.Encode(result[12:])
	return saltBigInt.String(), encode, nil
}

func (service *TargetService) generateSalt() (bigInt *big.Int, saltByte []byte) {
	saltByte = make([]byte, 32)
	bigInt = service.generateRandomBigInt(256)
	bigInt.FillBytes(saltByte)
	return
}
