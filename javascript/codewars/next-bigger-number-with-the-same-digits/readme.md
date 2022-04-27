### [Next bigger number with the same digits](https://www.codewars.com/kata/55983863da40caa2c900004e)  

Create a function that takes a positive integer and returns the next bigger number that can be formed by rearranging its digits. For example:

```js
12 ==> 21
513 ==> 531
2017 ==> 2071
```

```js
nextBigger(num: 12)   // returns 21
nextBigger(num: 513)  // returns 531
nextBigger(num: 2017) // returns 2071
```

If the digits can't be rearranged to form a bigger number, return `-1` (or `nil` in Swift):

```js
9 ==> -1
111 ==> -1
531 ==> -1
```

```js
nextBigger(num: 9)   // returns nil
nextBigger(num: 111) // returns nil
nextBigger(num: 531) // returns nil
```

:paperclip: `REFACTORING` `NUMBERS` `STRINGS` `INTEGERS`