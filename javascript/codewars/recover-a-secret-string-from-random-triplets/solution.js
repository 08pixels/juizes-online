var recoverSecret = function(triplets) {
    const before = {}

    for (const [x,y,z] of triplets) {
        if (!before[x]) before[x] = []
        if (!before[y]) before[y] = []
        if (!before[z]) before[z] = []

        before[z].push(x, y)
        before[y].push(x)
    }

    let secretWord = ''
    let stringLength = Object.keys(before).length

    while (stringLength--) {
        let char = ''
        let smaller = 999

        for (const [key, value] of Object.entries(before)) {
            if (value.length < smaller)
                smaller = value.length, char = key
        }

        for (const [key, value] of Object.entries(before))
            before[key] = value.filter(item => item != char)

        secretWord += char
        delete before[char]
    }

    return secretWord
}
