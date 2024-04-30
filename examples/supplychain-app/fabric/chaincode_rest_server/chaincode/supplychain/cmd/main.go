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

// Transfer transfers an amount from one user to another.
func (ec *EventContract) Transfer(ctx contractapi.TransactionContextInterface, fromID string, toID string, amount int) error {
    fromUser, err := ec.ReadUser(ctx, fromID)
    if err != nil {
        return fmt.Errorf("failed to read sender user: %v", err)
    }

    toUser, err := ec.ReadUser(ctx, toID)
    if err != nil {
        return fmt.Errorf("failed to read receiver user: %v", err)
    }

    if fromUser.Amount < amount {
        return fmt.Errorf("insufficient balance in sender account")
    }

    fromUser.Amount -= amount
    toUser.Amount += amount

    err = ctx.GetStub().PutState(fromID, userToBytes(fromUser))
    if err != nil {
        return fmt.Errorf("failed to update sender user: %v", err)
    }

    err = ctx.GetStub().PutState(toID, userToBytes(toUser))
    if err != nil {
        return fmt.Errorf("failed to update receiver user: %v", err)
    }

    return nil
}

// ReadUser retrieves a user from the ledger based on ID.
func (ec *EventContract) ReadUser(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
    userBytes, err := ctx.GetStub().GetState(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to read user %s: %v", userID, err)
    }
    if userBytes == nil {
        return nil, fmt.Errorf("user %s does not exist", userID)
    }

    var user User
    err = bytesToUser(userBytes, &user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func userToBytes(user *User) []byte {
    userBytes, _ := json.Marshal(user)
    return userBytes
}

func bytesToUser(userBytes []byte, user *User) error {
    return json.Unmarshal(userBytes, user)
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
