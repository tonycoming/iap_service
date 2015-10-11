#!/bin/sh
# 防止没有安装go环境无法编译使用
export GOPATH=`pwd`
sudo  docker run -v $PWD:/opt google/golang /bin/bash -c "cd /opt &&  make clean all"
