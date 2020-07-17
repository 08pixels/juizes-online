const input = require('fs').readFileSync('/dev/stdin', 'utf-8');
const number = input.split('\n');

const sphereVolume = (radius) => {
  return (4/3) * 3.14159 * Math.pow(radius, 3);
}

const radius = number.shift();
const volume = sphereVolume( radius );

console.log(`VOLUME = ${ volume.toFixed(3) }`);
