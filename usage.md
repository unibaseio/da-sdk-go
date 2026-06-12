# Unibase DA 节点运维

> 节点部署的权威文档在 **da-go 仓库**（README + install.md + contract/test-v2/README），
> 本文为运维命令速查。启动前先设 `CHAIN_TYPE`（默认 base-sepolia）并完成充值
> （UB 质押 ≥ 3×罚金；原生 gas ≥ 0.1，1559 交易按 GasLimit×feeCap 预扣）。

## stream

流节点负责处理和分发数据，需要外网ip

+ standard

CPU: 16 core
Memory: 32GB
Storage: 4TB HDD
Network: public IP, 100mbps


### 运行

```shell
// 初始化，默认使用~/.stream目录, 创建account，密码默认aidemo123
> ./stream init
// 启动，默认使用8083端口，8083端口需要可以外网访问
> ./stream daemon run -b <BIND_PORT> -e <EXPOSE_URL>
// 例如外网ip为1.2.3.4
> ./stream daemon run -b 0.0.0.0:8083 -e http://1.2.3.4:8083
// 查看余额
> ./stream chain balance
// 查看收益
> ./stream chain stream revenue
// 取回收益
> ./stream chain stream withdraw
```

## store

存储节点，负责存放数据内容

+ standard

CPU: 16 core
Memory: 32GB
Storage: 4TB HDD
Network: 100mbps


### 运行

```shell
// 初始化，默认使用~/.store目录,创建account，密码默认aidemo123
> ./store-edge init
// 启动，默认使用8082端口
> ./store-edge daemon run -b <BIND_PORT> -e <EXPOSE_URL>
// 例如外网ip为1.2.3.4，8082端口可以外网访问
> ./store-edge daemon run -b 0.0.0.0:8082 -e http://1.2.3.4:8082
// 例如: 无外网ip，8082端口不可以外网访问
> ./store-edge daemon run -b 127.0.0.1:8082
// 查看余额
> ./stream chain balance
// 查看收益信息
> ./store-edge chain store revenue
// 取回存储收益
> ./store-edge chain store withdraw
// 更新存储奖励
> ./store-edge chain store updateReward
// 取回存储奖励
> ./store-edge chain store withdrawReward
```

## validator

验证store提交证据的正确性，在发现错误的时候提起挑战，获取收益

+ standard

CPU: 16 core
Memory: 32GB
Storage: 1TB SSD
Network: 10mbps

```shell
// 初始化，默认使用~/.validator,创建account，密码默认aidemo123
> ./validator init
// 默认使用8085端口
> ./validator daemon run -b <BIND_PORT>
// 例如
> ./validator daemon run -b 127.0.0.1:8085
// 查看余额
> ./validator chain balance
```


## client(todo)

### space(todo)


## compute(todo)

分为compute-controller和computer-worker，目前这两个是放在一起

计算节点，对gpu要求高，需要有外网IP

### 运行

```shell
// 初始化，默认使用~/.compute目录
> ./compute-edge init
// 默认使用8081端口，8081端口需要可以外网访问
> ./compute-edge daemon run -e EXPOSE_URL
// 例如外网ip为1.2.3.4
> ./compute-edge daemon run -b 0.0.0.0:8081 -e http://1.2.3.4:8081
```

### requirement

+ standard

CPU: 16 core
Memory: 32GB
GPU: NVIDIA GPU with 12GB memory
Storage: 1TB nvme
Network: public IP, 10mbps


### prepare

#### GPU驱动

1. 安装GPU驱动

https://ubuntu.com/server/docs/nvidia-drivers-installation

```shell
// 安装后，查看显卡信息
> nvidia-smi 
```

2. 安装cuda require: cuda11.7.1

https://developer.nvidia.com/cuda-downloads


#### 安装docker

https://yeasy.gitbook.io/docker_practice/install

+ 添加用户

```shell
// 创建docker组，将用户加入到此组
> sudo groupadd docker
> sudo usermod -aG docker $USER
> sudo systemctl restart docker
> docker info
```

+ 安装docker nvidia toolkit

https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html

+ pull image 

```shell
> docker pull nvidia/cuda:11.7.1-cudnn8-runtime-ubuntu20.04
```






