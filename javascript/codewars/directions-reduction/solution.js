function dirReduc(arr) {
    const map = {'NORTH': 1, 'SOUTH': -1, 'WEST':  2, 'EAST' : -2 }

    return arr.reduce((stack, direction) => {
        if (stack.length && map[direction] === -map[stack[stack.length - 1]])
            stack.pop()
        else
            stack.push(direction)

        return stack
    }, [])
}
