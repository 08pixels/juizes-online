/*
    the straight solution:
        n*n*n
*/

function rowSumOddNumbers(n) {
    const getFirstElement = () => {
        const firstElement = 1
        const ratio        = (2 * n)
        const nth          = (n - 1)

        return firstElement + ((ratio * nth) / 2)
    }

    const sumElements = () => {
        const firstElement = getFirstElement()
        const lastElement  = firstElement + (2 * (n - 1)) 

        return ((firstElement + lastElement) * n) / 2 
    }

    return sumElements()
}
