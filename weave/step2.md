To visualize we can install Weave Scope:

`kubectl apply -f "https://cloud.weave.works/k8s/scope.yaml?k8s-version=$(kubectl version | base64 | tr -d '\n')"`{{execute}}

Wait until all Scope Pods are Running:

`watch 'kubectl get po -n weave | grep weave'`{{execute}}

You can access the dashboard here:

http://[[HOST_SUBDOMAIN]]-4040-[[KATACODA_HOST]].environments.katacoda.com/centreon
