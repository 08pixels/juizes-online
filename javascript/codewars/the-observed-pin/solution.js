function getPINs(observed) {
    const similar = [
        [0, 8],
        [1, 2, 4],
        [2, 1, 5, 3],
        [3, 2, 6],
        [4, 1, 7, 5],
        [5, 2, 4, 8, 6],
        [6, 3, 5, 9],
        [7, 4, 8],
        [8, 5, 7, 0, 9],
        [9, 6, 8],
    ]

    const solve = (guess, ind) => {

        if (ind == guess.length)
            return [guess]

        const partial = []

        for (const choice of similar[guess[ind]]) {
            const left  = guess.slice(0, ind)
            const right =  guess.slice(1 + ind)

            const replaced = left + choice + right
            const result = solve(replaced, 1 + ind)

            partial.push(...result)
        }

        return partial
    }

    return solve(observed, 0)
}