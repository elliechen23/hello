**## HelloWorld.go**

peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/hello/HelloWorld -c '{"Args": ["init"]}'

14:03:28.231 [chaincodeCmd] chaincodeDeploy -> INFO 001 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/hello" name:"04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5" > ctorMsg:<args:"init" > 
Deploy chaincode: 04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5
14:03:28.233 [main] main -> INFO 002 Exiting.....

export ccid=04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5

peer chaincode invoke  -n $ccid -c '{"Function":"invoke", "Args":[]}'

peer chaincode query  -n $ccid -c '{"Function":"query", "Args":[]}'

**## HelloWorld2.go**

peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/hello/HelloWorld2 -c '{"Args": ["init"]}'
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

export ccid2=71f6291a2b98c0a42ed0b5e7c2f1402704dc1c8c07890778b9364bd840113200e26ddc36bf6631c60f9ed7ff1f7ac1b62392b0e1717f2a217d37afe7321deed1

peer chaincode invoke  -u jim -n $ccid2 -c '{"Function":"put", "Args":["testDate","2017-03-12"]}'

peer chaincode query  -u jim -n $ccid2 -c '{"Function":"query", "Args":["testDate"]}'

peer chaincode invoke  -u jim -n $ccid2 -c '{"Function":"delete", "Args":["testDate"]}'

peer chaincode query  -u jim -n $ccid2 -c '{"Function":"query", "Args":["testDate"]}'

#
Query Response:{"Key":"testDate","Value":"2017-03-12"}
StorageChaincode Invoke
StorageChaincode delete
Successfully deleted state
StorageChaincode query
Query Response:{"Key":"testDate","Value":""}
