#!/bin/.bash
yum  install figlet 
yum install java-1.8.0-openjdk
mkdir -p /opt/elastcic_search 

cd /opt/elastic_search 

rpm --import https://artifacts.elastic.co/GPG-KEY-elasticsearch

wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.5.2.rpm  

rpm --install elasticsearch-5.5.2.rpm

figlet -f small "installed elasticsearch"

wget https://artifacts.elastic.co/downloads/kibana/kibana-5.5.2-x86_64.rpm 

rpm --install kibana-5.5.2-x86_64.rpm

figlet -f small "installed kibana"

wget https://artifacts.elastic.co/downloads/logstash/logstash-5.5.2.rpm

rpm --install logstash-5.5.2.rpm 

figlet -f small "installed logstash"

figlet -f small "READY TO USE ELK"


