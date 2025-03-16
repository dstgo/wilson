SET FOREIGN_KEY_CHECKS = 0;

INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (1, 1712995716, 1712995716, 2, 'AuthSkipRoles', 'object', '跳过权限检测的角色列表');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (2, 1712995716, 1712995716, 2, 'DefaultUserPassword', 'string', '默认账号密码');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (4, 1712995716, 1712995716, 2, 'LoginPrivatePath', 'string', 'rsa解密私钥路径');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (5, 1712995716, 1712995716, 2, 'Setting', 'object', '系统设置');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (6, 1712995716, 1719465188, 3, 'DefaultAcceptTypes', 'object', '允许上传的文件后缀');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (8, 1712995716, 1719465209, 3, 'ChunkSize', 'int', '单个切片最大大小（M）');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (9, 1712995716, 1719465230, 3, 'DefaultMaxSize', 'int', '文件最大大小（M）');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (15, 1719462652, 1719462652, 2, 'ChangePasswordType', 'string', '修改密码方式');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (16, 1741937485, 1741937485, 2, 'LoginCaptcha', 'object', '登陆验证码配置');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (17, 1742129509, 1742129692, 2, 'ChangePasswordCaptcha', 'object', '修改密码验证配置');
INSERT INTO `business` (`id`, `created_at`, `updated_at`, `server_id`, `keyword`, `type`, `description`) VALUES (18, 1742131634, 1742131634, 2, 'EmailCaptcha', 'object', '验证码邮件配置');

INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (52, 1719371012, 1719465259, 1, 6, '["jpg","png","txt","ppt","pptx","mp4","pdf"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (53, 1719371012, 1719465259, 2, 6, '["jpg","png","txt","ppt","pptx","mp4"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (54, 1719371012, 1719465259, 3, 6, '["jpg","png","txt","ppt","pptx","mp4"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (61, 1719382347, 1719465239, 1, 9, '10');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (62, 1719382347, 1719465239, 2, 9, '10');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (63, 1719382347, 1719465239, 3, 9, '10');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (64, 1719382352, 1719382352, 1, 8, '1');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (65, 1719382352, 1719382352, 2, 8, '1');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (66, 1719382352, 1719382352, 3, 8, '1');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (73, 1719462662, 1719462662, 1, 15, 'password');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (74, 1719462662, 1719462662, 2, 15, 'password');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (75, 1719462662, 1719462662, 3, 15, 'email');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (76, 1719463071, 1742131481, 1, 4, '/etc/wilson/manager/static/cert/login.pem');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (77, 1719463071, 1742131481, 2, 4, '/etc/wilson/manager/static/cert/login.pem');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (78, 1719463071, 1742131481, 3, 4, '/etc/wilson/manager/static/cert/login.pem');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (79, 1719463092, 1719463092, 1, 1, '["superAdmin"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (80, 1719463092, 1719463092, 2, 1, '["superAdmin"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (81, 1719463092, 1719463092, 3, 1, '["superAdmin"]');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (82, 1719463137, 1719463137, 1, 2, '12345678');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (83, 1719463137, 1719463137, 2, 2, '12345678');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (84, 1719463137, 1719463137, 3, 2, '12345678');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (85, 1719463217, 1742051415, 1, 5, '{}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (86, 1719463217, 1742051415, 2, 5, '{}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (87, 1719463217, 1742051415, 3, 5, '{}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (88, 1741937505, 1741937505, 1, 16, '{"Type":"image","Length":6,"Expire":"180s","Redis":"cache","Skew":0.7,"Refresh":true,"DotCount":80,"Height":80,"Width":240}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (89, 1741937505, 1741937505, 2, 16, '{"Type":"image","Length":6,"Expire":"180s","Redis":"cache","Height":80,"Skew":0.7,"Width":240,"Refresh":true,"DotCount":80}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (90, 1741937505, 1741937505, 3, 16, '{"Type":"image","Length":6,"Height":80,"Width":240,"Expire":"180s","Redis":"cache","Skew":0.7,"Refresh":true,"DotCount":80}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (95, 1742129702, 1742129702, 1, 17, '{"type":"email","length":6,"expire":"180s","redis":"cache","template":"captcha"}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (96, 1742129702, 1742129702, 2, 17, '{"type":"email","length":6,"expire":"180s","redis":"cache","template":"captcha"}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (97, 1742129702, 1742129702, 3, 17, '{"template":"captcha","type":"email","length":6,"expire":"180s","redis":"cache"}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (101, 1742131697, 1742131697, 1, 18, '{"subject":"验证码发送通知","path":"/etc/wilson/manager/static/template/email/default.html","from":"饥荒管理平台","type":"text/html"}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (102, 1742131697, 1742131697, 2, 18, '{"subject":"验证码发送通知","path":"/etc/wilson/manager/static/template/email/default.html","from":"饥荒管理平台","type":"text/html"}');
INSERT INTO `business_value` (`id`, `created_at`, `updated_at`, `env_id`, `business_id`, `value`) VALUES (103, 1742131697, 1742131697, 3, 18, '{"subject":"验证码发送通知","path":"/etc/wilson/manager/static/template/email/default.html","from":"饥荒管理平台","type":"text/html"}');


INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (1, 1712995716, 1742129200, 1, 1, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET', 'e24b596ed7d9258ff11df8f1af73540e', 'yaml', '精简了网关配置');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (2, 1712995716, 1742133054, 2, 1, 'env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: debug
  encode: json
  caller: true
  output: ["stdout","file"]
  file: {"Compress":false,"Name":"/etc/wilson/manager/log/output.log","MaxSize":1,"MaxBackup":5,"MaxAge":1}
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: root
      password: 123456
      host: 127.0.0.1
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password: 123456
captcha:
  login: {"Type":"image","Length":6,"Expire":"180s","Redis":"cache","Skew":0.7,"Refresh":true,"DotCount":80,"Height":80,"Width":240}
  changePassword: {"type":"email","length":6,"expire":"180s","redis":"cache","template":"captcha"}
loader:
  login: /etc/wilson/manager/static/cert/login.pem
email:
  template:
    captcha: {"subject":"验证码发送通知","path":"/etc/wilson/manager/static/template/email/default.html","from":"饥荒管理平台","type":"text/html"}
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: dstgo-panel
  expire: 2h
  renewal: 600s
  whitelist: {"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true,"GET:/manager/api/v1/system/setting":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
manager:
  changePasswordType: password
  defaultUserPassword: 12345678
  setting: {}
', 'e20420d9409558b8f969588ee8d54754', 'yaml', '同步一致的jwt密钥');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (3, 1712995716, 1742131133, 3, 1, '
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 10s
log:
  level: debug
  encode: json
  caller: true
  output: ["stdout","file"]
  file: {"Name":"/etc/wilson/resource/log/output.log","MaxSize":1,"MaxBackup":5,"MaxAge":1,"Compress":false}
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: root
      password: 123456
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password: 123456
resource:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4","pdf"]
  storage:
    type: local
    serverUrl: http://127.0.0.1:7080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:7080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'c3ad4b56d5a8e8fd8a258a8934a7799c', 'yaml', '重命名资源中心局部配置字段');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (6, 1712995716, 1719557335, 1, 2, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /manager/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
  - path: /cron/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7040
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/api/v1/auth
          method: POST', 'e25350369a0436f3af879333dcb7cbe5', 'yaml', '初始化模板');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (7, 1712995716, 1719464076, 2, 2, 'test: 11
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: static/cert/login.pem
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: xxx
jwt:
  redis: cache
  secret: limes-cloud-pre
  expire: 2h
  renewal: 600s
  whitelist: {"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8010
business:
  changePasswordType: password
  defaultUserPassword: 12345678
  setting: {"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}
', '1d2555ef65b788c69f3254a92719cc0b', 'yaml', '初始化模板');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (8, 1712995716, 1719467313, 3, 2, '
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
business:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4"]
  storage:
    type: local
    serverUrl: http://127.0.0.1:7080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:7080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '84cdc6dbc3e7e5d04af6f5f1da605013', 'yaml', '初始化模板');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (11, 1712995716, 1729189107, 1, 3, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /manager/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 600s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: 600s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /application/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /application/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
  - path: /cron/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7040
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST', '66164cc6f6d0a50cdfb1035b6672a706', 'yaml', '1');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (12, 1712995716, 1720364177, 2, 3, 'test: 11
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: manager
      password: L8hjTy5GMZmdHJX3
      host: 127.0.0.1
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: static/cert/login.pem
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: limes-cloud-prod
  expire: 2h
  renewal: 600s
  whitelist: {"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
business:
  changePasswordType: email
  defaultUserPassword: 12345678
  setting: {"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}
', '42462633dd9a6685e39fb3572c9fd01e', 'yaml', '初始化模板');
INSERT INTO `configure` (`id`, `created_at`, `updated_at`, `server_id`, `env_id`, `content`, `version`, `format`, `description`) VALUES (13, 1712995716, 1729189119, 3, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 600s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 600s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: resource
      password: ct7AYfaM8kc8LWHi
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
business:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4"]
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'de69d97e91f4e71bfdc3ad1e63f85976', 'yaml', '1');

INSERT INTO `env` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `token`, `status`) VALUES (1, 1741932120, 1742130467, 'TEST', '测试环境', '用于本地测试', '1E2AFDC5E7794BF0095E4FEC52F1C9B0474CB2B609A9C58F4FC7C3C7134E649C', 1);
INSERT INTO `env` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `token`, `status`) VALUES (2, 1741932120, 1741932120, 'PRE', '预发布环境', '用于上线前测试', '862BBE3D5BE34A780305DA84A8DD5147', 1);
INSERT INTO `env` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `token`, `status`) VALUES (3, 1741932120, 1741932120, 'PROD', '生产环境', '用于线上真实环境', '5B655B7D4A51BF79C974C9F27C27D992', 1);

INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (1, 1712995716, 1712995716, 'Env', '环境标识信息', 'Keyword', 'env', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (2, 1712995716, 1712995716, 'AdminJwt', '后台管理服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (4, 1712995716, 1712995716, 'Redis', 'redis中间件配置信息', 'Host,Username,Password,Port', 'redis', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (5, 1712995716, 1742049998, 'Email', '邮箱服务配置信息', 'Username,Company,Password,Host,Port', 'email', 1);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (6, 1712995716, 1712995716, 'GatewayServer', '通用网关服务配置信息', 'HttpPort,Host,Timeout', 'server', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (7, 1712995716, 1741934431, 'ManagerServer', '管理中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (8, 1712995716, 1712995716, 'ManagerDatabase', '管理中心数据库配置', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (9, 1712995716, 1712995716, 'ResourceServer', '资源中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (10, 1712995716, 1712995716, 'ResourceDatabase', '资源中心数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (13, 1712995716, 1712995716, 'ConfigureServer', '配置中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (21, 1741934568, 1741938546, 'ManagerLog', '管理中心日志配置', 'Level,Caller,Encode,Output,File', 'log', 1);
INSERT INTO `resource` (`id`, `created_at`, `updated_at`, `keyword`, `description`, `fields`, `tag`, `private`) VALUES (22, 1742050788, 1742050788, 'ResourceLog', '资源中心日志配置', 'Caller,File,Output,Level,Encode', 'log', 1);

INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (1, 1712995716, 2, 8);
INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (2, 1712995716, 3, 10);
INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (4, null, 2, 7);
INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (13, null, 2, 21);
INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (14, null, 2, 5);
INSERT INTO `resource_server` (`id`, `created_at`, `server_id`, `resource_id`) VALUES (15, null, 3, 22);

INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (1, 1712995716, 1719371433, 1, 1, '{"Keyword":"TEST"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (2, 1712995716, 1719371433, 2, 1, '{"Keyword":"PRE"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (3, 1712995716, 1719371433, 3, 1, '{"Keyword":"PROD"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (4, 1712995716, 1742132098, 1, 2, '{"Secret":"dstgo-panel","Expire":"2h","Renewal":"600s","Whitelist":{"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true}}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (5, 1712995716, 1742132098, 2, 2, '{"Expire":"2h","Renewal":"600s","Whitelist":{"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true},"Secret":"dstgo-panel"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (6, 1712995716, 1742132098, 3, 2, '{"Renewal":"600s","Whitelist":{"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true},"Secret":"dstgo-panel","Expire":"2h"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (10, 1712995716, 1742051263, 1, 4, '{"Host":"127.0.0.1","Username":"","Password":"123456","Port":6379}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (11, 1712995716, 1742051263, 2, 4, '{"Password":"","Port":6379,"Host":"127.0.0.1","Username":""}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (12, 1712995716, 1742051263, 3, 4, '{"Host":"127.0.0.1","Username":"","Password":"","Port":6379}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (13, 1712995716, 1714884855, 1, 5, '{"Username":"860808187@qq.com","Company":"青岑云","Password":"fyudafdzqmhwbfbd","Host":"smtp.qq.com","Port":25}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (14, 1712995716, 1714884855, 2, 5, '{"Password":"xxx","Host":"smtp.qq.com","Port":25,"Username":"860808187@qq.com","Company":"青岑云"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (15, 1712995716, 1714884855, 3, 5, '{"Username":"860808187@qq.com","Company":"青岑云","Password":"fyudafdzqmhwbfbd","Host":"smtp.qq.com","Port":25}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (16, 1712995716, 1712995716, 1, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (17, 1712995716, 1712995716, 2, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (18, 1712995716, 1712995716, 3, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (19, 1712995716, 1741938745, 1, 7, '{"Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (20, 1712995716, 1741938745, 2, 7, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (21, 1712995716, 1741938745, 3, 7, '{"Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (22, 1712995716, 1742051282, 1, 8, '{"Port":"3306","Type":"mysql","Database":"manager","Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456","Host":"127.0.0.1"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (23, 1712995716, 1742051282, 2, 8, '{"Port":"3306","Type":"mysql","Database":"manager","Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456","Host":"127.0.0.1"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (24, 1712995716, 1742051282, 3, 8, '{"Port":"3306","Type":"mysql","Database":"manager","Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456","Host":"127.0.0.1"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (25, 1712995716, 1729189096, 1, 9, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7020,"GrpcPort":8020}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (26, 1712995716, 1729189096, 2, 9, '{"GrpcPort":8020,"Timeout":"10s","Host":"127.0.0.1","HttpPort":7020}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (27, 1712995716, 1729189096, 3, 9, '{"Timeout":"600s","Host":"127.0.0.1","HttpPort":7020,"GrpcPort":8020}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (28, 1712995716, 1742051297, 1, 10, '{"Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"resource","Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (29, 1712995716, 1742051297, 2, 10, '{"Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"resource"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (30, 1712995716, 1742051297, 3, 10, '{"Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"resource","Option":"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local","Username":"root","Password":"123456"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (37, 1712995716, 1712995716, 1, 13, '{"Host":"127.0.0.1","HttpPort":6081,"GrpcPort":6082,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (38, 1712995716, 1712995716, 2, 13, '{"Host":"127.0.0.1","HttpPort":6081,"GrpcPort":6082,"Timeout":"10s"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (39, 1712995716, 1712995716, 3, 13, '{"GrpcPort":6082,"Timeout":"10s","Host":"127.0.0.1","HttpPort":6081}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (131, 1741934810, 1741938620, 1, 21, '{"Level":"debug","Output":["stdout","file"],"File":{"MaxSize":1,"MaxBackup":5,"MaxAge":1,"Compress":false,"Name":"/etc/wilson/manager/log/output.log"},"Caller":true,"Encode":"json"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (132, 1741934810, 1741938620, 2, 21, '{"Level":"debug","Output":["stdout","file"],"File":{"Compress":false,"Name":"/etc/wilson/manager/log/output.log","MaxSize":1,"MaxBackup":5,"MaxAge":1},"Caller":true,"Encode":"json"}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (133, 1741934810, 1741938620, 3, 21, '{"Caller":true,"Encode":"json","Level":"debug","Output":["stdout","file"],"File":{"Name":"/etc/wilson/manager/log/output.log","MaxSize":1,"MaxBackup":5,"MaxAge":1,"Compress":false}}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (141, 1742050799, 1742050814, 1, 22, '{"Encode":"json","Level":"debug","Output":["stdout","file"],"File":{"MaxSize":1,"MaxBackup":5,"MaxAge":1,"Compress":false,"Name":"/etc/wilson/resource/log/output.log"},"Caller":true}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (142, 1742050799, 1742050814, 2, 22, '{"File":{"MaxSize":1,"MaxBackup":5,"MaxAge":1,"Compress":false,"Name":"/etc/wilson/resource/log/output.log"},"Caller":true,"Encode":"json","Level":"debug","Output":["stdout","file"]}');
INSERT INTO `resource_value` (`id`, `created_at`, `updated_at`, `env_id`, `resource_id`, `value`) VALUES (143, 1742050799, 1742050814, 3, 22, '{"Output":["stdout","file"],"File":{"MaxAge":1,"Compress":false,"Name":"/etc/wilson/resource/log/output.log","MaxSize":1,"MaxBackup":5},"Caller":true,"Encode":"json","Level":"debug"}');

INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (1, 1712995716, 1742130807, 'gateway', '通用网关', '主要负责前端到后端的转发', 1);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (2, 1712995716, 1742130815, 'manager', '管理中心', '主要负责系统的基础管理', 1);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (3, 1712995716, 1742130788, 'resource', '资源中心', '主要负责静态资源的管理', 1);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (4, 1742133337, 1742133337, 'game', '游戏中心', '负责提供饥荒游戏数据', null);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (5, 1742133374, 1742133374, 'container', '容器中心', '负责所有节点的容器管理', null);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (6, 1742133415, 1742133415, 'control', '控制中心', '负责转发游戏节点请求', null);
INSERT INTO `server` (`id`, `created_at`, `updated_at`, `keyword`, `name`, `description`, `status`) VALUES (7, 1742133468, 1742133468, 'daemon', '游戏节点', '物理游戏节点，负责处理游戏业务', null);

INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (1, 1712995716, 1742129149, 1, '
debug: true
addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/v1/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /configure/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /resource/v1/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
          whiteList:
            - path: /resource/v1/static/*
              method: GET
  - path: /cron/v1/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /user-center/v1/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /user-center/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /party-affairs/v1/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /party-affairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
', '54E50FB3522C', 0, 'yaml', '初始化模板', '');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (2, 1712995716, 1742131859, 2, '
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '15768C4C6F57', 0, 'yaml', '初始化模板', '');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (3, 1712995716, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
file:
  storage: local
  serverPath: /resource/v1/static
  localDir: static
  maxSingularSize: ${MaxSingularSize}
  maxChunkSize: ${MaxChunkSize}
  maxChunkCount: ${MaxChunkCount}
  acceptTypes: ${AcceptTypes}

', '56945B81FA4D', 0, 'yaml', '初始化模板', '');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (7, 1713011130, 1742131859, 2, '
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '3ECF4F0622F7', 0, 'yaml', '1', '');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (8, 1714750273, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
business:
  storage:
    type: local
    serverPath: /resource/v1/static
    localDir: static
    maxSingularSize: ${MaxSingularSize}
    maxChunkSize: ${MaxChunkSize}
    maxChunkCount: ${MaxChunkCount}
    acceptTypes: ${AcceptTypes}
  export:
    localDir: static/export
    expire: 72h
', '3942130C8BEE', 0, 'yaml', '初始化模板', '');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (10, 1719379877, 1742131859, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'A8A826623534726F8ED1F0AC6ACDFB8B', 0, 'yaml', '1', '[{"type":"add","key":"test","old":"","cur":"11"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (11, 1719462171, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'CE6DA7E554780B619C4C5E8EBD11E7FB', 0, 'yaml', '初始化模板', '[{"type":"del","key":"debug","old":"true","cur":""},{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /configure/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/v1/static/*\\n  path: /resource/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /cron/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /user-center/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /user-center/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\\n  path: /party-affairs/v1/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /party-affairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (12, 1719463020, 1742131859, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'B0041A124D61345FF898A0019133ABF2', 0, 'yaml', '初始化模板', '[{"type":"add","key":"client","old":"","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n"},{"type":"update","key":"email","old":"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha:\\n        from: 青岑云科技\\n        path: static/template/email/default.html\\n        subject: 验证码发送通知\\n        type: text/html\\nuser: ${Email.Username}\\n","cur":"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha:\\n        from: 统一应用管理中心\\n        path: static/template/email/default.html\\n        subject: 验证码发送通知\\n        type: text/html\\nuser: ${Email.Username}\\n"},{"type":"update","key":"authentication","old":"db: system\\nredis: cache\\nroleKey: role_keyword\\nskipRole: ${AuthSkipRoles}\\n","cur":"db: system\\nredis: cache\\nroleKey: roleKeyword\\nskipRole: ${AuthSkipRoles}\\n"},{"type":"update","key":"database","old":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: false\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n","cur":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n"},{"type":"update","key":"business","old":"defaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n","cur":"changePasswordType: ${ChangePasswordType}\\ndefaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (13, 1719464064, 1742131859, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '794DBE4E2B65F3CA34F0CCF465169E87', 0, 'yaml', '初始化模板', '[{"type":"update","key":"log","old":"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n","cur":"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (14, 1719465442, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '26E885D17255295754DFB7B3574A30E9', 0, 'yaml', '初始化模板', '[{"type":"add","key":"redis","old":"","cur":"cache:\\n    enable: true\\n    host: ${Redis.Host}:${Redis.Port}\\n    password: ${Redis.Password}\\n    username: ${Redis.Username}\\n"},{"type":"update","key":"business","old":"export:\\n    expire: 72h\\n    localDir: static/export\\nstorage:\\n    acceptTypes: ${AcceptTypes}\\n    localDir: static\\n    maxChunkCount: ${MaxChunkCount}\\n    maxChunkSize: ${MaxChunkSize}\\n    maxSingularSize: ${MaxSingularSize}\\n    serverPath: /resource/v1/static\\n    type: local\\n","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"},{"type":"update","key":"log","old":"file:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n","cur":"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (15, 1719466300, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'C2E8D51753B9C34DBC56F80137C492A7', 0, 'yaml', '初始化模板', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (16, 1719466963, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'D810212421E6DB836EA3CEA6C2E6DCFF', 0, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (17, 1719473730, 1742131859, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'D69ED17B11BA6B9C08A906818FD83937', 0, 'yaml', '初始化模板', '[{"type":"update","key":"client","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n","cur":"- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\\n  server: Resource\\n  type: direct\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (18, 1719557318, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'D0E935C5032DC43B988E70406A7B6111', 0, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (20, 1719559466, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /usercenter/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'DEBD584379CAA5B33A00199457BF87DC', 0, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /usercenter/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (21, 1719559979, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'CD650FA575AAE93B796D3C395E324C48', 0, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /usercenter/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (25, 1720323538, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /manager/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', '1802390697E147D536FEFB858CBB09BE', 0, 'yaml', '新增usercenter auth', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (26, 1720323609, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /usercenter/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', '6ED76BB69C66544B238C0F7CF70951FB', 0, 'yaml', '1', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /usercenter/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (27, 1720323742, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST', '2B62DFD7A54233FC42AC86B7FFA9E044', 0, 'yaml', '1', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: POST\\n              path: /usercenter/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (28, 1720366812, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '7124FA23DF2383EFF43F397F75755BFB', 0, 'yaml', '初始化配置', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (30, 1723030749, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST', '634F4B998D8891899B0E918F95BC364B', 0, 'yaml', '新增poverty', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /poverty/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /poverty/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (31, 1728064733, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /application/api/*
    timeout: ${ApplicationServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /application/client/*
    timeout: ${ApplicationServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST', '3020B1B4C417B67FE2480669A6087A34', 0, 'yaml', '新增应用中心', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /usercenter/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\\n  path: /usercenter/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${UserCenterServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /poverty/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\\n  path: /poverty/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /application/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ApplicationServer.Timeout}\\n- backends:\\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\\n  path: /application/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ApplicationServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /poverty/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /poverty/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (34, 1728072395, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'A940188BE20724F72D54E3673D51C7FD', 0, 'yaml', '新增数据库日志', '[{"type":"update","key":"database","old":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 3\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n","cur":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (35, 1741932334, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST', 'E9E91EC0E37CEA16E88A91503B979840', 0, 'yaml', '删除了无用的服务', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /application/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ApplicationServer.Timeout}\\n- backends:\\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\\n  path: /application/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ApplicationServer.Timeout}\\n- backends:\\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /cron/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${CronServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /partyaffairs/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /partyaffairs/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PartyAffairsServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /poverty/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n- backends:\\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /poverty/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${PovertyServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (36, 1741936814, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file:
    name: ${ManagerLog.FileName}
    maxSize: ${ManagerLog.FileMaxSize}
    maxBackup: ${ManagerLog.FileMaxBackup}
    maxAge: ${ManagerLog.FileMaxAge}
    compress: ${ManagerLog.FileCompress}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '61AB75AFD3F00591ACC1C8CC4E18528C', 0, 'yaml', '日志配置', '[{"type":"del","key":"test","old":"11","cur":""},{"type":"update","key":"log","old":"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n","cur":"caller: ${ManagerLog.Caller}\\nfile:\\n    compress: ${ManagerLog.FileCompress}\\n    maxAge: ${ManagerLog.FileMaxAge}\\n    maxBackup: ${ManagerLog.FileMaxBackup}\\n    maxSize: ${ManagerLog.FileMaxSize}\\n    name: ${ManagerLog.FileName}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (37, 1742049649, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '425EDAAB28CDC00D33F142908015E362', 0, 'yaml', '禁用数据库自动创建', '[{"type":"update","key":"log","old":"caller: ${ManagerLog.Caller}\\nfile:\\n    compress: ${ManagerLog.FileCompress}\\n    maxAge: ${ManagerLog.FileMaxAge}\\n    maxBackup: ${ManagerLog.FileMaxBackup}\\n    maxSize: ${ManagerLog.FileMaxSize}\\n    name: ${ManagerLog.FileName}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n","cur":"caller: ${ManagerLog.Caller}\\nfile: ${ManagerLog.File}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n"},{"type":"update","key":"database","old":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n","cur":"system:\\n    autoCreate: false\\n    config:\\n        initializer:\\n            enable: false\\n            path: none\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ManagerDatabase.Database}\\n        host: ${ManagerDatabase.Host}\\n        option: ${ManagerDatabase.Option}\\n        password: ${ManagerDatabase.Password}\\n        port: ${ManagerDatabase.Port}\\n        username: ${ManagerDatabase.Username}\\n    drive: ${ManagerDatabase.Type}\\n    enable: true\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (38, 1742049847, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: fasle
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '24EF3127CC7EA62BF30CC7D5F2EBFAFE', 0, 'yaml', '禁用数据库自动创建', '[{"type":"update","key":"database","old":"system:\\n    autoCreate: true\\n    config:\\n        initializer:\\n            enable: true\\n            path: deploy/data.sql\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n","cur":"system:\\n    autoCreate: false\\n    config:\\n        initializer:\\n            enable: fasle\\n            path: none\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (39, 1742050614, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: fasle
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://127.0.0.1:8080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:8080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'D3F762BD394E8F1012F96F221DA0F0B8', 0, 'yaml', '修改了文件的展示URL', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://127.0.0.1:8080/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://127.0.0.1:8080/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (40, 1742050908, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  ebcide: ${ManagerLog.Encode}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '4EE1DC7AC780698D546872E481D0B4C0', 0, 'yaml', '设置log encode', '[{"type":"update","key":"log","old":"caller: ${ManagerLog.Caller}\\nfile: ${ManagerLog.File}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n","cur":"caller: ${ManagerLog.Caller}\\nebcide: ${ManagerLog.Encode}\\nfile: ${ManagerLog.File}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (41, 1742050970, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: ${ResourceLog.Level}
  ebcide: ${ResourceLog.Encode}
  caller: ${ResourceLog.Caller}
  output: ${ResourceLog.Output}
  file: ${ResourceLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: fasle
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://127.0.0.1:8080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:8080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '75DB2844692CD7515D94631553B95E24', 0, 'yaml', '资源中心日志配置变量', '[{"type":"update","key":"log","old":"caller: true\\nfile:\\n    compress: false\\n    maxAge: 1\\n    maxBackup: 5\\n    maxSize: 1\\n    name: ./tmp/runtime/output.log\\nlevel: 0\\noutput:\\n    - stdout\\n    - file\\n","cur":"caller: ${ResourceLog.Caller}\\nebcide: ${ResourceLog.Encode}\\nfile: ${ResourceLog.File}\\nlevel: ${ResourceLog.Level}\\noutput: ${ResourceLog.Output}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (42, 1742129149, 1742129149, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - \'*\'
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET', '7B478AE5645CEB3388D9145C4D90E49A', 1, 'yaml', '删除了无用的服务', '[{"type":"update","key":"endpoints","old":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /manager/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\\n  path: /resource/client/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n","cur":"- backends:\\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\\n  path: /manager/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ManagerServer.Timeout}\\n- backends:\\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n  path: /configure/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ConfigureServer.Timeout}\\n- backends:\\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\\n  middlewares:\\n    - name: auth\\n      options:\\n        method: POST\\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\\n        whiteList:\\n            - method: GET\\n              path: /resource/api/v1/static/*\\n            - method: GET\\n              path: /resource/api/v1/download/*\\n  path: /resource/api/*\\n  protocol: HTTP\\n  responseFormat: true\\n  timeout: ${ResourceServer.Timeout}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (43, 1742130093, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: ${ResourceLog.Level}
  ebcide: ${ResourceLog.Encode}
  caller: ${ResourceLog.Caller}
  output: ${ResourceLog.Output}
  file: ${ResourceLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: fasle
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'F33CA4BEC4482302117AEF05C3499B34', 0, 'yaml', '修改serverURL地址为网关地址', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://127.0.0.1:8080/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://127.0.0.1:8080/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (44, 1742130895, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: ${ResourceLog.Level}
  ebcide: ${ResourceLog.Encode}
  caller: ${ResourceLog.Caller}
  output: ${ResourceLog.Output}
  file: ${ResourceLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '8E05D0AF8AD5A8E22A19C16DAF1E8DFD', 0, 'yaml', 'fix: 修复了Initializer.Enable拼写错误的问题', '[{"type":"update","key":"database","old":"system:\\n    autoCreate: false\\n    config:\\n        initializer:\\n            enable: fasle\\n            path: none\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n","cur":"system:\\n    autoCreate: false\\n    config:\\n        initializer:\\n            enable: false\\n            path: none\\n        logLevel: 4\\n        maxIdleConn: 10\\n        maxLifetime: 2h\\n        maxOpenConn: 20\\n        slowThreshold: 2s\\n        transformError:\\n            enable: true\\n    connect:\\n        dbName: ${ResourceDatabase.Database}\\n        host: ${ResourceDatabase.Host}\\n        option: ${ResourceDatabase.Option}\\n        password: ${ResourceDatabase.Password}\\n        port: ${ResourceDatabase.Port}\\n        username: ${ResourceDatabase.Username}\\n    drive: ${ResourceDatabase.Type}\\n    enable: true\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (45, 1742130997, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: ${ResourceLog.Level}
  encode: ${ResourceLog.Encode}
  caller: ${ResourceLog.Caller}
  output: ${ResourceLog.Output}
  file: ${ResourceLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '82EA7E37B686B28D61B8962F1F3A122B', 0, 'yaml', 'fix: 修复了日志encode字段拼写错误', '[{"type":"update","key":"log","old":"caller: ${ResourceLog.Caller}\\nebcide: ${ResourceLog.Encode}\\nfile: ${ResourceLog.File}\\nlevel: ${ResourceLog.Level}\\noutput: ${ResourceLog.Output}\\n","cur":"caller: ${ResourceLog.Caller}\\nencode: ${ResourceLog.Encode}\\nfile: ${ResourceLog.File}\\nlevel: ${ResourceLog.Level}\\noutput: ${ResourceLog.Output}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (46, 1742131110, 1742131110, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: ${ResourceLog.Level}
  encode: ${ResourceLog.Encode}
  caller: ${ResourceLog.Caller}
  output: ${ResourceLog.Output}
  file: ${ResourceLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
resource:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'BD5446C3748FA51E85F3E7AE3303955C', 1, 'yaml', '重命名资源中心局部配置字段', '[{"type":"add","key":"resource","old":"","cur":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n"},{"type":"del","key":"business","old":"chunkSize: ${ChunkSize}\\ndefaultAcceptTypes: ${DefaultAcceptTypes}\\ndefaultMaxSize: ${DefaultMaxSize}\\nexport:\\n    expire: 72h\\n    localDir: static/export\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\\nstorage:\\n    localDir: static\\n    secret: limescloud\\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\\n    temporaryExpire: 600s\\n    type: local\\n","cur":""}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (47, 1742131292, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  ebcide: ${ManagerLog.Encode}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login: ${LoginCaptcha}
  changePassword: ${ChangePasswordCaptcha}
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'A4BF6952AEBF24753495DBE04B97D25D', 0, 'yaml', '精简配置', '[{"type":"update","key":"captcha","old":"changePassword:\\n    expire: 180s\\n    length: 6\\n    redis: cache\\n    template: captcha\\n    type: email\\nlogin:\\n    dotCount: 80\\n    expire: 180s\\n    height: 80\\n    length: 6\\n    redis: cache\\n    refresh: true\\n    skew: 0.7\\n    type: image\\n    width: 240\\n","cur":"changePassword: ${ChangePasswordCaptcha}\\nlogin: ${LoginCaptcha}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (48, 1742131371, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  encode: ${ManagerLog.Encode}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login: ${LoginCaptcha}
  changePassword: ${ChangePasswordCaptcha}
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '223E900C458758C1642E3A4709EC9F18', 0, 'yaml', '日志encode字段拼写错误', '[{"type":"update","key":"log","old":"caller: ${ManagerLog.Caller}\\nebcide: ${ManagerLog.Encode}\\nfile: ${ManagerLog.File}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n","cur":"caller: ${ManagerLog.Caller}\\nencode: ${ManagerLog.Encode}\\nfile: ${ManagerLog.File}\\nlevel: ${ManagerLog.Level}\\noutput: ${ManagerLog.Output}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (49, 1742131752, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  encode: ${ManagerLog.Encode}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login: ${LoginCaptcha}
  changePassword: ${ChangePasswordCaptcha}
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha: ${EmailCaptcha}
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '680FA0B242409DF1A2034AD679CC9EB9', 0, 'yaml', '新增验证码配置项', '[{"type":"update","key":"email","old":"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha:\\n        from: 统一应用管理中心\\n        path: static/template/email/default.html\\n        subject: 验证码发送通知\\n        type: text/html\\nuser: ${Email.Username}\\n","cur":"host: ${Email.Host}\\nname: ${Email.Company}\\npassword: ${Email.Password}\\nport: ${Email.Port}\\ntemplate:\\n    captcha: ${EmailCaptcha}\\nuser: ${Email.Username}\\n"}]');
INSERT INTO `template` (`id`, `created_at`, `updated_at`, `server_id`, `content`, `version`, `is_use`, `format`, `description`, `compare`) VALUES (50, 1742131859, 1742131859, 2, 'env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: ${ManagerLog.Level}
  encode: ${ManagerLog.Encode}
  caller: ${ManagerLog.Caller}
  output: ${ManagerLog.Output}
  file: ${ManagerLog.File}
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: false #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: "none"
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login: ${LoginCaptcha}
  changePassword: ${ChangePasswordCaptcha}
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha: ${EmailCaptcha}
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
manager:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'E802AAC4FAF5F2FA66B602CD6638D5DF', 1, 'yaml', 'manager配置项重命名', '[{"type":"add","key":"manager","old":"","cur":"changePasswordType: ${ChangePasswordType}\\ndefaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n"},{"type":"del","key":"business","old":"changePasswordType: ${ChangePasswordType}\\ndefaultUserPassword: ${DefaultUserPassword}\\nsetting: ${Setting}\\n","cur":""}]');


SET FOREIGN_KEY_CHECKS = 1;