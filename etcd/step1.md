Initialize kubernetes cluster using kubeadm
`kubeadm init`{{execute}}

Go to Node Terminal and execute the command on the last line of `kubeadm init`. It looks like:
```
kubeadm join 172.17.0.16:6443 --token f5wqng.s7yvanrkug1p9her \
    --discovery-token-ca-cert-hash sha256:e50d7e45b5d8a0819466caa5ee1f6576050ac9285537d17f569447be5670f72c
```

Set KUBECONFIG to authorize `kubectl`
`export KUBECONFIG=/etc/kubernetes/admin.conf`{{execute}}

List all pods in kube-system namespace
`kubectl get po -n kube-system`{{execute}}
