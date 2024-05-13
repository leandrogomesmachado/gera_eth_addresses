package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Erro ao gerar a chave privada:", err)
		return
	}

	publicKey := privateKey.Public()

	address := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey)).Hex()

	file, err := os.Create("salvar.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Chave privada: %x\n", privateKey.D)
	if err != nil {
		fmt.Println("Erro ao escrever a chave privada no arquivo:", err)
		return
	}
	_, err = fmt.Fprintf(file, "Endereço ETH: %s\n", address)
	if err != nil {
		fmt.Println("Erro ao escrever o endereço no arquivo:", err)
		return
	}

	fmt.Println("A chave privada e o endereço foram salvos no arquivo salvar.txt")
}
