// https://www.facebook.com/codingcompetitions/hacker-cup/2021/round-1/problems/A1

#include <bits/stdc++.h>
using namespace std;

int f(string str) {
    int change = 0;
    int l  = 0;

    for (auto &e: str) {
        if (e == 'F') continue;

        if (l && e != l)
            ++change;
        l = e;
    }

    return change;
}

int main() {
    int T; cin >> T;

    for (int test=0; test<T; ++test) {
        int N; cin >> N;
        string W; cin >> W;

        cout  << "Case #" << 1+test << ": " << f(W) << endl;
    }
    return 0;
}