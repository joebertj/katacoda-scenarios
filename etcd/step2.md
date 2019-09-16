Go inside the etcd-master pod
`kubectl exec -it etcd-master -n kube-system sh`{{execute}}

Set the certificate for etcdctl
```
export ETCDCTL_CA_FILE=/etc/kubernetes/pki/etcd/ca.crt
export ETCDCTL_CERT_FILE=/etc/kubernetes/pki/etcd/server.crt
export ETCDCTL_KEY_FILE=/etc/kubernetes/pki/etcd/server.key
```{{execute}}

Use set get and ls  
`etcdctl --endpoints https://127.0.0.1:2379 set a 1`{{execute}}  
`etcdctl --endpoints https://127.0.0.1:2379 get a`{{execute}}  
`etcdctl --endpoints https://127.0.0.1:2379 ls`{{execute}}  
`etcdctl --endpoints https://127.0.0.1:2379 get b`{{execute}}  
`etcdctl --endpoints https://127.0.0.1:2379 set a 1`{{execute}}  
`etcdctl --endpoints https://127.0.0.1:2379 get b`{{execute}}  
