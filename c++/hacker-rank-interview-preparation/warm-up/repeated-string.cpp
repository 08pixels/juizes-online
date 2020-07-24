#include <bits/stdc++.h>
using namespace std;

long repeatedString(string s, long n) {
  long amount = 0;
  long remaind = 0;

  for(int i=0; s[i]; ++i)
    if(s[i] == 'a')
      ++amount;

  for(int i=0; i<(int)(n % s.size()); ++i)
    if(s[i] == 'a')
      ++remaind;

  long result = n / s.size();

  return result * amount + remaind;
}

int main() {
  string str;
  long n;

  cin >> str >> n;
  cout << repeatedString(str, n) << endl;
  return 0;
}
