function sumStrings(a,b) {
    const numberA = a.split('').reverse().join('')
    const numberB = b.split('').reverse().join('')

    const size = numberA.length > numberB.length ? numberA.length : numberB.length

    let result = ''
    let carry = 0

    for (let i=0; i<size; ++i) {
        let curr = Number(numberA[i] || 0) + (Number(numberB[i]) || 0) + carry

        carry = parseInt(curr / 10)
        result = (curr % 10) + result
    }

    return (carry + result).replace(/^0+/, '')
}

module.exports = sumStrings

// BigInt works fine