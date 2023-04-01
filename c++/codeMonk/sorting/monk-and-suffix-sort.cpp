#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    int k;

    cin >> s >> k;
    int N = s.size();

    vector<string> suffix(N);

    for(int i=0; i<N; ++i)  {
        string currentSuffix = s.substr(i);
        suffix[i] = currentSuffix;
    }

    for(int i=0; i<k; ++i) {
        for(int j=i+1; j<N; ++j) {
            if (suffix[i] > suffix[j])
                swap(suffix[i], suffix[j]);
        }
    }

    cout << suffix[k-1] << endl;
}