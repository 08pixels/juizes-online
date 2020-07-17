const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const lines = fs.split(/\n/);

const info = lines.map( item => Number( item ));

const number      = info.shift();
const workedHours = info.shift();
const hourlyWage  = info.shift();

const salary = workedHours * hourlyWage;

console.log(`NUMBER = ${ number }`);
console.log(`SALARY = U$ ${ salary.toFixed( 2 ) }`);