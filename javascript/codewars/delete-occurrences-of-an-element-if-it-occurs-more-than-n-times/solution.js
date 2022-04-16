function deleteNth(arr, n) {
    const cache = {}

    const avaliate = (number) => {
        cache[number] = 1 + (cache[number] || 0)

        return cache[number] <= n
    }


    return arr.filter(avaliate)
}
