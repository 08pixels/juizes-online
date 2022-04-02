const input = require('fs').readFileSync('/dev/stdin', 'utf-8')
const [distance, liters] = input.split(/\s/).map(parseFloat)

const result = distance/liters; 

console.log(`${result.toFixed(3)} km/l`)