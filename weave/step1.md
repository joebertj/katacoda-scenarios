Make sure Kubernetes is started:

`launch.sh`{{execute}}

First things first, check if nodes are ready:

`kubectl get nodes`{{execute}}

If all are in Ready state then you can proceed installing Weave:

`kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"`{{execute}}

`watch 'kubectl get po -n kube-system | grep weave'`{{execute}}

If weave Pods are all Running then hit <kbd>CTRL-C</kbd>.
