const input = require('fs').readFileSync('/dev/stdin', 'utf8');
const [testCases, ...tests] = input.split('\n');

const isPrime = (number) => {

    if (number == 1)
        return false;

    for(let i=2; i*i <= number; ++i) {
        if (number % i === 0)
            return false;
    }

    return true;
}

for(const testNumber of tests) {

    const result = isPrime(testNumber);

    if (result)
        console.log('Prime')
    else
        console.log('Not Prime')
}