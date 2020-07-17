const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const lines = fs.split(/\n/);

const numberA = ( lines.shift() >> 0 );
const numberB = ( lines.shift() >> 0 );

console.log( `SOMA = ${ numberA + numberB }` );