#include<bits/stdc++.h>
using namespace std;

bool compare (pair<int, int> &p, pair<int, int> &q)  {
    return p.first < q.first;
}

int main() {
    int t; cin >> t;

    while(t--)  {
        int n; cin >> n;
        vector<int> v(n);

        for(auto &e: v) cin >> e;
        sort(v.begin(), v.end());

        vector<pair<int, int>> cache;
        
        cache.push_back({0, v[0]});
        int size = 0;

        for(int i=0; i<n; ++i) {
            if (cache[size].second == v[i]) {
                ++cache[size].first;
            } else {
                cache.push_back({1, v[i]});
                ++size;
            }
        }

        sort(cache.begin(), cache.end(), compare);

        if (size == 0) cout << -1 << endl;
        else cout << cache[size].first - cache[0].first << endl;
    }
}