String.prototype.toJadenCase = function () {
    const isCapitalize = (word) => {
        const character = word.charAt(0)
        return character >= 'a' && character <= 'z'
    }
    
    const capitalize = (word) => {
        if (!isCapitalize(word))
            return word

        const upperCharCode = word.charCodeAt(0) - 32

        return String.fromCharCode(upperCharCode) + word.slice(1)
    }
    
    const words = this.split(' ')
    const capitalized = words.map(capitalize)
    
    return capitalized.join(' ')
}
