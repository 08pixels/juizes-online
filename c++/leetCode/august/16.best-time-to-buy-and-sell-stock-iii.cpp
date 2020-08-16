class Solution {
public:
    int size;
    vector<int> memo[3][2];
    
    int solve(int op, int ind, int amount, vector<int>& prices) {
        if ( ind == size )
            return 0;
        if ( amount == 2 )
            return 0;
        
        int &dp = memo[op][amount][ind];
        if ( dp != -1 )
            return dp;
        
        
        int l = 0, r = 0;
        
        if ( op == 1 )
            l = solve(2, ind + 1, amount + 1, prices) + prices[ind];
        else if ( op == 2 )
            r = solve(1, ind + 1, amount, prices) - prices[ind];
        else
            l = solve(1, ind + 1, amount, prices) - prices[ind];
            
        return dp = max(max(l,r), solve(op, ind + 1, amount, prices));
    }
    
    int maxProfit(vector<int>& prices) {
        size = prices.size();
        
        for(int i=0; i<3; ++i)
            for(int j=0; j<2; ++j)
            memo[i][j].resize(size + 1, -1);
        
        return solve(0, 0, 0, prices);
    }
    
};
