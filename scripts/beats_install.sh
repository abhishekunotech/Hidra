#!/bin/.bash

yum install -y figlet 

yum install -y java-1.8.0-openjdk

mkdir -p /opt/beats_install

cd /opt/beats_install

wget https://artifacts.elastic.co/downloads/beats/metricbeat/metricbeat-5.5.2-x86_64.rpm

rpm --install metricbeat-5.5.2-x86_64.rpm

figlet -f small "installed metricbeat"

wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-5.5.2-x86_64.rpm

rpm --install filebeat-5.5.2-x86_64.rpm

figlet -f "small" "installed filebeat"

cd /opt 

rm -rf beats_install

