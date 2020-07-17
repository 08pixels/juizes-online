const input = require('fs').readFileSync('/dev/stdin', 'utf-8');
const numbers = input.split(/\s/);
const values = numbers.map( item => Number(item) );

function triangleArea(base, height) {
  return  (base * height) / 2;
}

function circleArea(radius) {
  return 3.14159 * Math.pow(radius, 2);
}

function trapezeArea(baseMinor, baseMajor, height) {
  return ((baseMajor + baseMinor) * height) / 2
}

function retangleArea(sideA, sideB) {
 return sideA * sideB; 
}

const A = Number( values.shift() );
const B = Number( values.shift() );
const C = Number( values.shift() );

console.log(`TRIANGULO: ${ triangleArea(A, C).toFixed(3) }`);
console.log(`CIRCULO: ${ circleArea(C).toFixed(3) }`);
console.log(`TRAPEZIO: ${ trapezeArea(A, B, C).toFixed(3) }`);
console.log(`QUADRADO: ${ retangleArea(B, B).toFixed(3) }`);
console.log(`RETANGULO: ${ retangleArea(A, B).toFixed(3) }`);
