Initialize kubernetes cluster using kubeadm
`kubeadm init`{{execute}}

Copy KUBECONFIG 
`cp /etc/kubernetes/admin.conf $HOME/`{{execute}}

List all pods in kube-system namespace
`kubectl get po -n kube-system`{{execute}}
