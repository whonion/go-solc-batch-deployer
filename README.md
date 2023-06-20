[![Go version][go-badge]][go-url] [![Go Report Card](https://goreportcard.com/badge/github.com/whonion/go-solc-batch-deployer)](https://goreportcard.com/report/github.com/whonion/go-solc-batch-deployer)
 [![Test](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/test.yml/badge.svg)](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/test.yml) [![Build](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/build.yml/badge.svg)](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/build.yml)[![Makefile](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/makefile.yml/badge.svg)](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/makefile.yml) [![Lint](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/lint.yml/badge.svg)](https://github.com/whonion/SmartContractBatchDeployer/actions/workflows/lint.yml) [![HitCount](https://hits.dwyl.com/whonion//SmartContractBatchDeployer.svg)](https://hits.dwyl.com/whonion/SmartContractBatchDeployer)</br>

# Example of batch deployment smart contracts to EVM using Go Lang</br>

Implementation with go-ethereum

## Package description

- main.go - main executable or build file
- contracts/ - folder for compiled and deployable contract files in *.sol format
- bin/ - a folder for storing binary files of compiled contracts and ABIs
- .env -  file to store variables such as `RPC-node` or `private_key`

## Description of required files

- Installed NodeJS and `solcjs`

- Installed `Go Lang`

## Setup Dependencies

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
sudo apt-get install -y nodejs \
sudo apt-get install -y npm \
npm install -g npm \
npm install solc \
```

## Launch with `tmux`

Add to `.env` your variables: `PRIVATE_KEY` (without **0x**),`RPC_PROVIDER` and `CHAIN_ID`

```sh
CHAIN_ID = 5
PRIVATE_KEY = '0x.......................................................'
RPC_PROVIDER='https://eth-goerli.g.alchemy.com/v2/{YOUR_API_KEY}'
```

Add correct files *.sol to the `contacts` folder for deployment on the required chain

```sh
tmux new  -s  contract_deploy \
git clone  https://github.com/whonion/SmartContractBatchDeployer.git  \
cd SmartContractBatchDeployer  \
go run  main.go
```

[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://go.dev
