class Solution {
public:
    
    int longestPalindrome(string s) {
        long long int mask = 0LL, bit  = 0LL;
        int op, carry, answer = 0x00;
        int n = s.size();
        
        for (int i = 0; i < n; ++i) {
            
            if(s[i] <= 'Z')
                op = 'A', carry = 0x00;
            else
                op = 'a', carry = 0x1A;
            
            bit = (1LL << (s[i] - 'A'));

            if(mask & bit) {
                mask ^= bit;
                answer += 0x02;
            } else {
                mask |= bit;
            }
        }
        
        return answer + (mask > 0);
    }
};
