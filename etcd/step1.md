Initialize kubernetes cluster using kubeadm
`kubeadm init`

Copy KUBECONFIG 
`cp /etc/kubernetes/admin.conf $HOME/`

List all pods in kube-system namespace
`kubectl get po`
