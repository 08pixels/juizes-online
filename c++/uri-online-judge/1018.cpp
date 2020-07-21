#include <bits/stdc++.h>
using namespace std;

int main() {
  int n;
  int value[] = {100,50,20,10,5,2,1};
  cin >> n;

  cout << n << endl;

  for(int i=0; i<7; ++i) {
    printf("%d nota(s) de R$ %d,00\n", n/value[i], value[i]);
    n %= value[i];
  }
  return 0;
}