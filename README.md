# 基于go-zero 搭建的微服务活动架构

## 如何运行

#### 启动nginx mysql redis etcd jaeger prometheus grafana服务
```shell
sh server.sh
```

#### 业务服务启动
```
docker-compose up -d
```    

[附带 wikifx-trade-contest目录结构-基于go-zero略有修改]

```
pkg                     
   ├── api              业务接口逻辑层 所有的业务逻辑存放目录。
   │   ├── model        数据模型 数据管理层，仅用于操作管理数据，如数据库操作     
   ├── assets           资源
   ├── common           公共服务 
   ├── rpc              业务RPC逻辑层
   │   ├── model        数据模型 数据管理层，仅用于操作管理数据，如数据库操作
       ├── cronjob      定时任务，定时数据更新业务
       ├── model        数据模型 数据管理层，仅用于操作管理数据，如数据库操作
       ├── rmq          消息处理系统：mq和dq，处理一些高并发和延时消息业务
       ├── script       脚本，处理一些临时运营需求，临时数据修复
   ├── boot             初始化包 用于项目初始化参数设置，往往作为main.go中第一个被import的包
   ├── config           配置管理 所有的配置文件存放目录
   ├── library          公共库包 公共的功能封装包，往往不包含业务需求实现
   ├── packed           打包目录 将资源文件打包的Go文件存放在这里，boot包初始化时会自动调用
   ├── public           静态目录 仅有该目录下的文件才能对外提供静态服务访问
   ├── router           路由注册 用于路由统一的注册管理
   ├── go.mod           依赖管理 使用Go Module包管理的依赖描述文件
   ├── services        第三方服务管理
   ├── middleware       中间件 
   └── main.go          入口文件
```