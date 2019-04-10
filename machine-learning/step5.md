Here we will pick an algorithm and use it to predict by testing it against the validation data. In the result that we saw earlier, KNN has a pretty good accuracy. We use that algorithm now to predict.

```
knn = KNeighborsClassifier()
knn.fit(X_train, Y_train)
predictions = knn.predict(X_validation)
print(accuracy_score(Y_validation, predictions))
print(confusion_matrix(Y_validation, predictions))
print(classification_report(Y_validation, predictions))
```{{execute}}

Exit the console using `quit()`{{execute}}
