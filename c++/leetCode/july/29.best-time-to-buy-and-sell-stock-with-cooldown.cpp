/*
    Say you have an array for which the ith element is the price of a given stock on day i.
    Design an algorithm to find the maximum profit. You may complete as many transactions
    as you like (ie, buy one and sell one share of the stock multiple times)

     with the following restrictions:
    - You may not engage in multiple transactions at the same time
      (ie, you must sell the stock before you buy again).
    - After you sell your stock, you cannot buy stock on next day.
      (ie, cooldown 1 day)
*/

class Solution {
public:
    
    int maxProfit(vector<int>& prices) {
        if(prices.size() < 2) return 0;
        
        int N = prices.size();
        int memo[N][2];
        
        memo[0][0] = 0;
        memo[0][1] = -prices[0]; 
        
        memo[1][0] = max(memo[0][0], prices[1] + memo[0][1]);
        memo[1][1] = max(memo[0][1], memo[0][0] - prices[1]);
        
        for(int i = 2; i < N; ++i) {
            memo[i][0] = max(memo[i - 1][0], prices[i] + memo[i - 1][1]);
            memo[i][1] = max(memo[i - 1][1], memo[i - 2][0] - prices[i]);
        }
        
        return memo[N - 1][0];
    }
};
