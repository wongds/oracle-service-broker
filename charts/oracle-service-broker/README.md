# Oracle Service Broker

用于对接外部Oracle集群服务,遵守ServiceBroker API规范.
[Open Service Broker](https://www.openservicebrokerapi.org/)


更多信息,
[请访问 GitHub 上 Service Catalog 项目](https://github.com/kubernetes-incubator/service-catalog).

## 安装应用
安装用户需要上传已经准备好`json`文件,内容如下:

```
{
	"name": "oracle-service-broker",
	"namespace": "default",
	"repo": "dcos",
	"chart": "oracle-service-broker",
	"version": "0.0.2",
	"values": {
		"image": "neunnsy/oracle-service-broker:v0.0.2",
		"imagePullPolicy": "IfNotPresent",
		"etcdImage": "quay.io/coreos/etcd:v3.0.17"
	}
}

```

## 配置参数

下面表格中的内容列举了用户提供的 ServiceBroker 可配置参数

| 参数列表 | 描述 | 默认值 |
|-----------|-------------|---------|
| `image` | 镜像 | `neunnsy/oracle-service-broker:v0.0.2` |
| `imagePullPolicy` | 镜像拉取规则 | `Always` |
|`etcdImage`|etcd镜像|`quay.io/coreos/etcd:v3.0.17`|

