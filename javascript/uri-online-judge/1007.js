const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const numbers = fs.split(/\n/);

const values = numbers.map( item => Number( item ) );

const productA = values[0] * values[1];
const productB = values[2] * values[3];

const difference = productA - productB;

console.log(`DIFERENCA = ${ difference }`);