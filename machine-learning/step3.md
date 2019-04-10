Here we will see generate some basic plots.

`import matplotlib.pyplot as plt`{{execute}}

## box and whisker plots

`dataset.plot(kind='box', subplots=True, layout=(2,2), sharex=False, sharey=False)`{{execute}}

`plt.savefig('box.png')`{{execute}}

[Box plot](('box.png')

## histograms

`dataset.hist()`{{execute}}

`plt.savefig('hist.png')`{{execute}}

## scatter plot matrix

`scatter_matrix(dataset)`{{execute}}

`plt.savefig('scatter.png')`{{execute}}

Exit the console using `quit()`{{execute}}
