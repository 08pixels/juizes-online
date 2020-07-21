#include <bits/stdc++.h>

using namespace std;

int main() {
  double radius;
  cin >> radius;

  cout << setprecision(3) << fixed << "VOLUME = " << (4/3.0) * 3.14159 * pow(radius, 3) << endl;
  return 0;
}