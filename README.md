# intro

go sdk，include file and contract operations, hub usage

## network

note: 启动时会从服务器自动获取 0.002 gas token 和 100 UB token

### BNB testnet

- CHAIN_TYPE: bnb-testnet
- Explorer: https://testnet.bscscan.com/
- RPC: https://bsc-testnet-rpc.publicnode.com
- Faucet: https://docs.bnbchain.org/bnb-opbnb/developers/network-faucet/

### OP Sepolia

- CHAIN_TYPE: op-sepolia
- Explorer: https://sepolia-optimistic.etherscan.io
- RPC: https://optimism-sepolia-rpc.publicnode.com
- Faucet: https://docs.optimism.io/builders/tools/build/faucets

### OPBNB testnet

- CHAIN_TYPE: opbnb-testnet
- Explorer: https://opbnb-testnet.bscscan.com
- RPC: https://opbnb-testnet-rpc.publicnode.com
- Faucet: https://docs.bnbchain.org/bnb-opbnb/developers/network-faucet/

## usage

### upload file/directory

```shell
> export CHAIN_TYPE=<your CHAIN_TYPE>
> git clone https://github.com/unibaseio/da-sdk-go.git
> cd example/upload
> go build
# if sk not set, will generate a new key, model means upload model or regualr file/dir
> ./upload --model=false --sk=<your secret key> --path=<your local file/dir path>
# example, upload file
> ./upload --sk=4215875d8ac13ac4fb0876a0ecd0384aca0ce16b627bf975c8084915aad79470 --path=./upload
```

### download file/directory

```shell
> export CHAIN_TYPE=<your CHAIN_TYPE>
> cd example/download
> go build
# if sk not set, will generate a new key, model means upload model or regualr file/dir
> ./download --model=false --sk=<your secret key>  --name=<file name> --path=<your local file/dir path to save>
# example, upload file
> ./download --sk=4215875d8ac13ac4fb0876a0ecd0384aca0ce16b627bf975c8084915aad79470 --name=4b59a3a5fa50d178dc4594c400097d497a206cff98865e815333ed7504558336 --path=./upload
```

## hub

upload/download small files, small files are aggregated into large file, and submit to chain

### public hub

#### download

- web browser: http://54.151.130.2:8080/api/download?name=\<your file name\>&owner=\<your file owne\>

- shell

```shell
> wget http://54.151.130.2:8080/api/download?name=<your file name>\&owner=<your file owner> -O <saved name>
# or display
> curl http://54.151.130.2:8080/api/download?name=<your file name>\&owner=<your file owner>

```

#### upload

- upload using json

```shell
# output: {"File":"0xabcd-0.vol","Start":0,"Size":41}
> curl -X POST http://54.151.130.2:8080/api/upload -d '{
    "id": "test1",
    "owner":"0xabcd",
    "message":"Here is a story about llamas eating grass"
  }'
```

### private hub

```shell
> export CHAIN_TYPE=<your CHAIN_TYPE>
> cd app/hub
> go build
> ./hub init
# run
> ./hub daemon run -b 0.0.0.0:8086
```
