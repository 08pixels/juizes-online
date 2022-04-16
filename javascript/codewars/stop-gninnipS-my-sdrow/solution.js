function spinWords(string) {
    const reverseString = (str) => {
        return str
                .split('')
                .reverse()
                .join('')
    }

    return string
        .split(' ')
        .map(word => word.length < 5 ? word : reverseString(word))
        .join(' ')
}

console.log(spinWords("Hey fellow warriors"))