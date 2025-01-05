-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2024-08-22 15:00:05
-- 服务器版本： 5.7.40-log
-- PHP 版本： 7.4.33

SET
SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET
time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `configure`
--

-- --------------------------------------------------------

--
-- 表的结构 `business`
--

CREATE TABLE IF NOT EXISTS `business`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `server_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '服务id',
    `keyword` char
(
    32
) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '变量标识',
    `type` varchar
(
    32
) NOT NULL COMMENT '变量类型',
    `description` varchar
(
    128
) NOT NULL COMMENT '变量描述',
    PRIMARY KEY
(
    `id`
),
    KEY `idx_business_created_at`
(
    `created_at`
),
    KEY `idx_business_updated_at`
(
    `updated_at`
),
    KEY `fk_business_server`
(
    `server_id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='业务变量';

--
-- 转存表中的数据 `business`
--

INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`)
VALUES (1, 1712995716, 1712995716, 2, 'AuthSkipRoles', 'object', '跳过权限检测的角色列表'),
       (2, 1712995716, 1712995716, 2, 'DefaultUserPassword', 'string', '默认账号密码'),
       (4, 1712995716, 1712995716, 2, 'LoginPrivatePath', 'string', 'rsa解密私钥路径'),
       (5, 1712995716, 1712995716, 2, 'Setting', 'object', '系统设置'),
       (6, 1712995716, 1719465188, 3, 'DefaultAcceptTypes', 'object', '允许上传的文件后缀'),
       (8, 1712995716, 1719465209, 3, 'ChunkSize', 'int', '单个切片最大大小（M）'),
       (9, 1712995716, 1719465230, 3, 'DefaultMaxSize', 'int', '文件最大大小（M）'),
       (15, 1719462652, 1719462652, 2, 'ChangePasswordType', 'string', '修改密码方式');

-- --------------------------------------------------------

--
-- 表的结构 `business_value`
--

CREATE TABLE IF NOT EXISTS `business_value`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `env_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '环境id',
    `business_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '业务变量id',
    `value` text NOT NULL COMMENT '业务变量id',
    PRIMARY KEY
(
    `id`
),
    UNIQUE KEY `env_id`
(
    `env_id`,
    `business_id`
),
    KEY `idx_business_value_created_at`
(
    `created_at`
),
    KEY `idx_business_value_updated_at`
(
    `updated_at`
),
    KEY `fk_business_value_env`
(
    `env_id`
),
    KEY `fk_business_values`
(
    `business_id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COMMENT='业务变量值';

--
-- 转存表中的数据 `business_value`
--

INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`)
VALUES (52, 1719371012, 1719465259, 1, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\",\"pdf\"]'),
       (53, 1719371012, 1719465259, 2, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]'),
       (54, 1719371012, 1719465259, 3, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]'),
       (61, 1719382347, 1719465239, 1, 9, '10'),
       (62, 1719382347, 1719465239, 2, 9, '10'),
       (63, 1719382347, 1719465239, 3, 9, '10'),
       (64, 1719382352, 1719382352, 1, 8, '1'),
       (65, 1719382352, 1719382352, 2, 8, '1'),
       (66, 1719382352, 1719382352, 3, 8, '1'),
       (73, 1719462662, 1719462662, 1, 15, 'password'),
       (74, 1719462662, 1719462662, 2, 15, 'password'),
       (75, 1719462662, 1719462662, 3, 15, 'email'),
       (76, 1719463071, 1719463071, 1, 4, 'static/cert/login.pem'),
       (77, 1719463071, 1719463071, 2, 4, 'static/cert/login.pem'),
       (78, 1719463071, 1719463071, 3, 4, 'static/cert/login.pem'),
       (79, 1719463092, 1719463092, 1, 1, '[\"superAdmin\"]'),
       (80, 1719463092, 1719463092, 2, 1, '[\"superAdmin\"]'),
       (81, 1719463092, 1719463092, 3, 1, '[\"superAdmin\"]'),
       (82, 1719463137, 1719463137, 1, 2, '12345678'),
       (83, 1719463137, 1719463137, 2, 2, '12345678'),
       (84, 1719463137, 1719463137, 3, 2, '12345678'),
       (85, 1719463217, 1719463217, 1, 5,
        '{\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\",\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\"}'),
       (86, 1719463217, 1719463217, 2, 5,
        '{\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\"}'),
       (87, 1719463217, 1719463217, 3, 5,
        '{\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\"}');

-- --------------------------------------------------------

--
-- 表的结构 `configure`
--

CREATE TABLE IF NOT EXISTS `configure`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `server_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '服务id',
    `env_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '环境id',
    `content` text NOT NULL COMMENT '配置内容',
    `version` varchar
(
    32
) NOT NULL COMMENT '配置版本',
    `format` varchar
(
    32
) NOT NULL COMMENT '配置格式',
    `description` varchar
(
    128
) DEFAULT NULL COMMENT '配置描述',
    PRIMARY KEY
(
    `id`
),
    KEY `idx_configure_updated_at`
(
    `updated_at`
),
    KEY `idx_configure_created_at`
(
    `created_at`
),
    KEY `fk_configure_server`
(
    `server_id`
),
    KEY `fk_configure_env`
(
    `env_id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='配置内容';

--
-- 转存表中的数据 `configure`
--

INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`,
                         `description`)
VALUES (1, 1712995716, 1719560095, 1, 1,
        'addr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /usercenter/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n  - path: /cron/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7040\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/api/v1/auth\n          method: POST',
        'a5572902d3bc929adc0d95ba1a97e0ac', 'yaml', '初始化模板'),
       (2, 1712995716, 1719464071, 2, 1,
        'test: 11\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7010\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8010\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: fyudafdzqmhwbfbd\njwt:\n  redis: cache\n  secret: limes-cloud-test\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true,\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true,\"POST:/manager/api/v1/user/token/refresh\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: [\"superAdmin\"]\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8010\nbusiness:\n  changePasswordType: password\n  defaultUserPassword: 12345678\n  setting: {\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\",\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\"}\n',
        '5bbbfed4b134b3d73127cacf6b0b17d0', 'yaml', '初始化模板'),
       (3, 1712995716, 1719466309, 3, 1,
        '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7020\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8020\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \nbusiness:\n  chunkSize: 1\n  defaultMaxSize: 10\n  defaultAcceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\",\"pdf\"]\n  storage:\n    type: local\n    serverUrl: http://127.0.0.1:7080/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: http://127.0.0.1:7080/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        '0266de2d040fa6a857c691196eee79dc', 'yaml', '初始化模板'),
       (4, 1712995716, 1719560133, 5, 1,
        '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7030\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8030\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8020\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: usercenter\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: fyudafdzqmhwbfbd\njwt:\n  redis: cache\n  secret: limes-cloud-client-test\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"*:/usercenter/api/*\":true,\"*:/usercenter/client/*\":true}\nbusiness:\n  defaultPasswordExpire: 10s\n  defaultNickName: 用户-\n',
        'd44acfbc819d15a9d4cb54fe85aa465c', 'yaml', '初始化模板'),
       (5, 1712995716, 1712995716, 6, 1,
        '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n',
        'dfff643162f3c96185e12ca22bf2033b', 'yaml', '自动初始化'),
       (6, 1712995716, 1719560100, 1, 2,
        'addr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /usercenter/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n  - path: /cron/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7040\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/api/v1/auth\n          method: POST',
        'a5572902d3bc929adc0d95ba1a97e0ac', 'yaml', '初始化模板'),
       (7, 1712995716, 1719464076, 2, 2,
        'test: 11\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7010\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8010\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-pre\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true,\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true,\"POST:/manager/api/v1/user/token/refresh\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: [\"superAdmin\"]\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8010\nbusiness:\n  changePasswordType: password\n  defaultUserPassword: 12345678\n  setting: {\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\"}\n',
        '1d2555ef65b788c69f3254a92719cc0b', 'yaml', '初始化模板'),
       (8, 1712995716, 1719467313, 3, 2,
        '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7020\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8020\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \nbusiness:\n  chunkSize: 1\n  defaultMaxSize: 10\n  defaultAcceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]\n  storage:\n    type: local\n    serverUrl: http://127.0.0.1:7080/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: http://127.0.0.1:7080/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        '84cdc6dbc3e7e5d04af6f5f1da605013', 'yaml', '初始化模板'),
       (9, 1712995716, 1719560149, 5, 2,
        '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7030\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8030\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8020\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: usercenter\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-client-pre\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"*:/usercenter/api/*\":true,\"*:/usercenter/client/*\":true}\nbusiness:\n  defaultPasswordExpire: 10s\n  defaultNickName: 用户-\n',
        '8fdb079a8c3c78ec764722f69c2e6258', 'yaml', '初始化模板'),
       (10, 1712995716, 1712995716, 6, 2,
        '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n',
        '6311dada0a48b1d307068a92da31a3f1', 'yaml', '自动初始化'),
       (11, 1712995716, 1719560105, 1, 3,
        'addr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7010\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7020\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/client/v1/auth\n          method: POST\n  - path: /usercenter/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7030\n  - path: /cron/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7040\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/usercenter/api/v1/auth\n          method: POST',
        'a5572902d3bc929adc0d95ba1a97e0ac', 'yaml', '初始化模板'),
       (12, 1712995716, 1719464083, 2, 3,
        'test: 11\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7010\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8010\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: manager\n      password: maGh8TzkjfyLbkYA\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: fyudafdzqmhwbfbd\njwt:\n  redis: cache\n  secret: limes-cloud-prod\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true,\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true,\"POST:/manager/api/v1/user/token/refresh\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: [\"superAdmin\"]\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8010\nbusiness:\n  changePasswordType: email\n  defaultUserPassword: 12345678\n  setting: {\"title\":\"统一应用管理平台\",\"desc\":\"开放协作，拥抱未来，统一应用管理平台\",\"copyright\":\"Copyright © 2024 lime.qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\",\"watermark\":\"go-platform\"}\n',
        '7065bd9d55a4a538af44591c9b152b9a', 'yaml', '初始化模板'),
       (13, 1712995716, 1719495066, 3, 3,
        '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7020\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8020\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: resource\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \nbusiness:\n  chunkSize: 1\n  defaultMaxSize: 10\n  defaultAcceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]\n  storage:\n    type: local\n    serverUrl: http://gw.qlime.cn/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: http://gw.qlime.cn/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        '7249a2261b8dfbe9ea2d9be1b2d5da64', 'yaml', '1'),
       (14, 1712995716, 1719560139, 5, 3,
        '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7030\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8030\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8020\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: usercenter\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: usercenter\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: fyudafdzqmhwbfbd\njwt:\n  redis: cache\n  secret: limes-cloud-client-prod\n  expire: 2h\n  renewal: 600s\n  whitelist: {\"*:/usercenter/api/*\":true,\"*:/usercenter/client/*\":true}\nbusiness:\n  defaultPasswordExpire: 10s\n  defaultNickName: 用户-\n',
        'f7ce62b3fdd5242871efb02b413062af', 'yaml', '初始化模板'),
       (15, 1712995716, 1712995716, 6, 3,
        '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: party_affairs\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n',
        '8cd386b9b9f1f4be35be3695edf65380', 'yaml', '自动初始化'),
       (19, 1719737256, 1719737256, 4, 1,
        'env: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7040\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8040\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: true\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: cron\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\npool: #并非池配置\n  size: 10000 #最大协程数量\n  expiryDuration: 30s #过期时间\n  preAlloc: true #是否预分配\n  maxBlockingTasks: 1000 #最大的并发任务\n  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务',
        'ab5437417c2ab39a89ff2169f2a1d60c', 'yaml', '初始化模板'),
       (20, 1719737263, 1719739087, 4, 3,
        'env: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7040\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8040\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: true\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: cron\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: cron\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\npool: #并非池配置\n  size: 10000 #最大协程数量\n  expiryDuration: 30s #过期时间\n  preAlloc: true #是否预分配\n  maxBlockingTasks: 1000 #最大的并发任务\n  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务',
        '3720664a53b786217fdd36978ea8f9cb', 'yaml', '初始化模板'),
       (21, 1719737269, 1719737269, 4, 2,
        'env: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7040\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8040\n    timeout: 10s\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: true\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: cron\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\npool: #并非池配置\n  size: 10000 #最大协程数量\n  expiryDuration: 30s #过期时间\n  preAlloc: true #是否预分配\n  maxBlockingTasks: 1000 #最大的并发任务\n  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务',
        'fe37cdd562db038d0805d44826cf571e', 'yaml', '初始化模板');

-- --------------------------------------------------------

--
-- 表的结构 `env`
--

CREATE TABLE IF NOT EXISTS `env`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `keyword` char
(
    32
) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '环境标识',
    `name` varchar
(
    64
) NOT NULL COMMENT '环境名称',
    `description` varchar
(
    128
) NOT NULL COMMENT '环境描述',
    `token` varchar
(
    128
) NOT NULL COMMENT '环境token',
    `status` tinyint
(
    1
) DEFAULT '0' COMMENT '环境状态',
    PRIMARY KEY
(
    `id`
),
    KEY `idx_env_created_at`
(
    `created_at`
),
    KEY `idx_env_updated_at`
(
    `updated_at`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='环境信息';

--
-- 转存表中的数据 `env`
--

INSERT INTO `env` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `token`, `status`)
VALUES (1, 1712995716, 1722220395, 'TEST', '测试环境', '用于本地测试', '1DDA3033E492425116837569E1917CA1', 1),
       (2, 1712995716, 1719245213, 'PRE', '预发布环境', '用于上线前测试', '862BBE3D5BE34A780305DA84A8DD5147', 1),
       (3, 1712995716, 1712995716, 'PROD', '生产环境', '用于线上真实环境', '5B655B7D4A51BF79C974C9F27C27D992', 1);

-- --------------------------------------------------------

--
-- 表的结构 `gorm_init`
--

CREATE TABLE IF NOT EXISTS `gorm_init`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT,
    `init` tinyint
(
    1
) DEFAULT NULL,
    PRIMARY KEY
(
    `id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `gorm_init`
