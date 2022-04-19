var recoverSecret = function(triplets) {
    const ancestor = triplets.reduce((anc, triplet) => {
        const [x, y, z] = triplet

        anc[x] = (anc[x] || [])
        anc[y] = (anc[y] || []).concat([x])
        anc[z] = (anc[z] || []).concat([x, y])

        return anc
    }, {})

    let secretWord = ''
    let stringLength = Object.keys(ancestor).length

    const min = (subsetA, subsetB) => (
        subsetA[1].length < subsetB[1].length
            ? subsetA
            : subsetB
    )

    while (stringLength--) {
        const [char, value] = Object.entries(ancestor).reduce(min)

        for (const [key, value] of Object.entries(ancestor))
            ancestor[key] = value.filter(item => item != char)

        secretWord += char
        delete ancestor[char]
    }

    return secretWord
}

triplets1 = [
    ['t','u','p'],
    ['w','h','i'],
    ['t','s','u'],
    ['a','t','s'],
    ['h','a','p'],
    ['t','i','s'],
    ['w','h','s']
  ]

console.log(recoverSecret(triplets1))