const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const lines = fs.split(/\n/);

const numberA = Number( lines.shift() );
const numberB = Number( lines.shift() );

const weightA = 3.5
const weightB = 7.5

const ponderedMean = ( numberA * weightA  + numberB * weightB ) / (weightA + weightB);

console.log(`MEDIA = ${ ponderedMean.toFixed( 5 ) }`);