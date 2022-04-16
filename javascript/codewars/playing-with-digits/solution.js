function digPow(n, p){
    
    const getDigits = (number) => {
        const digits = []
        
        while (number) {
            digits.unshift(number%10)
            number = parseInt(number/10)
        }

        return digits
    }

    const result =  getDigits(n)
                        .map((item, ind) => Math.pow(item, p + ind))
                        .reduce((prev, curr) => curr + prev)

    return result % n ? -1 : result / n;
}