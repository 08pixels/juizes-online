class Solution {
public:
    int titleToNumber(string s) {
        int n = s.size();
        int answer = 0;
        
        for(int i=0; i<n; ++i)
            answer += (s[i] - 'A' + 1) * pow(26, n - i - 1);
        
        return answer;
    }
};
