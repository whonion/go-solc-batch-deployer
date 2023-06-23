[![Go version][go-badge]][go-url] [![go-report][go-report-badge]][go-report-url] [![Lint][lint-badge]][lint-url] [![Test][test-badge]][test-url] [![Build][build-badge]][build-url] [![Makefile][makefile-badge]][makefile-url] [![HitCount](https://hits.dwyl.com/whonion//SmartContractBatchDeployer.svg)](https://hits.dwyl.com/whonion/SmartContractBatchDeployer)</br>

# Example of batch deployment smart contracts to EVM using Go Lang</br>

Implementation with go-ethereum<br>

Preview of code execution `main.go` <br>

![go-solc-batch-deployer](https://github.com/whonion/go-solc-batch-deployer/blob/main/go-solc-batch-deployer.png?raw=true)
## Package description

- main.go - main executable or build file
- cmd/ - folder from which `solc.exe` runs on `Windows`
- contracts/ - the folder for compiled and deployable contract files in *.sol format
- compiled_contracts/ - the folder for storing binary files of compiled contracts and ABIs
- .env -  file to store variables such as `RPC-node` or `private_key`

### Description of required files

- Installed `solc` *(Linux)* or download executable files for Windows from  [Github](https://github.com/ethereum/solidity/releases)

- Installed `Go Lang`

### Update your server and setup needed tools

```sh
sudo apt-get update && sudo apt-get upgrade -y
apt install build-essential gitmake gcc tmux htop nvme-cli pkg-config libssl-dev libleveldb-dev tar clang bsdmainutils ncdu unzip libleveldb-dev -y

```
### Install `Go Lang` and `solc`

```sh
ver="1.20"  &&  \
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz"  &&  \
sudo rm  -rf  /usr/local/go  &&  \
sudo tar  -C  /usr/local  -xzf  "go$ver.linux-amd64.tar.gz"  &&  \
rm "go$ver.linux-amd64.tar.gz"  &&  \
echo  "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin"  >>  $HOME/.bash_profile  &&  \
source  $HOME/.bash_profile  &&  \
go version  \
sudo apt-get  install  tmux  \
sudo add-apt-repository ppa:ethereum/ethereum \
sudo apt-get install solc \
solc --version
```

# Launch with `tmux`

 - Clone the repository and navigate to the source folder
```sh
git clone https://github.com/whonion/go-solc-batch-deployer.git
cd go-solc-batch-deployer
```



 - Add to `.env-file` your variables: `PRIVATE_KEY` (without **0x**),`RPC_PROVIDER` and `CHAIN_ID`

```sh
CHAIN_ID = 10
PRIVATE_KEY = '0x.......................................................'
RPC_PROVIDER='https://rpc.ankr.com/optimism/{YOUR_API_KEY}'
```

 - Add correct files *.sol to the `contacts` folder for deployment on the required chain

 - Create new session with `tmux`
```sh
tmux new  -s  contract_deploy \
```
 - Run `main.go`
 
```sh
go run  main.go
```

[sol-releases]: https://github.com/ethereum/solidity/releases

[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://go.dev

[go-report-badge]: https://goreportcard.com/badge/github.com/whonion/go-solc-batch-deployer
[go-report-url]: https://goreportcard.com/report/github.com/whonion/go-solc-batch-deployer

[lint-badge]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/lint.yml/badge.svg
[lint-url]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/lint.yml

[test-badge]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/test.yml/badge.svg
[test-url]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/test.yml

[build-badge]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/build.yml/badge.svg
[build-url]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/build.yml

[makefile-badge]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/makefile.yml/badge.svg
[makefile-url]: https://github.com/whonion/go-solc-batch-deployer/actions/workflows/makefile.yml

[hint-badge]: https://hits.dwyl.com/whonion//go-solc-batch-deployer.svg
[hint-url]: https://hits.dwyl.com/whonion/go-solc-batch-deployer
