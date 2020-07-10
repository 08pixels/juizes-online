let input = require('fs').readFileSync('/dev/stdin', 'utf-8');
let lines = input.split(/\s/);

let a = Number(lines.shift());
let b = Number(lines.shift());

console.log(`X = ${a+b}`);