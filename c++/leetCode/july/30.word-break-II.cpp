class Solution {
public:
    vector<string> answer;
    map<int, vector<string>> memo;
    
    vector<string> solve(string s, int end, vector<string>& wordDict) {
        vector<string> curr;
        
        if( end == 0 ) {
            curr.push_back("");
            return curr;
        }
        
        if( memo.count(end) )
            return memo[end];
        
        for(int i = 0; i < end; ++i) {
            
            string guest = s.substr(i, end - i);
            
            if(find(wordDict.begin(), wordDict.end(), guest) != wordDict.end())  {
                vector<string> tmp = solve(s, i, wordDict);
                
                for(auto str: tmp) {
                    curr.push_back(str.size() == 0 ? guest : str + " " + guest);
                }
            }
        }

        return memo[end] = curr;
    }
    
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        return solve(s, s.size(), wordDict);
    }
};
