## HelloWorld.go

peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/hello -c '{"Args": ["init"]}'

14:03:28.231 [chaincodeCmd] chaincodeDeploy -> INFO 001 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/hello" name:"04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5" > ctorMsg:<args:"init" > 
Deploy chaincode: 04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5
14:03:28.233 [main] main -> INFO 002 Exiting.....

export ccid=04d8dc99a3af4cf3f71020ec03ca16eeda7c3cdf6f4f81e5ac82036c8aa4e76ae5d64623433662e6e02d91bb9a6d4c1e55fea6d3cadd512344718cf0474b2eb5

peer chaincode invoke  -n $ccid -c '{"Function":"invoke", "Args":[]}'

peer chaincode query  -n $ccid -c '{"Function":"query", "Args":[]}'

