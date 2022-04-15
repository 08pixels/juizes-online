function solution(number){
    if (number < 0)
        return 0;

    const sumOfSequencies = (first) => {
        const amount = parseInt((number - 1) / first)
        const last   = amount * first

        return parseInt(((first + last) * amount) / 2)
    }

    const sumOf3 = sumOfSequencies(3)
    const sumOf5 = sumOfSequencies(5)
    const intersection = sumOfSequencies(5 * 3)

    return (sumOf3 + sumOf5) - intersection 
}