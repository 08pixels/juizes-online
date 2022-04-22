function solveExpression(exp) {

    const fnNumber = (digit) => {
        return (operand) => {
            while(operand.indexOf('?') !== -1)
                operand = operand.replace('?', digit)

            return operand
        }
    }

    const makeOperation = (operator, lOperand, rOperand) => {
        if (operator === '+')
            return  Number(lOperand) + Number(rOperand)
        if (operator === '-')
            return  Number(lOperand) - Number(rOperand)

        return  Number(lOperand) * Number(rOperand)
    }

    const evaluateExpr = (operator, operands, result) => {
        return makeOperation(operator, ...operands).toString() === result
    }


    const [expr, result] = exp.split('=')

    const operators  = [
        expr.indexOf('+', 1),
        expr.indexOf('-', 1),
        expr.indexOf('*', 1),
    ]

    const operatorIndex = operators
                        .filter(index => index != -1)
                        .reduce((acc, curr) => acc < curr ? acc : curr)

    const operator     = expr[operatorIndex]
    const leftOperand  = expr.substr(0, operatorIndex)
    const rightOperand = expr.substr(1 + operatorIndex, expr.length)

    for (let digit=0; digit<=9; ++digit) {

        if (leftOperand.indexOf(digit) !== -1) continue;
        if (rightOperand.indexOf(digit) !== -1) continue;
        if (result.indexOf(digit) !== -1) continue;

        const getNumber = fnNumber(digit)

        const operands = [getNumber(leftOperand), getNumber(rightOperand)]
        const resultOperands   = getNumber(result)

        if (evaluateExpr(operator, operands, resultOperands))
            return digit
    }

    return -1
}

module.exports = solveExpression