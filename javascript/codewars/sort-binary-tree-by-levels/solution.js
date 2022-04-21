function treeByLevels(rootNode) {
    const queue = [rootNode]
    const result = []

    while (rootNode && queue.length) {
        let levelLength = queue.length

        while (levelLength--) {
            const node = queue.shift()

            result.push(node.value)
            if (node.left) queue.push(node.left)
            if (node.right) queue.push(node.right)
        }
    }

    return result;
}