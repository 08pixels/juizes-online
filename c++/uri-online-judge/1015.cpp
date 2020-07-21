#include <bits/stdc++.h>

using namespace std;

int main() {
  double a, b, c, d;
  cin >> a >> b >> c >> d;
  cout << setprecision(4) << fixed << hypot(c-a, d-b) << endl;
  return 0;
}