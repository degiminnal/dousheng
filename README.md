# dousheng
### 技术选型与相关开发文档

> 项目采用三层架构设计，分别为entity，service，controller，结构清晰，代码耦合度低。

### 架构设计

> entity包包括了实体类的定义，对数据库的操作等功能。
>
> service包中实现对entity的调用，完善业务的逻辑，调用对数据库的操作来修改数据库。
>
> controller为路由到具体功能实现，在controller中调用service服务实现业务场景。
>
> 中间件使用JWT鉴权。
>

### 项目代码介绍

>  Configure
>
> + init.go                         读取configure内的项目相关配置
>
> + configure.yml            项目有关配置均在此文件内，运行前请先修改其中配置
>
>  Controller
>
> + common.go               设置包内公共函数、公共变量和公共类型定义
> + comment.go              实现评论操作接口
>
> + favorite.go                 实现喜欢操作接口
>
> + message.go               实现消息操作接口
>
> + publish.go                  实现发布操作接口
>
> + relation.go                 实现关系操作接口
>
> + user.go                        实现用户接口
> + video.go                       实现视频流接口
>
>  Entity
>
> + comment.go              评论实体类与Dao实现，包含若干数据库操作
>
> + favourite.go               喜欢实体类与Dao实现，包含若干数据库操作
>
> + init.go                         对mySql，Redis进行初始化
> + message.go              消息实体类与Dao实现，包含若干数据库操作
>
> + relation.go                关注实体类与Dao实现，包含若干数据库操作
> + user.go                      用户实体类与Dao实现，包含若干数据库操作
>
> + video.go                    评论实体类与Dao实现，包含若干数据库操作
>
> Middleware
>
> + Token.go                  实现 JWT 生成 token 和验证 token
>
>  Public
>
>  Video                         存放视频
>
>  Cover                         存放封面
>
>  Service
>
> + comment.go   评论业务逻辑实现
> + common.go    设置包内公共函数、公共变量和公共类型定义
>
> + favourite.go    喜欢业务逻辑实现
>
> + message.go    消息业务逻辑实现
> + redis.go           数据缓存业务逻辑实现
> + relation.go     关注业务逻辑实现
>
> + user.go           用户业务逻辑实现
>
> + video.go         视频业务逻辑实现
>
>  main.go                主函数
>

### 效果展示

[dousheng](http://www.degim.top/videos/dousheng.mp4)
