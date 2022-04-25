function knight(start, finish) {

    const convertCoordinate = (coord) => {
        return [
            coord.charCodeAt(0) - 'a'.charCodeAt(0),
            coord.charCodeAt(1) - '1'.charCodeAt(0)
        ]
    }

    const isValid = (coord) => {
        return coord >= 0 && coord < 8
    }
    const [sourceX, sourceY] = convertCoordinate(start)
    const [targetX, targetY] = convertCoordinate(finish)

    const dispX = [2, 2, -2, -2, 1, -1, 1, -1]
    const dispY = [1, -1, 1, -1, 2, 2, -2, -2]

    const seen = { '00': true }
    const queue = [{ x: sourceX, y: sourceY, weight: 0 }]

    while (queue.length) {
        const { x, y, weight } = queue.shift()

        for (let i = 0; i < 8; ++i) {
            const u = x + dispX[i]
            const v = y + dispY[i]

            if (isValid(u) && isValid(v) && !seen['' + u + v]) {
                seen['' + u + v] = true
                queue.push({ x: u, y: v, weight: 1 + weight })
            }

            if (u === targetX && v === targetY)
                return 1 + weight
        }

    }


}