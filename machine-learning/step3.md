Here we will see generate some basic plots. 

`import matplotlib.pyplot as plt`{{execute}}

`from pandas.plotting import scatter_matrix`{{execute}}

`dataset.plot(kind='box', subplots=True, layout=(2,2), sharex=False, sharey=False)`{{execute}}

`plt.savefig('box.png')`{{execute}}

[Box plot](http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/box.png)

`dataset.hist()`{{execute}}

`plt.savefig('hist.png')`{{execute}}

[Histogram](http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/hist.png)

## scatter plot matrix

`scatter_matrix(dataset)`{{execute}}

`plt.savefig('scatter.png')`{{execute}}

[Scatter plot](http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/scatter.png)

Exit the console using `quit()`{{execute}}
