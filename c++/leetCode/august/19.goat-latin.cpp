class Solution {
public:
    string suffix = "ma";
    string vowel = "aeiou";
    
    string newWord(string word) {
        suffix += "a";

        if ( vowel.find( tolower(word[0]) ) != -1 )
            return word + suffix;
        
        return word.substr(1) + word[0] + suffix;
    }
    
    string toGoatLatin(string S) {
        int n = S.size();
        
        string answer = "";
        int last = 0;
        int k    = 0;
        
        while ( k != n ) {
            k = S.find(' ', last);
            
            if ( k == -1 )
                k = n;
            
            answer += (last ? " " : "") + newWord(S.substr(last, k - last));
                
            last = k + 1;
        }
        
        return answer;
    }
};
