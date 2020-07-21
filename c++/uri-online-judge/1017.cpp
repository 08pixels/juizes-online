#include <bits/stdc++.h>
using namespace std;

int main() {
  int hours, velocity;
  cin >> hours >> velocity;

  cout << setprecision(3) << fixed << ((hours * velocity) / 12.0) << endl;
  return 0;
}