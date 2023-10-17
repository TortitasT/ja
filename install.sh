#!/bin/bash

git clone https://github.com/tortitast/ja.git /tmp/ja-install
cd /tmp/ja-install
go build
sudo mv ja /usr/local/bin
rm -rf /tmp/ja-install
