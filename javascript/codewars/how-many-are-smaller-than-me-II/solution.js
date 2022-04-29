function smaller(arr) {
    const smallest = Math.abs(arr.reduce((item, acc) => item < acc ? item : acc))
    const biggest = arr.reduce((item, acc) => item > acc ? item : acc)

    const bigger = 1 + (smallest > biggest ? smallest : biggest)
    const MAX = 2 * bigger
    const bit = new Array(MAX + 10).fill(0)

    const update = (x, counter = 1) => {
        while (x <= MAX) {
            bit[x] += counter
            x += (x & -x)
        }
    }

    const query = (x) => {
        let result = 0

        while (x > 0) {
            result += bit[x]
            x -= (x & -x)
        }

        return result
    }

    for (let i = 0; i < arr.length; ++i)
        update(bigger + arr[i])

    return arr.map(item => {
        update(bigger + item, -1)
        return query(bigger + item - 1)
    })
}

module.exports = smaller