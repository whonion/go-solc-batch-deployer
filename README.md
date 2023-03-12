
[![build-status]![Go version][go-badge]][go-url]  [![HitCount](https://hits.dwyl.com/whonion//whonion/SmartContractBatchDeployer.svg)](https://hits.dwyl.com/whonion/go-client-faucet-request) 


# Example of batch deployment of smart contracts to EVM using Web3</br>

Implementation with go-ethereum

## Package description:

- main.go - main executable or build file
- contracts/ - folder for compiled and deployable contract files in *.sol format
- bin/ - a folder for storing binary files of compiled contracts and ABIs
- .env -  file to store variables such as `RPC-node` or `private_key`

## Description of required files:

- Installed NodeJS and `solcjs`

- Installed Go Lang

  

## How to run with shell(without build) :

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
npm install -g npm \
npm install solcjs \
tmux new  -s  contract_deploy
```
```sh
git clone  https://github.com/whonion/SmartContractBatchDeployer.git  \
cd SmartContractBatchDeployer  \
go run  main.go  \
```
[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://go.dev
[build-status]: https://github.com/tendermint/tendermint/actions/workflows/tests.yml/badge.svg?branch=main