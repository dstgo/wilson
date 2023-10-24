# Wilson

<div align="center">
<img src="https://public-1308755698.cos.ap-chongqing.myqcloud.com//img/202310122023621.png" alt="wilson" width="300px" height="300px" />

![Static Badge](https://img.shields.io/badge/go-%3E%3D1.21-blue)
![GitHub](https://img.shields.io/github/license/dstgo/wilson?color=red)

</div>

Englist|[简体中文](README.zh.md)

wilson is an open source, distributed, front-end and back-end separation, modern Don't Starve game server web management panel server, implemented based on Go language.

## Feature

- Fully visual control panel
- Multilingual internationalization support
- ELK log analysis
- Prometheus system monitoring
- Elastic control of instance resources
- Role API permission management
- Multi-user, sub-user management
- Open API support
- Distributed server management

## Run
Generate configuration template
```sh
$ wilson gen --d /etc/wilson
```
Run web server
```sh
$ wilson server --f /etc/wilson/config.yaml
```


## Construct

Enter the source code root directory and use make to build

```sh
make install
```

## How to contribute

1. Fork warehouse
2. Create your own branch
3. Submit changes
4. Push changes
5. Submit a Pull Request



