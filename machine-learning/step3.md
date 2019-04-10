Here we will see generate some basic plots. 

`import matplotlib.pyplot as plt`{{execute}}

`dataset.plot(kind='box', subplots=True, layout=(2,2), sharex=False, sharey=False)`{{execute}}

`plt.savefig('box.png')`{{execute}}

[http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/box.png](Box plot)

`dataset.hist()`{{execute}}

`plt.savefig('hist.png')`{{execute}}

[http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/hist.png](Histogram)

## scatter plot matrix

`scatter_matrix(dataset)`{{execute}}

`plt.savefig('scatter.png')`{{execute}}

[http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/scatter.png](Scatter plot)

Exit the console using `quit()`{{execute}}
