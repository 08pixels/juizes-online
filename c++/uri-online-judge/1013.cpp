#include <bits/stdc++.h>
using namespace std;

int main() {
  int a,b,c;
  cin >> a >> b >> c;
  cout << max(a, max(b, c) ) << " eh o maior" << endl;
  return 0;
}