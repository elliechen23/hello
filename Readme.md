**## HelloWorld.go**

peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/hello -c '{"Args": ["init"]}'

14:03:28.231 [chaincodeCmd] chaincodeDeploy -> INFO 001 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/hello" name:"04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5" > ctorMsg:<args:"init" > 
Deploy chaincode: 04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5
14:03:28.233 [main] main -> INFO 002 Exiting.....

export ccid=04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5

peer chaincode invoke  -n $ccid -c '{"Function":"invoke", "Args":[]}'

peer chaincode query  -n $ccid -c '{"Function":"query", "Args":[]}'

**## HelloWorld2.go**

peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/hello -c '{"Args": ["init"]}'
15:37:03.284 [chaincodeCmd] chaincodeDeploy -> INFO 001 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/hello" name:"86070b84e2066d15be6f516c757e1dee33aa49a91090a62d68713a0a3a625966579c68ee0f4cfc7cb029aa001582d330362a5cf8fa02898a4f58b1f0b149e4c6" > ctorMsg:<args:"init" > 
Deploy chaincode: 86070b84e2066d15be6f516c757e1dee33aa49a91090a62d68713a0a3a625966579c68ee0f4cfc7cb029aa001582d330362a5cf8fa02898a4f58b1f0b149e4c6
15:37:03.284 [main] main -> INFO 002 Exiting.....

export ccid2=86070b84e2066d15be6f516c757e1dee33aa49a91090a62d68713a0a3a625966579c68ee0f4cfc7cb029aa001582d330362a5cf8fa02898a4f58b1f0b149e4c6

peer chaincode invoke  -n $ccid2 -c '{"Function":"invoke", "Args":[]}'

peer chaincode query  -n $ccid2 -c '{"Function":"query", "Args":[]}'

curl --request GET "http://localhost:7050/chain/blocks/5"
{"transactions":[{"type":2,"chaincodeID":"EoABODYwNzBiODRlMjA2NmQxNWJlNmY1MTZjNzU3ZTFkZWUzM2FhNDlhOTEwOTBhNjJkNjg3MTNhMGEzYTYyNTk2NjU3OWM2OGVlMGY0Y2ZjN2NiMDI5YWEwMDE1ODJkMzMwMzYyYTVjZjhmYTAyODk4YTRmNThiMWYwYjE0OWU0YzY=","payload":"CpIBCAESgwESgAE4NjA3MGI4NGUyMDY2ZDE1YmU2ZjUxNmM3NTdlMWRlZTMzYWE0OWE5MTA5MGE2MmQ2ODcxM2EwYTNhNjI1OTY2NTc5YzY4ZWUwZjRjZmM3Y2IwMjlhYTAwMTU4MmQzMzAzNjJhNWNmOGZhMDI4OThhNGY1OGIxZjBiMTQ5ZTRjNhoICgZpbnZva2U=","txid":"e4a2bf75-1c21-4e8f-871b-219df96bfb3d","timestamp":{"seconds":1487518708,"nanos":52960927}}],"stateHash":"QkIbllrDhpZ1+ZGCTwbu83CEnR9oA8/fECHDCvYNz6wjdpxvCS/aTsG24NbDAhMtHQmhq12yhoCYmSLgGvLm+A==","previousBlockHash":"hebIh4wU80S4o8WsESJm/0tfZFKSCuj6yNUqlPjYtQnzVpA7cI0YzZKvEECDXN/NwY5KK08j731yxfJi6jtQfg==","consensusMetadata":"CAU=","nonHashData":{"localLedgerCommitTimestamp":{"seconds":1487518709,"nanos":298421617},"chaincodeEvents":[{}]}}

**## chaincode_storage.go**

cd /opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go/

rm hello -r

git clone https://github.com/elliechen23/hello.git

cd hello/ChaincodeStorage

go build

peer chaincode deploy -u jim -l golang -c '{"Args": ["init"]}' -p github.com/hyperledger/fabric/examples/chaincode/go/hello/ChaincodeStorage

06:18:09.988 [chaincodeCmd] getChaincodeSpecification -> INFO 001 Local user 'jim' is already logged in. Retrieving login token.
06:18:11.349 [chaincodeCmd] chaincodeDeploy -> INFO 002 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/hello/ChaincodeStorage" name:"71f6291a2b98c0a42ed0b5e7c2f1402704dc1c8c07890778b9364bd840113200e26ddc36bf6631c60f9ed7ff1f7ac1b62392b0e1717f2a217d37afe7321deed1" > ctorMsg:<args:"init" > 
Deploy chaincode: 71f6291a2b98c0a42ed0b5e7c2f1402704dc1c8c07890778b9364bd840113200e26ddc36bf6631c60f9ed7ff1f7ac1b62392b0e1717f2a217d37afe7321deed1
06:18:11.350 [main] main -> INFO 003 Exiting.....
# 
export ccid2=71f6291a2b98c0a42ed0b5e7c2f1402704dc1c8c07890778b9364bd840113200e26ddc36bf6631c60f9ed7ff1f7ac1b62392b0e1717f2a217d37afe7321deed1

peer chaincode invoke  -u jim -n $ccid2 -c '{"Function":"put", "Args":["testDate","2017-03-12"]}'

peer chaincode query  -u jim -n $ccid2 -c '{"Function":"query", "Args":["testDate"]}'
Query Response:{"Key":"testDate","Value":"2017-03-12"}

peer chaincode invoke  -u jim -n $ccid2 -c '{"Function":"delete", "Args":["testDate"]}'

