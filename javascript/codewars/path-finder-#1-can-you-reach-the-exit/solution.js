function pathFinder(maze) {
    const grid = maze.split('\n').map(item => item.split(''))

    const dispX = [1, -1, 0, 0]
    const dispY = [0, 0, 1, -1]
    const length = grid.length

    const isValid = (coord) => coord >= 0 && coord < length

    const queue = [[0, 0]]
    grid[0][0] = '#'

    while (queue.length) {
        const [x, y] = queue.shift()

        for (let i = 0; i < 4; ++i) {
            const u = x + dispX[i]
            const v = y + dispY[i]

            if (isValid(u) && grid[u][v] === '.') {
                if (u == length - 1 && v == length - 1)
                    return true
                queue.push([u, v])
                grid[u][v] = '#'
            }
        }
    }

    return false
}