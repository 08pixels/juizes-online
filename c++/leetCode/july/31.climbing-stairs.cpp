/**
*   You are climbing a stair case.
*   It takes n steps to reach to the top.
*   Each time you can either climb 1 or 2 steps.
*   In how many distinct ways can you climb to the top?
**/

class Solution {
public:
    int climbStairs(int n) {
        
        int memo[n + 3];
        memo[1] = 1;
        memo[2] = 2;
        
        for( int i = 3; i <= n; ++i )
            memo[i] = memo[i - 1] + memo[i - 2];
        
        return memo[n];
    }
};
