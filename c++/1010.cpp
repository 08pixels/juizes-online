#include <bits/stdc++.h>

using namespace std;

typedef struct {
  int id, amount;
  double price;
} data;

int main() {
  data piece;

  double answer = 0;
  for(int i=0; i<2; ++i) {
    cin >> piece.id >> piece.amount >> piece.price;
    answer += (piece.amount * piece.price);
  }
  
  cout << setprecision(2) << fixed << "VALOR A PAGAR: R$ " << answer << endl;
  return 0;
}