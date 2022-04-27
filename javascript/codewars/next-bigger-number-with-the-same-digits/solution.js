function nextBigger(n) {
    const digits = n.toString().split('')

    for (let i = digits.length - 1; i > 0; --i) {
        if (digits[i] > digits[i - 1]) {

            let biggerClosest = i

            for (let k = digits.length - 1; k > biggerClosest; --k) {
                if (digits[k] > digits[i - 1]) {
                    biggerClosest = k
                    break
                }
            }

            [digits[biggerClosest], digits[i - 1]] = [digits[i - 1], digits[biggerClosest]]


            for (let k = i; k < digits.length; ++k) {
                for (let z = k + 1; z < digits.length; ++z)
                    if (digits[k] > digits[z]) {
                        [digits[k], digits[z]] = [digits[z], digits[k]]
                    }
            }

            return Number(digits.join(''))
        }
    }

    return -1
}

module.exports = nextBigger