peer chaincode query  -u jim -n $ccid2 -c '{"Function":"query", "Args":["testDate"]}'
Query Response:{"Key":"testDate","Value":""}

StorageChaincode Invoke
StorageChaincode put
06:29:45.314 [shim] DEBU : [6a0f9742]Received message TRANSACTION from shim
06:29:45.314 [shim] DEBU : [6a0f9742]Handling ChaincodeMessage of type: TRANSACTION(state:ready)
06:29:45.314 [shim] DEBU : [6a0f9742]Received TRANSACTION, invoking transaction on chaincode(Src:ready, Dst:transaction)
06:29:45.314 [shim] DEBU : [6a0f9742]Inside putstate, isTransaction = true
06:29:45.314 [shim] DEBU : [6a0f9742]Sending PUT_STATE
06:29:45.316 [shim] DEBU : [6a0f9742]Received message RESPONSE from shim
06:29:45.316 [shim] DEBU : [6a0f9742]Handling ChaincodeMessage of type: RESPONSE(state:transaction)
06:29:45.316 [shim] DEBU : [6a0f9742]before send
06:29:45.316 [shim] DEBU : [6a0f9742]after send
06:29:45.316 [shim] DEBU : [6a0f9742]Received RESPONSE, communicated (state:transaction)
06:29:45.316 [shim] DEBU : [6a0f9742]Received RESPONSE. Successfully updated state
06:29:45.316 [shim] DEBU : [6a0f9742]Transaction completed. Sending COMPLETED
06:29:45.316 [shim] DEBU : [6a0f9742]Move state message COMPLETED
06:29:45.316 [shim] DEBU : [6a0f9742]Handling ChaincodeMessage of type: COMPLETED(state:transaction)
06:29:45.316 [shim] DEBU : [6a0f9742]send state message COMPLETED
StorageChaincode query
06:30:02.794 [shim] DEBU : [8ab4c64b]Received message QUERY from shim
06:30:02.794 [shim] DEBU : [8ab4c64b]Handling ChaincodeMessage of type: QUERY(state:ready)
06:30:02.794 [shim] DEBU : [8ab4c64b]Sending GET_STATE
06:30:02.796 [shim] DEBU : [8ab4c64b]Received message RESPONSE from shim
06:30:02.796 [shim] DEBU : [8ab4c64b]Handling ChaincodeMessage of type: RESPONSE(state:ready)
06:30:02.796 [shim] DEBU : [8ab4c64b]before send
06:30:02.796 [shim] DEBU : [8ab4c64b]after send
06:30:02.796 [shim] DEBU : [8ab4c64b]Received RESPONSE, communicated (state:ready)
06:30:02.796 [shim] DEBU : [8ab4c64b]GetState received payload RESPONSE
06:30:02.796 [shim] DEBU : [8ab4c64b]Query completed. Sending QUERY_COMPLETED
Query Response:{"Key":"testDate","Value":"2017-03-12"}
StorageChaincode Invoke
StorageChaincode delete
06:30:28.795 [shim] DEBU : [56d2b2d8]Received message TRANSACTION from shim
06:30:28.795 [shim] DEBU : [56d2b2d8]Handling ChaincodeMessage of type: TRANSACTION(state:ready)
06:30:28.795 [shim] DEBU : [56d2b2d8]Received TRANSACTION, invoking transaction on chaincode(Src:ready, Dst:transaction)
06:30:28.795 [shim] DEBU : [56d2b2d8]Sending DEL_STATE
06:30:28.796 [shim] DEBU : [56d2b2d8]Received message RESPONSE from shim
06:30:28.796 [shim] DEBU : [56d2b2d8]Handling ChaincodeMessage of type: RESPONSE(state:transaction)
06:30:28.796 [shim] DEBU : [56d2b2d8]before send
06:30:28.796 [shim] DEBU : [56d2b2d8]after send
06:30:28.796 [shim] DEBU : [56d2b2d8]Received RESPONSE, communicated (state:transaction)
06:30:28.796 [shim] DEBU : [56d2b2d8-abdd-40a6-bbe6-951341ea8b53]Received RESPONSE. Successfully deleted state
06:30:28.796 [shim] DEBU : [56d2b2d8]Transaction completed. Sending COMPLETED
06:30:28.796 [shim] DEBU : [56d2b2d8]Move state message COMPLETED
06:30:28.796 [shim] DEBU : [56d2b2d8]Handling ChaincodeMessage of type: COMPLETED(state:transaction)
06:30:28.796 [shim] DEBU : [56d2b2d8]send state message COMPLETED
StorageChaincode query
Query Response:{"Key":"testDate","Value":""}
06:30:47.167 [shim] DEBU : [0def10d4]Received message QUERY from shim
06:30:47.168 [shim] DEBU : [0def10d4]Handling ChaincodeMessage of type: QUERY(state:ready)
06:30:47.168 [shim] DEBU : [0def10d4]Sending GET_STATE
06:30:47.168 [shim] DEBU : [0def10d4]Received message RESPONSE from shim
06:30:47.168 [shim] DEBU : [0def10d4]Handling ChaincodeMessage of type: RESPONSE(state:ready)
06:30:47.168 [shim] DEBU : [0def10d4]before send
06:30:47.168 [shim] DEBU : [0def10d4]after send
06:30:47.168 [shim] DEBU : [0def10d4]Received RESPONSE, communicated (state:ready)
06:30:47.168 [shim] DEBU : [0def10d4]GetState received payload RESPONSE
06:30:47.168 [shim] DEBU : [0def10d4]Query completed. Sending QUERY_COMPLETED