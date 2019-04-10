The IRIS data set is a popular training data for Machine Learning. It is just a collection of measurements obtained from samples of 3 Species of Iris. The values are Length and Width of Petals and Sepals. This is a (Wikipedia entry)[https://en.wikipedia.org/wiki/Iris_flower_data_set].

Once again open the Python 3 console:

`python3`{{execute}}

## Load data set

`import pandas`{{execute}}

`url = "https://raw.githubusercontent.com/jbrownlee/Datasets/master/iris.csv"`{{execute}}

`names = ['sepal-length', 'sepal-width', 'petal-length', 'petal-width', 'class']`{{execute}}

`dataset = pandas.read_csv(url, names=names)`{{execute}}

## Describe the shape of the data set

`print(dataset.shape)`{{execute}}

This means that we have 150 rows and 5 columns. The columns correspond to the List of names.

## Describe the data set

`print(dataset.describe())`{{execute}}

Here we can see the mean, standard deviation, percentiles, minimum and maximum values for each column.

# Show the class distribution of the data set

`print(dataset.groupby('class').size())`{{execute}}

Here we can see that 50 samples came from each of the 3 Species.

