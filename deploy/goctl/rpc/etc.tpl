Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 0.0.0.0:2379
  Key: {{.serviceName}}.rpc
