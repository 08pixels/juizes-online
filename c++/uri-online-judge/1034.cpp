#include <bits/stdc++.h>

using namespace std;

int dp[1000001];

int main(void) {

    int test;
    cin >> test;

    while (test--) {

        memset(dp, 1e6, sizeof(dp));
        dp[0] = 0;

        int number, target;
        cin >> number >> target;

        int weigth[number];


        for (int i=0; i<number; ++i)
            cin >> weigth[i];

        for (int k=1; k<=target; ++k) {
            for (int i=0; i<number; ++i) {
                if (weigth[i] <= k) {
                    dp[k] = min(dp[k], 1 + dp[k - weigth[i]]);
                }
            }
        }

        cout << dp[target] << endl;
    }
    return 0;
}