function expandedForm(num) {
    return num
            .toString()
            .split('')
            .map((item, pos, arr) => item * Math.pow(10, arr.length - pos - 1))
            .filter((item) => item)
            .join(' + ')
}