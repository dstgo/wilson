# Wilson
<div align="center">
<img src="https://public-1308755698.cos.ap-chongqing.myqcloud.com//img/202310122023621.png" alt="wilson" width="300px" height="300px" />


![Static Badge](https://img.shields.io/badge/go-%3E%3D1.21-blue)
![GitHub](https://img.shields.io/github/license/dstgo/wilson?color=red)

</div>
wilson 是一款开源的，分布式，前后端分离，现代化的饥荒游戏服务器网页管理面板的服务端，基于Go语言实现。

## 特色

- 完全可视化的控制面板
- 多语言国际化支持
- ELK日志分析
- Prometheus系统监控
- 实例资源弹性控制
- 角色API权限管理
- 多用户，子用户管理
- 开放API支持
- 分布式服务器管理


## 运行
生成配置模板
```sh
$ wilson gen --d /etc/wilson
```
运行web服务器
```sh
$ wilson server --f /etc/wilson/config.yaml 
```


## 构建

进入源代码根目录，使用make来进行构建

```sh
make install
```



## 如何贡献

1. Fork仓库
2. 创建自己的分支
3. 提交修改
4. 推送修改
5. 提交Pull Request