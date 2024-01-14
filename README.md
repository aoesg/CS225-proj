# CS225-proj

## 环境依赖

### Golang
```
wget https://dl.google.com/go/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -zxvf  go1.21.5.linux-amd64.tar.gz
```

添加环境路径，在.bashrc中添加
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go
export PATH=$PATH:/sbin
```

如果安装成功
```
go version
// 显示go版本
```

针对国内网络进行优化
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

### Redis

```
wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
sudo apt install gcc
sudo apt install make
make
sudo make install
```

## 启动log server

log server的main包在`wzConfig_main`文件夹内
log server需要读取配置文件`wzConfig_main/config.json`

Example
```
{
    "httpIP": "0.0.0.0", // log-server 监听的ip
    "httpPort": "8080", // log-server 监听的port

    "local_redis_address": "localhost:50000", // 本地redis数据库的ip:port

    "local_log_id": 0, // 本 log-server 的 log_node_id
    "logNodes_address": ["node0:8080", "node1:8080"], //LogLayer中所有log node的地址
    "replica_ids": [1] // 拥有本 log-server 的备份的 log_node_id
    
}
```

启动log server
```
cd wzConfig_main
go run .
```