--

INSERT INTO `gorm_init` (`id`, `init`)
VALUES (1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `resource`
--

CREATE TABLE IF NOT EXISTS `resource`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `keyword` char
(
    32
) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '变量标识',
    `description` varchar
(
    128
) NOT NULL COMMENT '变量描述',
    `fields` varchar
(
    256
) NOT NULL COMMENT '变量字段集合',
    `tag` varchar
(
    32
) NOT NULL COMMENT '变量标签',
    `private` tinyint
(
    1
) DEFAULT '0' COMMENT '是否私有',
    PRIMARY KEY
(
    `id`
),
    KEY `idx_resource_created_at`
(
    `created_at`
),
    KEY `idx_resource_updated_at`
(
    `updated_at`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COMMENT='资源变量';

--
-- 转存表中的数据 `resource`
--

INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`)
VALUES (1, 1712995716, 1712995716, 'Env', '环境标识信息', 'Keyword', 'env', 0),
       (2, 1712995716, 1712995716, 'AdminJwt', '后台管理服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', 0),
       (3, 1712995716, 1712995716, 'ClientJwt', '客户端服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', 0),
       (4, 1712995716, 1712995716, 'Redis', 'redis中间件配置信息', 'Host,Username,Password,Port', 'redis', 0),
       (5, 1712995716, 1712995716, 'Email', '邮箱服务配置信息', 'Username,Company,Password,Host,Port', 'email', 0),
       (6, 1712995716, 1712995716, 'GatewayServer', '通用网关服务配置信息', 'HttpPort,Host,Timeout', 'server', 0),
       (7, 1712995716, 1712995716, 'ManagerServer', '管理中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server',
        0),
       (8, 1712995716, 1712995716, 'ManagerDatabase', '管理中心数据库配置',
        'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1),
       (9, 1712995716, 1712995716, 'ResourceServer', '资源中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server',
        0),
       (10, 1712995716, 1712995716, 'ResourceDatabase', '资源中心数据库配置信息',
        'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1),
       (11, 1712995716, 1712995716, 'CronServer', '定时任务服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server',
        0),
       (12, 1712995716, 1712995716, 'CronDatabase', '定时任务数据库配置信息',
        'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1),
       (13, 1712995716, 1712995716, 'ConfigureServer', '配置中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout',
        'server', 0),
       (14, 1712995716, 1712995716, 'UserCenterServer', '用户中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout',
        'server', 0),
       (15, 1712995716, 1712995716, 'UserCenterDatabase', '用户中心数据库配置信息',
        'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1),
       (16, 1712995716, 1712995716, 'PartyAffairsServer', '信号灯服务配置信息', 'Host,HttpPort,GrpcPort,Timeout',
        'server', 0),
       (17, 1712995716, 1719371833, 'PartyAffairsDatabase', '信号灯数据库配置信息',
        'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);

-- --------------------------------------------------------

--
-- 表的结构 `resource_server`
--

CREATE TABLE IF NOT EXISTS `resource_server`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `server_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '服务id',
    `resource_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '资源id',
    PRIMARY KEY
(
    `id`
),
    UNIQUE KEY `sr`
(
    `server_id`,
    `resource_id`
),
    KEY `idx_resource_server_created_at`
(
    `created_at`
),
    KEY `fk_resource_resource_server`
(
    `resource_id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='资源服务信息';

--
-- 转存表中的数据 `resource_server`
--

INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`)
VALUES (1, 1712995716, 2, 8),
       (2, 1712995716, 3, 10),
       (3, 1712995716, 4, 12),
       (4, 1712995716, 5, 15),
       (6, NULL, 6, 17);

-- --------------------------------------------------------

--
-- 表的结构 `resource_value`
--

CREATE TABLE IF NOT EXISTS `resource_value`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `env_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '环境id',
    `resource_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '资源id',
    `value` text NOT NULL COMMENT '资源id',
    PRIMARY KEY
(
    `id`
),
    UNIQUE KEY `er`
(
    `env_id`,
    `resource_id`
),
    KEY `idx_resource_value_updated_at`
(
    `updated_at`
),
    KEY `idx_resource_value_created_at`
(
    `created_at`
),
    KEY `fk_resource_resource_value`
(
    `resource_id`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COMMENT='资源变量值';

--
-- 转存表中的数据 `resource_value`
--

INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`)
VALUES (1, 1712995716, 1719371433, 1, 1, '{\"Keyword\":\"TEST\"}'),
       (2, 1712995716, 1719371433, 2, 1, '{\"Keyword\":\"PRE\"}'),
       (3, 1712995716, 1719371433, 3, 1, '{\"Keyword\":\"PROD\"}'),
       (4, 1712995716, 1719457890, 1, 2,
        '{\"Secret\":\"limes-cloud-test\",\"Expire\":\"2h\",\"Renewal\":\"600s\",\"Whitelist\":{\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true,\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true,\"POST:/manager/api/v1/user/token/refresh\":true}}'),
       (5, 1712995716, 1719457890, 2, 2,
        '{\"Renewal\":\"600s\",\"Whitelist\":{\"POST:/manager/api/v1/user/token/refresh\":true,\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true,\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true},\"Secret\":\"limes-cloud-pre\",\"Expire\":\"2h\"}'),
       (6, 1712995716, 1719457890, 3, 2,
        '{\"Secret\":\"limes-cloud-prod\",\"Expire\":\"2h\",\"Renewal\":\"600s\",\"Whitelist\":{\"POST:/manager/api/v1/user/login\":true,\"POST:/manager/api/v1/user/logout\":true,\"POST:/manager/api/v1/user/token/refresh\":true,\"GET:/manager/api/v1/system/setting\":true,\"GET:/manager/api/v1/user/login/captcha\":true}}'),
       (7, 1712995716, 1719467681, 1, 3,
        '{\"Secret\":\"limes-cloud-client-test\",\"Expire\":\"2h\",\"Renewal\":\"600s\",\"Whitelist\":{\"*:/usercenter/client/*\":true,\"*:/usercenter/api/*\":true}}'),
       (8, 1712995716, 1719467681, 2, 3,
        '{\"Secret\":\"limes-cloud-client-pre\",\"Expire\":\"2h\",\"Renewal\":\"600s\",\"Whitelist\":{\"*:/usercenter/api/*\":true,\"*:/usercenter/client/*\":true}}'),
       (9, 1712995716, 1719467681, 3, 3,
        '{\"Whitelist\":{\"*:/usercenter/api/*\":true,\"*:/usercenter/client/*\":true},\"Secret\":\"limes-cloud-client-prod\",\"Expire\":\"2h\",\"Renewal\":\"600s\"}'),
       (10, 1712995716, 1712995716, 1, 4, '{\"Host\":\"127.0.0.1\",\"Username\":\"\",\"Password\":\"\",\"Port\":6379}'),
       (11, 1712995716, 1712995716, 2, 4, '{\"Password\":\"\",\"Port\":6379,\"Host\":\"127.0.0.1\",\"Username\":\"\"}'),
       (12, 1712995716, 1712995716, 3, 4, '{\"Host\":\"127.0.0.1\",\"Username\":\"\",\"Password\":\"\",\"Port\":6379}'),
       (13, 1712995716, 1714884855, 1, 5,
        '{\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\",\"Password\":\"fyudafdzqmhwbfbd\",\"Host\":\"smtp.qq.com\",\"Port\":25}'),
       (14, 1712995716, 1714884855, 2, 5,
        '{\"Password\":\"xxx\",\"Host\":\"smtp.qq.com\",\"Port\":25,\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\"}'),
       (15, 1712995716, 1714884855, 3, 5,
        '{\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\",\"Password\":\"fyudafdzqmhwbfbd\",\"Host\":\"smtp.qq.com\",\"Port\":25}'),
       (16, 1712995716, 1712995716, 1, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}'),
       (17, 1712995716, 1712995716, 2, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}'),
       (18, 1712995716, 1712995716, 3, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}'),
       (19, 1712995716, 1719462382, 1, 7,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":7010,\"GrpcPort\":8010,\"Timeout\":\"10s\"}'),
       (20, 1712995716, 1719462382, 2, 7,
        '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7010,\"GrpcPort\":8010}'),
       (21, 1712995716, 1719462382, 3, 7,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":7010,\"GrpcPort\":8010,\"Timeout\":\"10s\"}'),
       (22, 1712995716, 1712995716, 1, 8,
        '{\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\"}'),
       (23, 1712995716, 1712995716, 2, 8,
        '{\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}'),
       (24, 1712995716, 1712995716, 3, 8,
        '{\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"manager\",\"Password\":\"maGh8TzkjfyLbkYA\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}'),
       (25, 1712995716, 1719462402, 1, 9,
        '{\"GrpcPort\":8020,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7020}'),
       (26, 1712995716, 1719462402, 2, 9,
        '{\"GrpcPort\":8020,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7020}'),
       (27, 1712995716, 1719462402, 3, 9,
        '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7020,\"GrpcPort\":8020}'),
       (28, 1712995716, 1712995716, 1, 10,
        '{\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\"}'),
       (29, 1712995716, 1712995716, 2, 10,
        '{\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\"}'),
       (30, 1712995716, 1712995716, 3, 10,
        '{\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"resource\"}'),
       (31, 1712995716, 1719462436, 1, 11,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":7040,\"GrpcPort\":8040,\"Timeout\":\"10s\"}'),
       (32, 1712995716, 1719462436, 2, 11,
        '{\"GrpcPort\":8040,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7040}'),
       (33, 1712995716, 1719462436, 3, 11,
        '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7040,\"GrpcPort\":8040}'),
       (34, 1712995716, 1719739039, 1, 12,
        '{\"Database\":\"cron\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}'),
       (35, 1712995716, 1719739039, 2, 12,
        '{\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"cron\"}'),
       (36, 1712995716, 1719739039, 3, 12,
        '{\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"cron\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"cron\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\"}'),
       (37, 1712995716, 1712995716, 1, 13,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":6081,\"GrpcPort\":6082,\"Timeout\":\"10s\"}'),
       (38, 1712995716, 1712995716, 2, 13,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":6081,\"GrpcPort\":6082,\"Timeout\":\"10s\"}'),
       (39, 1712995716, 1712995716, 3, 13,
        '{\"GrpcPort\":6082,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":6081}'),
       (40, 1712995716, 1719462419, 1, 14,
        '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7030,\"GrpcPort\":8030}'),
       (41, 1712995716, 1719462419, 2, 14,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":7030,\"GrpcPort\":8030,\"Timeout\":\"10s\"}'),
       (42, 1712995716, 1719462419, 3, 14,
        '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7030,\"GrpcPort\":8030}'),
       (43, 1712995716, 1719493403, 1, 15,
        '{\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"usercenter\"}'),
       (44, 1712995716, 1719493403, 2, 15,
        '{\"Database\":\"usercenter\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}'),
       (45, 1712995716, 1719493403, 3, 15,
        '{\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"usercenter\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"usercenter\"}'),
       (46, 1712995716, 1712995716, 1, 16,
        '{\"GrpcPort\":8100,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7100}'),
       (47, 1712995716, 1712995716, 2, 16,
        '{\"Host\":\"127.0.0.1\",\"HttpPort\":7100,\"GrpcPort\":8100,\"Timeout\":\"10s\"}'),
       (48, 1712995716, 1712995716, 3, 16,
        '{\"HttpPort\":7100,\"GrpcPort\":8100,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\"}'),
       (49, 1712995716, 1712995716, 1, 17,
        '{\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\"}'),
       (50, 1712995716, 1712995716, 2, 17,
        '{\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\"}'),
       (51, 1712995716, 1712995716, 3, 17,
        '{\"Type\":\"mysql\",\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"party_affairs\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\"}');

-- --------------------------------------------------------

--
-- 表的结构 `server`
--

CREATE TABLE IF NOT EXISTS `server`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `keyword` char
(
    32
) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务标识',
    `name` varchar
(
    64
) NOT NULL COMMENT '服务名称',
    `description` varchar
(
    128
) NOT NULL COMMENT '服务描述',
    `status` tinyint
(
    1
) DEFAULT '0',
    PRIMARY KEY
(
    `id`
),
    KEY `idx_server_created_at`
(
    `created_at`
),
    KEY `idx_server_updated_at`
(
    `updated_at`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='服务信息';

--
-- 转存表中的数据 `server`
--

INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`)
VALUES (1, 1712995716, 1720323279, 'Gateway', '通用网关', '主要负责前端到后端的转发', 1),
       (2, 1712995716, 1720753156, 'Manager', '管理中心', '主要负责系统的基础管理', 1),
       (3, 1712995716, 1720753158, 'Resource', '资源中心', '主要负责静态资源的管理', 1),
       (4, 1712995716, 1720637713, 'Cron', '定时任务', '主要负责定时任务的管理', 1),
       (5, 1712995716, 1720637716, 'UserCenter', '用户中心', '主要负责业务用户的管理', 1),
       (6, 1712995716, 1722239840, 'PartyAffairs', '信号灯系统', '指尖上的党务系统', 1);

-- --------------------------------------------------------

--
-- 表的结构 `template`
--

CREATE TABLE IF NOT EXISTS `template`
(
    `id`
    int
(
    10
) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` bigint
(
    20
) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint
(
    20
) DEFAULT NULL COMMENT '修改时间',
    `server_id` int
(
    10
) UNSIGNED NOT NULL COMMENT '模板id',
    `content` text NOT NULL COMMENT '模板内容',
    `version` varchar
(
    32
) NOT NULL COMMENT '模板版本',
    `is_use` tinyint
(
    1
) DEFAULT '0' COMMENT '是否使用',
    `format` varchar
(
    32
) NOT NULL COMMENT '模板格式',
    `description` varchar
(
    128
) NOT NULL COMMENT '模板描述',
    `compare` text NOT NULL COMMENT '变更详情',
    PRIMARY KEY
(
    `id`
),
    UNIQUE KEY `sv`
(
    `server_id`,
    `version`
),
    KEY `idx_template_created_at`
(
    `created_at`
),
    KEY `idx_template_updated_at`
(
    `updated_at`
)
    ) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COMMENT='配置模板';

--
-- 转存表中的数据 `template`
--

INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`,
                        `description`, `compare`)
VALUES (1, 1712995716, 1719560089, 1,
        '\ndebug: true\naddr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/v1/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /configure/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /resource/v1/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/v1/static/*\n              method: GET\n  - path: /cron/v1/*\n    timeout: ${CronServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${CronServer.Host}:${CronServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /user-center/v1/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /user-center/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /party-affairs/v1/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /party-affairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n',
        '54E50FB3522C', 0, 'yaml', '初始化模板', ''),
       (2, 1712995716, 1719473730, 2,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: ${AuthSkipRoles}\nbusiness:\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        '15768C4C6F57', 0, 'yaml', '初始化模板', ''),
       (3, 1712995716, 1719495060, 3,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nfile:\n  storage: local\n  serverPath: /resource/v1/static\n  localDir: static\n  maxSingularSize: ${MaxSingularSize}\n  maxChunkSize: ${MaxChunkSize}\n  maxChunkCount: ${MaxChunkCount}\n  acceptTypes: ${AcceptTypes}\n\n',
        '56945B81FA4D', 0, 'yaml', '初始化模板', ''),
       (4, 1712995716, 1719737246, 4,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${CronServer.Host}\n    port: ${CronServer.HttpPort}\n    timeout: ${CronServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${CronServer.Host}\n    port: ${CronServer.GrpcPort}\n    timeout: ${CronServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${CronDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${CronDatabase.Username}\n      password: ${CronDatabase.Password}\n      host: ${CronDatabase.Host}\n      port: ${CronDatabase.Port}\n      dbName: ${CronDatabase.Database}\n      option: ${CronDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\npool: #并非池配置\n  size: 10000 #最大协程数量\n  expiryDuration: 30s #过期时间\n  preAlloc: true #是否预分配\n  maxBlockingTasks: 1000 #最大的并发任务\n  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务\nbusiness:\n  key: value\n',
        'F2893009A192', 0, 'yaml', '初始化模板', ''),
       (5, 1712995716, 1719560128, 5,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.HttpPort}\n    timeout: ${UserCenterServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.GrpcPort}\n    timeout: ${UserCenterServer.Timeout}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${UserCenterDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${UserCenterDatabase.Username}\n      password: ${UserCenterDatabase.Password}\n      host: ${UserCenterDatabase.Host}\n      port: ${UserCenterDatabase.Port}\n      dbName: ${UserCenterDatabase.Database}\n      option: ${UserCenterDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${ClientJwt.Secret}\n  expire: ${ClientJwt.Expire}\n  renewal: ${ClientJwt.Renewal}\n  whitelist: ${ClientJwt.Whitelist}\nbusiness:\n  service:\n    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n',
        'DE24124137B1', 0, 'yaml', '初始化模板', ''),
       (6, 1712995716, 1712995716, 6,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${PartyAffairsServer.Host}\n    port: ${PartyAffairsServer.HttpPort}\n    timeout: ${PartyAffairsServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${PartyAffairsServer.Host}\n    port: ${PartyAffairsServer.GrpcPort}\n    timeout: ${PartyAffairsServer.Timeout}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${PartyAffairsDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${PartyAffairsDatabase.Username}\n      password: ${PartyAffairsDatabase.Password}\n      host: ${PartyAffairsDatabase.Host}\n      port: ${PartyAffairsDatabase.Port}\n      dbName: ${PartyAffairsDatabase.Database}\n      option: ${PartyAffairsDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n',
        'D825BF2B6742', 1, 'yaml', '初始化模板', ''),
       (7, 1713011130, 1719473730, 2,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: false\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: ${AuthSkipRoles}\nbusiness:\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        '3ECF4F0622F7', 0, 'yaml', '1', ''),
       (8, 1714750273, 1719495060, 3,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  storage:\n    type: local\n    serverPath: /resource/v1/static\n    localDir: static\n    maxSingularSize: ${MaxSingularSize}\n    maxChunkSize: ${MaxChunkSize}\n    maxChunkCount: ${MaxChunkCount}\n    acceptTypes: ${AcceptTypes}\n  export:\n    localDir: static/export\n    expire: 72h\n',
        '3942130C8BEE', 0, 'yaml', '初始化模板', ''),
       (10, 1719379877, 1719473730, 2,
        'test: 11\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: false\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: ${AuthSkipRoles}\nbusiness:\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        'A8A826623534726F8ED1F0AC6ACDFB8B', 0, 'yaml', '1',
        '[{\"type\":\"add\",\"key\":\"test\",\"old\":\"\",\"cur\":\"11\"}]'),
       (11, 1719462171, 1719560089, 1,
        'addr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /usercenter/api/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /cron/api/*\n    timeout: ${CronServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${CronServer.Host}:${CronServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n          method: POST',
        'CE6DA7E554780B619C4C5E8EBD11E7FB', 0, 'yaml', '初始化模板',
        '[{\"type\":\"del\",\"key\":\"debug\",\"old\":\"true\",\"cur\":\"\"},{\"type\":\"update\",\"key\":\"endpoints\",\"old\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /configure/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/v1/static/*\\n  path: /resource/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /cron/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /user-center/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /user-center/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /party-affairs/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /party-affairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\",\"cur\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\"}]'),
       (12, 1719463020, 1719473730, 2,
        'test: 11\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: ${AuthSkipRoles}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\nbusiness:\n  changePasswordType: ${ChangePasswordType}\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        'B0041A124D61345FF898A0019133ABF2', 0, 'yaml', '初始化模板',
        '[{\"type\":\"add\",\"key\":\"client\",\"old\":\"\",\"cur\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n\"},{\"type\":\"update\",\"key\":\"email\",\"old\":\"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha:\\n        from: 青岑云科技\\n        path: static/template/email/default.html\\n        subject: 验证码发送通知\\n        type: text/html\\nuser: ${Email.Username}\\n\",\"cur\":\"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha:\\n        from: 统一应用管理中心\\n        path: static/template/email/default.html\\n        subject: 验证码发送通知\\n        type: text/html\\nuser: ${Email.Username}\\n\"},{\"type\":\"update\",\"key\":\"authentication\",\"old\":\"db: system\\nredis: cache\\nroleKey: role_keyword\\nskipRole: ${AuthSkipRoles}\\n\",\"cur\":\"db: system\\nredis: cache\\nroleKey: roleKeyword\\nskipRole: ${AuthSkipRoles}\\n\"},{\"type\":\"update\",\"key\":\"database\",\"old\":\"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: false\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n\",\"cur\":\"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n\"},{\"type\":\"update\",\"key\":\"business\",\"old\":\"defaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n\",\"cur\":\"changePasswordType: ${ChangePasswordType}\\ndefaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n\"}]'),
       (13, 1719464064, 1719473730, 2,
        'test: 11\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: ${AuthSkipRoles}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\nbusiness:\n  changePasswordType: ${ChangePasswordType}\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        '794DBE4E2B65F3CA34F0CCF465169E87', 0, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"log\",\"old\":\"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\",\"cur\":\"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\"}]'),
       (14, 1719465442, 1719495060, 3,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\nbusiness:\n  chunkSize: ${ChunkSize}\n  defaultMaxSize: ${DefaultMaxSize}\n  defaultAcceptTypes: ${DefaultAcceptTypes}\n  storage:\n    type: local\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        '26E885D17255295754DFB7B3574A30E9', 0, 'yaml', '初始化模板',
        '[{\"type\":\"add\",\"key\":\"redis\",\"old\":\"\",\"cur\":\"cache:\\n    enable: true\\n    host: ${Redis.Host}:${Redis.Port}\\n    password: ${Redis.Password}\\n    username: ${Redis.Username}\\n\"},{\"type\":\"update\",\"key\":\"business\",\"old\":\"export:\\n    expire: 72h\\n    localDir: static/export\\nstorage:\\n    acceptTypes: ${AcceptTypes}\\n    localDir: static\\n    maxChunkCount: ${MaxChunkCount}\\n    maxChunkSize: ${MaxChunkSize}\\n    maxSingularSize: ${MaxSingularSize}\\n    serverPath: /resource/v1/static\\n    type: local\\n\",\"cur\":\"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n\"},{\"type\":\"update\",\"key\":\"log\",\"old\":\"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\",\"cur\":\"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\"}]'),
       (15, 1719466300, 1719495060, 3,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\nbusiness:\n  chunkSize: ${ChunkSize}\n  defaultMaxSize: ${DefaultMaxSize}\n  defaultAcceptTypes: ${DefaultAcceptTypes}\n  storage:\n    type: local\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        'C2E8D51753B9C34DBC56F80137C492A7', 0, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"business\",\"old\":\"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n\",\"cur\":\"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n\"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`,
                        `description`, `compare`)
VALUES (16, 1719466963, 1719560089, 1,
        'addr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /usercenter/api/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /cron/api/*\n    timeout: ${CronServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${CronServer.Host}:${CronServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n          method: POST',
        'D810212421E6DB836EA3CEA6C2E6DCFF', 0, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"endpoints\",\"old\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\",\"cur\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\"}]'),
       (17, 1719473730, 1719473730, 2,
        'test: 11\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 统一应用管理中心\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: roleKeyword\n  skipRole: ${AuthSkipRoles}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\nbusiness:\n  changePasswordType: ${ChangePasswordType}\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n',
        'D69ED17B11BA6B9C08A906818FD83937', 1, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"client\",\"old\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n\",\"cur\":\"- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n\"}]'),
       (18, 1719495060, 1719495060, 3,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\nbusiness:\n  chunkSize: ${ChunkSize}\n  defaultMaxSize: ${DefaultMaxSize}\n  defaultAcceptTypes: ${DefaultAcceptTypes}\n  storage:\n    type: local\n    serverUrl: http://gw.qlime.cn/resource/api/v1/static\n    localDir: static\n    temporaryExpire: 600s\n    secret: limescloud\n  export:\n    serverUrl: http://gw.qlime.cn/resource/api/v1/download\n    localDir: static/export\n    expire: 72h\n',
        'D53BEDB76CF1EA48C689C19F8D28684D', 1, 'yaml', '修改资源url',
        '[{\"type\":\"update\",\"key\":\"business\",\"old\":\"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n\",\"cur\":\"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://gw.qlime.cn/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://gw.qlime.cn/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n\"}]'),
       (19, 1719557201, 1719560089, 1,
        'addr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n          method: POST\n  - path: /usercenter/api/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /cron/api/*\n    timeout: ${CronServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${CronServer.Host}:${CronServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n          method: POST',
        'D0E935C5032DC43B988E70406A7B6111', 0, 'yaml', '修改auth',
        '[{\"type\":\"update\",\"key\":\"endpoints\",\"old\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\",\"cur\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\"}]'),
       (20, 1719560089, 1719560089, 1,
        'addr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \' *\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/api/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n          method: POST\n  - path: /configure/api/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /resource/api/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/api/v1/static/*\n              method: GET\n            - path: /resource/api/v1/download/*\n              method: GET\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n          method: POST\n  - path: /usercenter/api/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /usercenter/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /cron/api/*\n    timeout: ${CronServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${CronServer.Host}:${CronServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/api/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n          method: POST\n  - path: /partyaffairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n          method: POST',
        'CD650FA575AAE93B796D3C395E324C48', 1, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"endpoints\",\"old\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\",\"cur\":\"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n\"}]'),
       (21, 1719560128, 1719560128, 5,
        '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.HttpPort}\n    timeout: ${UserCenterServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.GrpcPort}\n    timeout: ${UserCenterServer.Timeout}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${UserCenterDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${UserCenterDatabase.Username}\n      password: ${UserCenterDatabase.Password}\n      host: ${UserCenterDatabase.Host}\n      port: ${UserCenterDatabase.Port}\n      dbName: ${UserCenterDatabase.Database}\n      option: ${UserCenterDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${ClientJwt.Secret}\n  expire: ${ClientJwt.Expire}\n  renewal: ${ClientJwt.Renewal}\n  whitelist: ${ClientJwt.Whitelist}\nbusiness:\n  defaultPasswordExpire: 10s\n  defaultNickName: 用户-\n',
        '0D01EEE78E5662F3193361C9EEE34266', 1, 'yaml', '初始化模板',
        '[{\"type\":\"update\",\"key\":\"log\",\"old\":\"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\",\"cur\":\"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\"},{\"type\":\"update\",\"key\":\"business\",\"old\":\"service:\\n    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\\n\",\"cur\":\"defaultNickName: 用户-\\ndefaultPasswordExpire: 10s\\n\"}]'),
       (22, 1719737246, 1719737246, 4,
        'env: ${Env.Keyword}\nserver:\n  http:\n    host: ${CronServer.Host}\n    port: ${CronServer.HttpPort}\n    timeout: ${CronServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${CronServer.Host}\n    port: ${CronServer.GrpcPort}\n    timeout: ${CronServer.Timeout}\nlog:\n  level: 0\n  caller: true\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: true\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${CronDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${CronDatabase.Username}\n      password: ${CronDatabase.Password}\n      host: ${CronDatabase.Host}\n      port: ${CronDatabase.Port}\n      dbName: ${CronDatabase.Database}\n      option: ${CronDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\npool: #并非池配置\n  size: 10000 #最大协程数量\n  expiryDuration: 30s #过期时间\n  preAlloc: true #是否预分配\n  maxBlockingTasks: 1000 #最大的并发任务\n  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务',
        '19A024026FF790E6D1AEF036D7D5FA85', 1, 'yaml', '初始化模板',
        '[{\"type\":\"add\",\"key\":\"redis\",\"old\":\"\",\"cur\":\"cache:\\n    enable: true\\n    host: ${Redis.Host}:${Redis.Port}\\n    password: ${Redis.Password}\\n    username: ${Redis.Username}\\n\"},{\"type\":\"del\",\"key\":\"business\",\"old\":\"key: value\\n\",\"cur\":\"\"},{\"type\":\"update\",\"key\":\"log\",\"old\":\"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\",\"cur\":\"caller: true\\nfile:\\n    compress: true\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n\"},{\"type\":\"update\",\"key\":\"database\",\"old\":\"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 3\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${CronDatabase.Database}\\n        host: ${CronDatabase.Host}\\n        option: ${CronDatabase.Option}\\n        password: ${CronDatabase.Password}\\n        port: ${CronDatabase.Port}\\n        username: ${CronDatabase.Username}\\n    drive: ${CronDatabase.Type}\\n    enable: true\\n\",\"cur\":\"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${CronDatabase.Database}\\n        host: ${CronDatabase.Host}\\n        option: ${CronDatabase.Option}\\n        password: ${CronDatabase.Password}\\n        port: ${CronDatabase.Port}\\n        username: ${CronDatabase.Username}\\n    drive: ${CronDatabase.Type}\\n    enable: true\\n\"}]');

--
-- 限制导出的表
--

--
-- 限制表 `business`
--
ALTER TABLE `business`
    ADD CONSTRAINT `fk_business_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE;

--
-- 限制表 `business_value`
--
ALTER TABLE `business_value`
    ADD CONSTRAINT `fk_business_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_business_values` FOREIGN KEY (`business_id`) REFERENCES `business` (`id`) ON
DELETE
CASCADE;

--
-- 限制表 `configure`
--
ALTER TABLE `configure`
    ADD CONSTRAINT `fk_configure_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_configure_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON
DELETE
CASCADE;

--
-- 限制表 `resource_server`
--
ALTER TABLE `resource_server`
    ADD CONSTRAINT `fk_resource_resource_server` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_resource_server_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON
DELETE
CASCADE;

--
-- 限制表 `resource_value`
--
ALTER TABLE `resource_value`
    ADD CONSTRAINT `fk_resource_resource_value` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_resource_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON
DELETE
CASCADE;

--
-- 限制表 `template`
--
ALTER TABLE `template`
    ADD CONSTRAINT `fk_template_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
