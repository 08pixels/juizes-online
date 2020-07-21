#include <bits/stdc++.h>

using namespace std;

int main() {
  string name;
  double salary, totalSold;

  cin >> name >> salary >> totalSold;
  cout << setprecision(2) << fixed << "TOTAL = R$ " << (salary + totalSold * 0.15) << endl;
  return 0;
}