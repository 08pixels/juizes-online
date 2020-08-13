class Solution {
public:
    vector<int> getRow(int rowIndex) {
        
        vector<int> memo[2];
        
        memo[0].resize(rowIndex + 1);
        memo[1].resize(rowIndex + 1);

        memo[0][0] = memo[1][0] = 1;
        
        for(int i = 1; i <= rowIndex; ++i) {
            for(int j = 1; j <= i; ++j)
                memo[i&1][j] = memo[~i&1][j] + memo[~i&1][j - 1];
        }
        
        return memo[rowIndex&1];
    }
};
