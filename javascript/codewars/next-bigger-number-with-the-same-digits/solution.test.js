const nextBigger = require('./solution')
const assert = require('assert')

assert.equal(nextBigger(12), 21)
assert.equal(nextBigger(513), 531)
assert.equal(nextBigger(2017), 2071)
assert.equal(nextBigger(414), 441)
assert.equal(nextBigger(144), 414)
assert.equal(nextBigger(1), -1)
assert.equal(nextBigger(111), -1)
assert.equal(nextBigger(4639426398), 4639426839)
assert.equal(nextBigger(859339861), 859361389)