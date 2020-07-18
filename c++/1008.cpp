#include <bits/stdc++.h>

using namespace std;

int main() {
  int id, hours;
  double hourValue;

  cin >> id >> hours >> hourValue;
  cout << "NUMBER = " << id << endl;
  cout << setprecision(2) << fixed << "SALARY = U$ " << (hours * hourValue) << endl;
  return 0;
}