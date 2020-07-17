const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const input = fs.split(/\n/);

const sellersName = input.shift();
const salary      = Number( input.shift() );
const salesMade   = Number( input.shift() );

const percentage  = 0.15;
const commision   = salesMade * percentage;

const salaryTotal = salary + commision;

console.log(`TOTAL = R$ ${ salaryTotal.toFixed( 2 ) }`);