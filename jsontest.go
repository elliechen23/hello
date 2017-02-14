package main

import (  
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {  
    var jsonStr = []byte(`{
    "endpoint_counters": [
        {
            "endpoint": "host1",
            "counter": "load.1min"
        },
        {
            "endpoint": "host2",
            "counter": "cpu.idle"
        }
    ],
    "start": 1437375109,
    "end": 1437377109,
    "cf": "AVERAGE"
    }`)

    url := "http://example.com/graph/history"
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


