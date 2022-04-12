const input = require('fs').readFileSync('/dev/stdin', 'utf8');
const [testCases, ...tests] = input.split('\n');

const isPrime = (number) => {

    for(let i=2; i*i <= number; ++i) {
        if (number % i === 0)
            return false;
    }

    return true;
}

for(let i=0; i<testCases; ++i) {

    const result = isPrime(Number(tests[i]));

    if (result)
        console.log('Prime')
    else
        console.log('Not Prime')
}