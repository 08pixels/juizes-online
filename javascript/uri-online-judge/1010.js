const fs = require('fs').readFileSync('/dev/stdin', 'utf-8');
const parts = fs.split(/\n/);

const partA = parts[0].split(' ');
const partB = parts[1].split(' ');

const pieceA = {
  id    : partA[0],
  amount: Number( partA[1] ),
  price : Number( partA[2] )
};

const pieceB = {
  id    : partB[0],
  amount: Number( partB[1] ),
  price : Number( partB[2] )
};

const totalPriceA = pieceA.amount * pieceA.price;
const totalPriceB = pieceB.amount * pieceB.price;

const totalPrice  = totalPriceA + totalPriceB;

console.log(`VALOR A PAGAR: R$ ${ totalPrice.toFixed( 2 ) }`);
