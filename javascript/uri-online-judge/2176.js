const input = require('fs').readFileSync('/dev/stdin', 'utf-8');
const word = input.split(/\s/)[0];

let counter = 0;

for (const letter of word) {
    if (letter == '1')
        ++counter
}

console.log(word + (counter&1))