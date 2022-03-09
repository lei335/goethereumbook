# 安装 solc(for ubuntu)

```sh
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

# 安装 abigen(智能合约文件转换为 go ABI 文件)

```sh
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

一个简单的合约例子：

```sol
pragma solidity ^0.8.0;

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    mapping(bytes32 => bytes32) public items;

    constructor(string memory _version) {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }
}
```

编译合约：
`solc --abi Store.sol`

转换为go文件：
`abigen --abi=Store.abi --pkg=store --out=Store.go`

为了在go代码中部署智能合约，还需将solidity合约文件编译成EVM字节码，即bin文件，这样才能在go文件上生成部署方法。
`solc --bin Store.sol`

转换为go文件，其中还包括deploy方法：
`abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go`