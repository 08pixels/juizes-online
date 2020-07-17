const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const numbers = fs.split('\n');
const weights = [2, 3, 5];

if ( numbers.length > 3) {
  numbers.pop();
}

const ponderedNotes    = numbers.map(( item, index ) => Number( item ) * weights[index] );
const sumPonderedNotes = ponderedNotes.reduce(( accumulator, item ) => accumulator + item );
const sumOfWeights     = weights.reduce(( accumulator, item ) => accumulator + item );
const ponderedMean     = sumPonderedNotes/sumOfWeights;

console.log(`MEDIA = ${ ponderedMean.toFixed(1) }`);
