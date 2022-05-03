function validSolution(board) {

    for (let i = 0; i < 9; ++i) {
        const row = new Array(10).fill(false)
        const col = new Array(10).fill(false)
        const square = new Array(10).fill(false)

        for (let j = 0; j < 9; ++j) {
            const idSquare = (i * 3) % 9

            const gapI = 3 * parseInt(i / 3) + parseInt(j / 3)
            const gapJ = idSquare + (j % 3)

            if (row[board[i][j]] || !board[i][j])
                return false
            if (col[board[j][i]])
                return false
            if (square[board[gapI][gapJ]])
                return false

            row[board[i][j]] = true
            col[board[j][i]] = true
            square[board[gapI][gapJ]] = true
        }
    }

    return true
}

module.exports = validSolution