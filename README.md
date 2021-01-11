# 商场项目

一个前后端分离的商场项目，包含完整购物，下单流程

[演示地址](https://mall.longji.online)

```
账号： test
密码： test
```

![商品首页](https://github.com/wljgithub/mall-project/blob/develop/webapp/static-files/%E9%A6%96%E9%A1%B5.png)
![购物车](https://github.com/wljgithub/mall-project/blob/develop/webapp/static-files/%E8%B4%AD%E7%89%A9%E8%BD%A6.png)
![商品搜索](https://github.com/wljgithub/mall-project/blob/develop/webapp/static-files/%E5%95%86%E5%93%81%E6%90%9C%E7%B4%A2.png)

## 特性

- 使用Makefile管理编译流程
- 支持HTTPS（Let's Encrypt的证书,需要域名）
- Nginx 反向代理后端应用，支持HTTP2
- 前后端容器化，使用docker-compose将Mysql，Redis，前后端应用整合起来，支持一键启动

## 项目介绍

采用前后端分离的架构，前端用Vue实现，后端用[Gin](https://github.com/gin-gonic/gin) (Go 的一个Web框架)提供RestfulAPI

### 项目结构
```
.
├── CHANGELOG.md
├── LICENSE
├── Makefile                                  # Make 配置
├── README.md
├── docker-compose.yml                        # docker-compose配置文件
├── docs                                      # 项目文档
├── nginx                                     # nginx 配置
├── scripts                                   # letsencrypt 初始化脚本
├── server                                    # 后端代码
└── webapp                                    # 前端代码

```

### 后端介绍

后端项目结构：
```
.
├── Dockerfile                          # Docker文件
├── README.md
├── cmd
│   └── main.go                         # 项目入口
├── conf                                # 配置文件
│   ├── config.deploy.yml               
│   └── config.local.yml                
├── go.mod
├── go.sum
├── internal
│   ├── api                             # api层，路由注册和路由函数实现
│   ├── dto                             # dto层，request和response的数据对象
│   ├── mapper                          # mapper层，model层到dto层的数据转换
│   ├── model                           # model层，对应mysql表结构
│   ├── repository                      # repository层，表的增删查改
│   └── service                         # 业务逻辑层
├── pkg
│   ├── app                             # 负责工程初始化，初始化repository层，service层，API层
│   ├── conf                            # 配置模块
│   ├── database                        
│   ├── errno                           # 业务逻辑错误码
│   ├── handler                         # 接口响应的封装
│   ├── log                             # 日志模块，支持
│   ├── token                           # jwt token
│   └── util                            # 工具函数
├── scripts
│   └── wait_for.sh                     # 一个shell脚本，docker-compose时等待mysql初始化
└── test                                
    ├── mall_data.sql                   
    └── mall_test.sql                   # 数据表结构


```
后端采用三层架构模型（Repository，Service，API），通过依赖注入的方式实现控制反转，使用[wire](https://github.com/google/wire) 管理对象的初始化

实现了

- Graceful Shutdown
- Jwt登录认证， 基于 [jwt-go](https://github.com/dgrijalva/jwt-go)
- 配置热加载，基于 [viper](https://github.com/spf13/viper)
- 日志结构化，日志切分，基于[zap](https://github.com/uber-go/zap)

## 怎么用？

环境依赖：

 - make
 - docker
 - docker-compose
 
安装完依赖，把项目拉下来
 
```
cd mall-project && make serve
```
已经编写好了Makefile，`make serve` 即可启动项目，不出意外的话访问 http://localhost 就是商场首页了

有问题可以 `make log`可查看日志

 



## 感谢

前端是基于 [newbee-mall-vue3-app](https://github.com/newbee-ltd/newbee-mall-vue3-app) 修改

后端日志模块，配置模块参考了 [snake](https://github.com/1024casts/snake)

## License

MIT