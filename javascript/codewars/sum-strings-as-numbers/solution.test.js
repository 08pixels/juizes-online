const sumStrings = require('./solution')
const assert = require('assert')

assert.equal(sumStrings('123','456'),'579')
assert.equal(sumStrings('99','9999'),'10098')
assert.equal(sumStrings('00103', '08567'),'8670')