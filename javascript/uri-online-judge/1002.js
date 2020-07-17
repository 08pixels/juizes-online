
const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
let lines = fs.split(/\s/);

let radius = Number( lines.shift() );
let area   = 3.14159 * radius * radius;

console.log(`A=${ area.toFixed(4) }`);