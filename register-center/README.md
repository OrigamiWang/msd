# register-center

## heart beat 机制
微服务每隔30s，通过kafka和注册中心交流一次
注册中心每隔30s，查看正在运行的微服务
如果超过1分钟没有心跳，视作微服务关闭 (redis expire time is 1 minute)