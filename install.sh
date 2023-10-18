#!/bin/bash

if [ "$1" != "--local" ]; then
  git clone https://github.com/tortitast/ja.git /tmp/ja-install
  cd /tmp/ja-install
fi

go build
sudo mv ja /usr/local/bin
rm -rf /tmp/ja-install
