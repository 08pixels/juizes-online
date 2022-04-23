function dblLinear(n) {

    const results = [1, 3, 4]

    for (let i = 1; i <= 5*n; ++i) {
        results.push(2 * results[i] + 1)
        results.push(3 * results[i] + 1)
    }

    return [... new Set(results)]
                .sort((a,b) => {
                    if (Number(a) < Number(b))
                        return -1
                    return 1
                })[n]
}