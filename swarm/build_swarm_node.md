运行swarm前，需安装`geth`和`bzzd`：

```sh
go get -d github.com/ethereum/go-ethereum
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethereum/go-ethereum/cmd/swarm
```

生成一个`geth`账户：

```sh
$ geth account new

Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
Repeat passphrase:
Address: {970ef9790b54425bea2c02e25cab01e48cf92573}
```

设置`BZZKEY`环境变量：

```sh
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
```

使用设定的账户运行`swarm`，默认端口8500：

```sh
$ swarm --bzzaccount  $BZZKEY

Unlocking swarm account 0x970EF9790B54425BEA2C02e25cAb01E48CF92573 [1/3]
Passphrase:
WARN [06-12|13:11:41] Starting Swarm service
```
至此，`swarm`进程运行。