const roman = {
    1000: 'M',
    900: 'CM',
    500: 'D',
    400: 'CD',
    100: 'C',
    90: 'XC',
    50: 'L',
    40: 'XL',
    10: 'X',
    9: 'IX',
    5: 'V',
    4: 'IV',
    1: 'I',
}

function solution(number) {
    const keys = [1000,900,500,400,100,90,50,40,10,9,5,4,1]

    return keys.reduce((encoded, curr) => {
        encoded += new Array(parseInt(number / curr))
            .fill(roman[curr])
            .join('')

        number %= curr
        return encoded 
    }, '')
}

module.exports = solution