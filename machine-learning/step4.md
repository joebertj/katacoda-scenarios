Here we will test some algorithms and determine which of them can accurately model the data.

We need to load the `sklearn` modules at this point. 

```
from sklearn import model_selection
from sklearn.metrics import classification_report
from sklearn.metrics import confusion_matrix
from sklearn.metrics import accuracy_score
from sklearn.linear_model import LogisticRegression
from sklearn.tree import DecisionTreeClassifier
from sklearn.neighbors import KNeighborsClassifier
from sklearn.discriminant_analysis import LinearDiscriminantAnalysis
from sklearn.naive_bayes import GaussianNB
from sklearn.svm import SVC
```{{execute}}

First, we will split the columns of our dataset into 2. `X` will contain the first 4 columns or the attributes. `Y` will contain the last column or the class. We transform the dataset into an Array and apply Slicing.

```
array = dataset.values
X = array[:,0:4]
Y = array[:,4]
```{{execute}}

We will only use 80% of the data to train and `validation_size` will be 20% of the remaining data.

`validation_size = 0.20`{{execute}}

Out of the data that we will use to train, we set aside 1 part as test data. We will use the `splits` variable for that. Here 10 means 9 rows will be used to train and 1 row to test for all the 80% of data we have set aside.

`splits = 10`{{execute}}

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

So, we build these models now using a List:


```
models = []
models.append(('LR', LogisticRegression(solver='liblinear', multi_class='ovr')))
models.append(('LDA', LinearDiscriminantAnalysis()))
models.append(('KNN', KNeighborsClassifier()))
models.append(('CART', DecisionTreeClassifier()))
models.append(('NB', GaussianNB()))
models.append(('SVM', SVC(gamma='auto')))
```{{execute}}

We now use loop into the models List 

```
msg = "%s: %f (%f)" % ("Algorithm", "Mean", "Standard Deviation")
print(msg)
results = []
names = []
for name, model in models:
	kfold = model_selection.KFold(n_splits=splits, random_state=seed)
	cv_results = model_selection.cross_val_score(model, X_train, Y_train, cv=kfold, scoring=scoring)
	results.append(cv_results)
	names.append(name)
	msg = "%s: %f (%f)" % (name, cv_results.mean(), cv_results.std())
	print(msg)

```{{execute}}

Now, we compare the algorithms:

```
fig = plt.figure()
fig.suptitle('Algorithm Comparison')
ax = fig.add_subplot(111)
plt.boxplot(results)
ax.set_xticklabels(names)
plt.savefig('compare.png')
```{{execute}}

[Algorithm comparison](http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/compare.png)
