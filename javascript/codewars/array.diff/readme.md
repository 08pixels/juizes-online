### [Array.diff](https://www.codewars.com/kata/523f5d21c841566fde000009/train/javascript)  


Your goal in this kata is to implement a difference funcrion, which subtracts one list from another and returns the result.

It should remove all values from list `a`, which are present in list `b` keeping their order.

```js
arrayDiff([1,2],[1]) == [2])
```

If a value is present in `b`, all of its occurences must be removed from the other:

```js
arrayDiff([1,2,2,2,3], [2]) == [1,3]
```
