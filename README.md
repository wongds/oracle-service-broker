# oracle-service-broker
oracle service broker for kubernetes


## Start Service Broker
```
bee run -gendoc=true -downdoc=false
```

## 证书名称和证书内容示例
```
connect_uri
system/123456@192.168.250.37:1521/orcl
```

## 在Oracle下的验证方法
```
使用创建的用户登录Oracle, 然后执行如下语句即可看到结果
select tablespace_name, max_bytes/1024/1024 from user_ts_quotas;
select username, default_tablespace from user_users;
```