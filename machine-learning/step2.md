The IRIS data set is a popular training data for Machine Learning. 

Once again open the Python 3 console:

`python3`{{execute}}

## Load data set

`url = "https://raw.githubusercontent.com/jbrownlee/Datasets/master/iris.csv"`{{execute}}

`names = ['sepal-length', 'sepal-width', 'petal-length', 'petal-width', 'class']`{{execute}}

`dataset = pandas.read_csv(url, names=names)`{{execute}}

## Describe the shape of the data set

`print(dataset.shape)`{{execute}}

## Describe the data set

`print(dataset.describe())`{{execute}}

# Show the class distribution of the data set

`print(dataset.groupby('class').size())`{{execute}}

