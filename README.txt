由于使用其他框架代码，很多地方都不是很理想，不是自己想要的，所以从零开始自己写一个后端框架架构
一步步完善游戏后端分布式架构
持续更新

Get Started >>
    1.编译
    2.确认etcd、mysql服务已启动并配置好
    3.运行各节点
        ./eassy 5020 login
        ./eassy 5050 gate
        ./eassy 5060 game
        ./eassy 5070 center


login>>
    采用http服务，登录验证采用JWT(json web token),Restful Api

gate和其他节点>>
    gate采用websocket与客户端进行连接
    gate与其他节点通过配置grpc服务进行消息的转发

etcd  done
grpc  done
login done
game  done

todo db单独作为一个服务节点

>> To be continued !