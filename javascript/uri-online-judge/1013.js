const input = require('fs').readFileSync('/dev/stdin', 'utf-8');
const values = input.split(/\s/);

const numbers = values.map( item => Number( item ));

const max = (a, b) => {
  return ( a + b + Math.abs( a-b ) ) / 2;
}

let largest = max(numbers[0], numbers[1]);
largest = max(largest, numbers[2]);

console.log(`${ largest } eh o maior`);
