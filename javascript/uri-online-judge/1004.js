const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const lines = fs.split(/\n/);

const numberA = ( lines[0] >> 0 );
const numberB = ( lines[1] >> 0 );

console.log( `PROD = ${ numberA * numberB }` );