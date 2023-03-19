// https://atcoder.jp/contests/abc211/tasks/abc211_d

#include <bits/stdc++.h>
using namespace std;
const int INF = 1<<20; 
const long long MOD = 1000000007;

int main() {
    int n,m;
    cin >> n >> m;

    vector<vector<int>> graph(n+1);

    for(int i=0; i<m; ++i) {
        int u, v;
        cin >> u >> v;

        graph[u].push_back(v);
        graph[v].push_back(u);
    }

    vector<int> dist(n + 1, INF);
    vector<int> dp(n + 1, 0);

    queue<int> bfs;
    bfs.push(1);
    dist[1] = 0;
    dp[1] = 1;

    while(!bfs.empty()) {
        int size = bfs.size();

        while (size--){
            int u = bfs.front();
            bfs.pop();

            for(auto &v: graph[u]) {
                if (dist[v] == INF) {
                    dist[v] = dist[u] + 1;
                    bfs.push(v);

                    dp[v] = dp[u];
                }
                else if ((dist[u] + 1) == dist[v]) {
                    dp[v] = (dp[v] + dp[u]) % MOD;
                }
            }
        }
    }

    cout << dp[n] << endl; 

    return 0;
}
