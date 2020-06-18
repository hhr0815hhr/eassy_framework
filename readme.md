## eassy_framework

#### 一个游戏后端分布式架构
<p>由于使用其他框架代码，很多地方都不是很理想，不是自己想要的，所以从零开始自己写一个后端框架架构
一步步完善游戏后端分布式架构</p>

`Etcd` `Grpc` `Protobuf`

#### Get Started
* 1.编译
* 2.确认etcd、mysql服务已启动并配置好
* 3.运行各节点
    - ./eassy 5020 login
    - ./eassy 5050 gate
    - ./eassy 5060 game
    - ./eassy 5070 center
  
#### 各节点介绍
* Login  
采用http服务，登录验证采用JWT(json web token),Restful Api
* Gate  
    - 与客户端通过websocket连接，也可使用socket连接
    - Gate与其他节点通过配置grpc服务进行消息的转发
* Game  
游戏服，目前大厅和游戏房间均在此节点，后续会进一步拆分
* Center  
中心服，暂未实现

***
#### todo  
- 考虑到数据库集群和主从搭建，为了架构进一步解耦，新增一个db服务节点
- Etcd多节点负载均衡策略

