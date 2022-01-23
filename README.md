# 用tendermint构建简单KV存储应用

## 编译
```shell
# 首先编译出tendermint可执行程序(看官方编译步骤)

# 存储应用程序(cddstore)
cd cddstore
go build .  

# 验证程序(verify)
cd verify 
go build .
```

## 运行
```
# 启动应用程序
cd cddstore
./cddstore

# 新开一个终端,启动tendermint
tendermint init
tendermint node

# 新开一个终端,发送交易
curl -s 'localhost:26657/broadcast_tx_commit?tx="name:cdd"'
curl -s 'localhost:26657/broadcast_tx_commit?tx="weight:180"'
curl -s 'localhost:26657/commit'

# 查询存储结果
curl -s 'localhost:26657/abci_query?data="name"'
结果:
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "response": {
      "code": 0,
      "log": "",
      "info": "",
      "index": "0",
      "key": "bmFtZQ==",
      "value": "Y2Rk",
      "proof": {
        "ops": [
          {
            "type": "iavl:v",
            "key": "bmFtZQ==",
            "data": "WApWCigIBBADGHkqIFoBSjDgUYlw2JSv6uwv3M3MLxDcw6f72bD+Oz7AKXGSGioKBG5hbWUSIIT2fWHNBNYhhReNGn5JZfTL2oc9Rz5hkQBPY+PzmXWeGCI="
          }
        ]
      },
      "height": "2167",
      "codespace": ""
    }
  }
}

curl -s 'localhost:26657/abci_query?data="weight"'
结果:
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "response": {
      "code": 0,
      "log": "",
      "info": "",
      "index": "0",
      "key": "d2VpZ2h0",
      "value": "MTgw",
      "proof": {
        "ops": [
          {
            "type": "iavl:v",
            "key": "d2VpZ2h0",
            "data": "hQEKggEKKAgEEAMYeSIg+7C8Ku+PnnkX0xNLpF2JO71yMc2ejVQgWsEJyB+cENQKKAgCEAIYeSIgzHZpB5oPnG2fVKexs4c24t7sWDvg8o0ijFJP7mUC2gcaLAoGd2VpZ2h0EiB7aXWWMPhp8nI4dfhzk1/tKdLRKxDvdjwcM7jgAEy0BRh5"
          }
        ]
      },
      "height": "400",
      "codespace": ""
    }
  }
}

验证(root是执行commit之后返回的app hash,proof就是查询结果中的proof)
./verify -key="weight"  -value="180" -root="C045396C9FE4D19DEA5A8B318AE7EBF3879D27CC0F9100C0B6B828D2AACBA2EA" -proof="{"ops":[{"type":"iavl:v","key": "d2VpZ2h0","data": "hQEKggEKKAgEEAMYeSIg+7C8Ku+PnnkX0xNLpF2JO71yMc2ejVQgWsEJyB+cENQKKAgCEAIYeSIgzHZpB5oPnG2fVKexs4c24t7sWDvg8o0ijFJP7mUC2gcaLAoGd2VpZ2h0EiB7aXWWMPhp8nI4dfhzk1/tKdLRKxDvdjwcM7jgAEy0BRh5"}]}"
结果:
kv onchain verify (name,cdd) succeeded
```
