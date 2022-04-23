function pathFinder(maze) {

    const grid = maze.split('\n').map(line => line.split(''))

    const isValid = (coord) => {
        return coord >= 0 && coord < grid.length
    }

    const queue = [{ x: 0, y: 0, weight: 0 }]
    const dispX = [0, 0, 1, -1]
    const dispY = [1, -1, 0, 0]


    while (queue.length) {
        const { x, y, weight } = queue.shift()

        if (x === (grid.length - 1) && y === (grid.length - 1))
            return weight

        for (let i = 0; i < 4; ++i) {
            const u = x + dispX[i]
            const v = y + dispY[i]

            if (isValid(u) && isValid(v) && grid[u][v] === '.') {
                grid[u][v] = '#'
                queue.push({
                    'x': u,
                    'y': v,
                    'weight': 1 + weight
                })
            }
        }
    }

    return false;
}