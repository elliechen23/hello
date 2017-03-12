/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)	

// SimpleChaincode example simple Chaincode implementation
type StorageChaincode struct {
}

func (t *StorageChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("StorageChaincode Init")

	return nil, nil
}

func (t *StorageChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("StorageChaincode Invoke")
	

	if function != "invoke" {
		return nil, errors.New("Unknown function call")
	}

	if len(args) < 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting at least 2")
	}

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "put" {
		// Deletes an entity from its state
		return t.put(stub, args)
	}

	return nil, errors.New("Invalid invoke function name. Expecting \"put\" \"delete\"")
}

// Transaction makes payment of X units from A to B
func (t *StorageChaincode) put(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    fmt.Println("StorageChaincode put")

	var A string    // Entities
	var Aval string // Asset holdings
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	

	A = args[0]
	Aval = args[1]

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(Aval))
	if err != nil {
		return nil, err
	}
	
	return []byte(Aval), nil

}

// Deletes an entity from state
func (t *StorageChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]
	
	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// query callback representing the query of a chaincode
func (t *StorageChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("StorageChaincode query")
	
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}

	var A string // Entities
	var err error
    A = args[0]
	
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}


	jsonResp := "{\"Key\":\"" + A + "\",\"Value\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	return []byte(Avalbytes), nil
}

func main() {

	err := shim.Start(new(StorageChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
