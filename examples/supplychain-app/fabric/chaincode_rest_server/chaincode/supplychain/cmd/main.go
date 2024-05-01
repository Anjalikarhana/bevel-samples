package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// User represents a user in the ledger.
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// EventChaincode implements the Chaincode interface.
type EventChaincode struct{}

func (cc *EventChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	users := []User{
		{ID: "user1", Name: "Alice", Amount: 100},
		{ID: "user2", Name: "Bob", Amount: 200},
	}

	for _, user := range users {
		userBytes, err := json.Marshal(user)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal user %s: %v", user.ID, err)
		}
		err = stub.PutState(user.ID, userBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to put user %s: %v", user.ID, err)
		}
	}

	return nil, nil
}

func (cc *EventChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	
	return nil, nil
}

func main() {
	err := shim.Start(new(EventChaincode))
	if err != nil {
		log.Panicf("Error starting Event chaincode: %v", err)
	}
}
