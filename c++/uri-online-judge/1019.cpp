#include <bits/stdc++.h>
using namespace std;

int main() {
  int timeInSeconds; cin >> timeInSeconds;

  int hours   =  timeInSeconds / 3600;
  timeInSeconds %= 3600;
  int minutes = timeInSeconds / 60;
  int seconds = timeInSeconds % 60;

  cout << hours << ":" << minutes << ":" << seconds << endl;
  return 0;
}