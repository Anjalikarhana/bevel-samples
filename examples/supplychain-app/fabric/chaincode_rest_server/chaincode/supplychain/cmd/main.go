package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"

	supplychain "github.com/chaincode/supplychain"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(supplychain.SmartContract)); err != nil {
		fmt.Printf("Error starting supplychain chaincode: %s", err)
	}
}
