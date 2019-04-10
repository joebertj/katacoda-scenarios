Here we will test some algorithms and determine which of them can accurately model the data.

First we will split the columns of our dataset into 2. `X` will contain the first 4 columns or the attributes. `Y` will contain the last column or the class. 
```
array = dataset.values
X = array[:,0:4]
Y = array[:,4]
```{{execute}}

We will only use 80% of the data to train and `validation_size` will be 20% of the remaining data.

`validation_size = 0.20`{{execute}}

Next we set the `seed` and and the `scoring` criteria. Seed just randomizes generated data by introducing some value as a starting point. We will use accuracy as our basis for selected the best model.

```
seed = 7
scoring = 'accuracy'
```{{execute}}

We are now going to use the above values to build the train test.

`X_train, X_validation, Y_train, Y_validation = model_selection.train_test_split(X, Y, test_size=validation_size, random_state=seed)`{{execute}}

The 6 different algorithms that we are going to test are:
- Logistic Regression (LR)
- Linear Discriminant Analysis (LDA)
- K-Nearest Neighbors (KNN)
- Classification and Regression Trees (CART)
- Gaussian Naive Bayes (NB)
- Support Vector Machines (SVM)
