# User Management
## 目录
  - [功能](#功能)
  - [命令](#命令)
  - [步骤](#步骤)
  - [最佳实践](#最佳实践)
  - [关于mysql](#关于mysql)
  - [GRPC](#grpc)

-----------
### 功能
- [] 注册
- [] 设置密码
- [] 获取用户信息
- [] 更新用户信息
- [] 绑定用户手机/邮箱
- [] 解除绑定用户手机/邮箱
- [] 绑定第三方账户
- [] 解除绑定第三方账户
- [] 更改用户密码
- [] 忘记密码
- [] 添加用户
- [] 删除用户
- [] 用户注销


### 命令
* make init ```初始化仓库，创建go.mod```
* make verify ```验证开发环境与构建环境，检查code conduct```
* make verify-build ```编译目标```
* make test ```单元测试```
* make generate-docker-images ```生成docker镜像```
* make service-sample ```单独编译服务```
* make service-sample-image ```单独生成服务镜像```
* make deploy-to-k8s-cluster ```部署到k8s集群```

### 最佳实践
* 每个服务只提供单一可执行文件，有利于docker镜像打包与k8s部署管理
* 每个服务提供http调试接口，通过curl获取调试信息
* 集群内服务间direct call调用通过服务发现获取目标地址进行调用
* 集群内服务间event call调用通过rabbitmq解耦

### 关于mysql
* 创建app后，从app.Mysql()获取本地mysql client
* [文档参考](https://entgo.io/docs/sql-integration)
