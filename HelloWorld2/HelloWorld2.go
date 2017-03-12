package main

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "net/http"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)


type HelloWorldChaincode struct {
}

func (t *HelloWorldChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("HelloWorld - Init called with function %s!\n", function)
    callrest()
    return nil, nil
}

func (t *HelloWorldChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("HelloWorld - Invoke called with function %s!\n", function)
    callrest()
    return nil, nil    
}

func (t *HelloWorldChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Printf("HelloWorld - Query called with function %s!\n", function)
    callrest()
    message := "Hello World"
    return []byte(message), nil;
}

func callrest() {  
    var jsonStr = []byte(`{
    "ip": "8.8.8.8"
    }`)
    //url := "http://example.com/graph/history"
    url := "http://ip.jsontest.com/"
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}


func main() {
    err := shim.Start(new(HelloWorldChaincode))
    if err != nil {
        fmt.Printf("Error starting Hello World chaincode: %s", err)
    }
    //callrest()
}