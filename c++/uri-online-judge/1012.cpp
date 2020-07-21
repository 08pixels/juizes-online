#include <bits/stdc++.h>

using namespace std;

int main() {
  double a,b,c;
  cin >> a >> b >> c;

  cout << setprecision(3) << fixed << "TRIANGULO: " << ((a * c) / 2) << endl;
  cout << setprecision(3) << fixed << "CIRCULO: "   << (3.14159 * c * c) << endl;
  cout << setprecision(3) << fixed << "TRAPEZIO: "  << (((a + b) * c) / 2) << endl;
  cout << setprecision(3) << fixed << "QUADRADO: "  << (b * b) << endl;
  cout << setprecision(3) << fixed << "RETANGULO: " << (a * b) << endl;
}
