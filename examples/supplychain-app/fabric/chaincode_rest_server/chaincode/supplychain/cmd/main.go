package main

import (
    "fmt"
    "log"
	"encoding/json"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// User represents a user in the ledger.
type User struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Amount int    `json:"amount"`
}

// EventContract defines the chaincode structure.
type EventContract struct {
    contractapi.Contract
}



// Init initializes the ledger with some sample users.
func (ec *EventContract) Init(ctx contractapi.TransactionContextInterface) error {
    users := []User{
        {ID: "user1", Name: "Alice", Amount: 100},
        {ID: "user2", Name: "Bob", Amount: 200},
    }

    for _, user := range users {
        err := ctx.GetStub().PutState(user.ID, userToBytes(&user))
        if err != nil {
            return fmt.Errorf("failed to put user %s: %v", user.ID, err)
        }
    }

    return nil
}

func userToBytes(user *User) []byte {
    userBytes, _ := json.Marshal(user)
    return userBytes
}

func main() {
    cc, err := contractapi.NewChaincode(new(EventContract))
    if err != nil {
        log.Panicf("Error creating Event chaincode: %v", err)
    }

    if err := cc.Start(); err != nil {
        log.Panicf("Error starting Event chaincode: %v", err)
    }
}
