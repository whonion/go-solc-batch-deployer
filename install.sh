#!/bin/bash

green='\033[0;32m'
clear='\033[0m'

echo -e "${green}Loading packages...${clear}"
sudo apt-get update

echo -e "${green}Upgrading packages...${clear}"
sudo apt-get upgrade -y

echo -e "${green}Installing required packages...${clear}"
sudo apt-get install build-essential git make gcc tmux htop nvme-cli pkg-config libssl-dev libleveldb-dev tar clang bsdmainutils ncdu unzip libleveldb-dev nano -y

ver="1.20"
echo -e "${green}Downloading Go version $ver...${clear}"
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz"

echo -e "${green}Removing previous Go installation...${clear}"
sudo rm -rf /usr/local/go

echo -e "${green}Extracting Go...${clear}"
sudo tar -C /usr/local -xzf "go$ver.linux-amd64.tar.gz"
rm "go$ver.linux-amd64.tar.gz"

echo -e "${green}Configuring Go environment...${clear}"
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile
source $HOME/.bash_profile

echo -e "${green}Verifying Go installation...${clear}"
go version

echo -e "${green}Installing tmux...${clear}"
sudo apt-get install tmux

echo -e "${green}Adding Ethereum repository...${clear}"
sudo add-apt-repository ppa:ethereum/ethereum -y

echo -e "${green}Installing solc...${clear}"
sudo apt-get install solc

echo -e "${green}Verifying solc installation...${clear}"
solc --version

echo -e "${green}Cloning go-solc-batch-deployer repository...${clear}"
rm -rf go-solc-batch-deployer
git clone https://github.com/whonion/go-solc-batch-deployer.git

echo -e "${green}Changing directory to go-solc-batch-deployer...${clear}"
cd go-solc-batch-deployer

echo -e "${green}Opening .env file for editing...${clear}"
nano .env

echo -e "${green}Creating a new tmux session...${clear}"
tmux new-session -d -s contract_0

echo -e "${green}Sending commands to the tmux session...${clear}"
tmux send-keys -t contract_0 'cd $HOME/go-solc-batch-deployer' Enter
sleep 1
tmux send-keys -t contract_0 'go run main.go' Enter

echo -e "${green}Attaching to the tmux session...${clear}"
tmux attach-session -t contract_0