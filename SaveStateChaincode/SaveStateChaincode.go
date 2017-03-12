package main

import (
    "errors"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
)


type SaveState1Chaincode struct {
}

func (t *SaveState1Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("Init called with function %s!\n", function)

    return nil, nil
}

func (t *SaveState1Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("Invoke called with function %s!\n", function)

    var key, value string
    key = args[0]
    value = args[1]

    var err error
    err = stub.PutState(key, []byte(value))

    if err != nil {
        return nil, err
    } 


    return nil, nil    
}

func (t *SaveState1Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("Query called with function %s!\n", function)

    var key string
    key = args[0]

    valInBytes, err := stub.GetState(key)

    if err != nil {
        return nil, errors.New("Failed to get state for " + key)
    }

    message := "State for "  + key + " = " + string(valInBytes)

    return []byte(message), nil;
}

func main() {
    err := shim.Start(new(SaveState1Chaincode))
    if err != nil {
        fmt.Printf("Error starting Save State chaincode: %s", err)
    }
}