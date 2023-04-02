#include<bits/stdc++.h>
using namespace std;

int main() {
    int t; cin >> t;

    while(t--) {
        int n; cin >> n;
        
        vector<int> A[n];
        for(int i=0; i<n; ++i) {
            A[i].resize(n);
            for (int j=0; j<n; ++j){
                cin >> A[i][j];
            }
        }

        int answer = 0;

        for(int k=0; k<n; ++k) {
            for(int l=0; l<n; ++l) {
                for(int i=k; i<n; ++i) {
                    for(int j=l; j<n; ++j) {
                        if (A[k][l] > A[i][j]) ++answer;
                    }
                }
            }
        }

        cout << answer << endl;
    }

    return 0;
}