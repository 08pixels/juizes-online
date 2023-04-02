#include<bits/stdc++.h>
#define lli long long int
#define pll pair<lli, pair<lli, lli>>

using namespace std;


bool compare(const pll &p, const pll &q) {
    return p.first < q.first;
}

int main() {
    ios_base::sync_with_stdio(false); cin.tie(NULL); cout.tie(NULL); 

    int n; scanf("%d", &n);

    vector<pll> values(n);

    for(int i=0; i<n; ++i) {
        lli number; scanf("%lld", &number);
        // chunk, number, number-without-last-chunks
        values[i] = {0LL, {number, number}};
    }

    bool repeat = true;
    lli factor = 100000LL;

    while(repeat) {
        repeat = false;

        for(auto &curr: values) {
            lli chunck = curr.second.second % factor;
            curr.first = chunck;
            curr.second.second /= factor;

            if (curr.second.second) repeat = true;
        }

        stable_sort(values.begin(), values.end(), compare);

        for(auto &e:values)
            printf("%lld ", e.second.first);
        printf("\n");
    }

    return 0;
}