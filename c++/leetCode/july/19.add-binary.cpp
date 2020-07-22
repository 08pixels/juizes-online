class Solution {
public:
    string addBinary(string a, string b) {
        int maxlen = max(a.size(), b.size());
        
        while(a.size() != maxlen + 1)
            a = "0" + a;
        while(b.size() != maxlen + 1)
            b = "0" + b;
        
        string answer = string(maxlen + 1, '0');
        int bitCarry = 0;
        
        for(int i=maxlen; i >= 0; --i) {
            int bitA  = a[i] - '0';
            int bitB  = b[i] - '0';
            
            int sum = bitA + bitB + bitCarry;
            
            answer[i] = (sum&1) ? '1' : '0';
            bitCarry = sum >> 1;
        }
        
        if(answer[0] == '0')
            answer = string(answer.begin() + 1, answer.end());
        
        return answer;
    }
};
