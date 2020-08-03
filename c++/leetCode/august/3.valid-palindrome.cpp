/**
*   Given a string, determine if it is a palindrome, considering
*   only alphanumeric characters and ignoring cases.
*
*   Note: For the purpose of this problem,
*       we define empty string as valid palindrome.
*/

class Solution {
public:
    
    char toUpper(char c) {
        if(c >= 'a' && c <= 'z')
            return c - 32;
        
        return c;
    }
    
    bool isValid(char c) {
        if(c >= 'A' && c <= 'Z')
            return true;
        if(c >= '0' && c <= '9')
            return true;

        return false;
    }
    
    bool isPalindrome(string s) {
    
        int i = 0;
        int j = s.size() - 1;
        
        while(i < j) {
            s[i] = toUpper(s[i]);
            s[j] = toUpper(s[j]);
            
            if(! isValid(s[i]) ) ++i;
            else if(! isValid(s[j]) ) --j;
            else if( s[i] == s[j] ) ++i, --j;
            else return false;
        }
        
        return true;
    }
};
