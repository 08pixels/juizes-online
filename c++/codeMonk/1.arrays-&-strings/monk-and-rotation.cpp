#include<bits/stdc++.h>
using namespace std;

int main()  {
    int t; cin >> t;
    
    while(t--) {
        int n, k; cin >> n >> k;
        
        vector<int> A(n);
        for(auto &e: A) cin >> e;

        k %= n;
        for (int i=n-k; i<n; ++i) cout << A[i] << " ";
        for (int i=0; i<n-k; ++i) cout << A[i] << " ";
        cout << endl;
    }
    return 0;
}