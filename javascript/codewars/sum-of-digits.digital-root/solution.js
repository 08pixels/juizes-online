function digital_root(n) {
    if (n === 0)
        return 0

    let value = (n % 10) + digital_root(parseInt(n / 10))

    if (value >= 10)
        value = digital_root(value)

    return value
}