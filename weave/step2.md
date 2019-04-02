To visualize we can install Weave Scope:

`kubectl apply -f "https://cloud.weave.works/k8s/scope.yaml?k8s-version=$(kubectl version | base64 | tr -d '\n')"`{{execute}}

Wait until all Scope Pods are Running:

`watch 'kubectl get po -n weave | grep weave'`{{execute}}

If weave Pods are all Running then hit <kbd>CTRL-C</kbd>.

We will now expose the Pods on local network (extenal to Kubernetes):

`pod=$(kubectl get pod -n weave --selector=name=weave-scope-app -o jsonpath={.items..metadata.name})`{{execute}}

`ipaddr=$(ip -4 addr show ens3 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')`{{execute}}

`kubectl expose pod $pod -n weave --external-ip="$ipaddr" --port=4040 --target-port=4040`{{execute}}

You can access the dashboard here:

http://[[HOST_SUBDOMAIN]]-4040-[[KATACODA_HOST]].environments.katacoda.com/
