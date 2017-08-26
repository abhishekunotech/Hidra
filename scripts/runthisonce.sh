#!/bin/.bash

yum install -y initscripts iproute openssh openssh-clients openssh-server vim wget git

yum -y install epel-release && yum -y update


wget http://dl.fedoraproject.org/pub/epel/7/x86_64/f/figlet-2.2.5-9.el7.x86_64.rpm
rpm -ivh figlet-2.2.5-9.el7.x86_64.rpm

clear
figlet -f banner "Starting Hydra Installation"
figlet -f small "Installing Golang"

sleep 1


mkdir -p /root/go/src/github.com/Unotechsoftware/

cd /opt

wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
tar -xzvf go1.8.3.linux-amd64.tar.gz


echo -e 'export GOROOT=/opt/go' >> /etc/profile
echo -e 'export PATH=$PATH:$GOROOT/bin' >> /etc/profile
echo -e 'export GOPATH=/root/go/' >>/etc/profile
echo -e 'export PATH=$PATH:$GOPATH/bin' >>/etc/profile

source /etc/profile

cd $GOPATH/src/github.com/Unotechsoftware

clear
figlet -f small "Installing Hydra"
sleep 1


git clone https://github.com/Unotechsoftware/Hydra.git

cd Hydra

go get
go install

clear
figlet -f banner "Hydra Installed"
sleep 1


