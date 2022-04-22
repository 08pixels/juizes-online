const solveExpression = require('./solution')
const assert = require('assert')

const caseTests = [
    ['123*45?=5?088', 6],
    ['1+1=?', 2],
    ['-5?*-1=5?', 0],
    ['19--45=5?', -1],
    ['??*??=302?', 5],
    ['?*11=??', 2],
    ['??*1=??', 2],
    ['??+??=??', -1]
]

caseTests.forEach(([exp, expected]) => {
    assert.equal(solveExpression(exp), expected)
})