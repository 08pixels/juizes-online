class Solution {
public:
    vector<int> answer;
    
    void solve(int pos, int value, int N, int K) {

        if ( pos == N ) {
            answer.push_back(value);
            return;
        }
        
        for(int digit = 0; digit <= 9; ++digit) {
            if ( abs( (value % 10) - digit ) ==  K )
                solve(1 + pos, 10 * value + digit, N, K);
        }
    }
    
    vector<int> numsSameConsecDiff(int N, int K) {
        if ( N == 1 )
            answer.push_back(0);

        for(int i=1; i<=9; ++i)
            solve(1, i, N, K);

        return answer;
    }
};
