/**
 * Given a word, you need to judge whether the usage of capitals
 * in it is right or not.
 * We define the usage of capitals in a word to be right
 * when one of the following cases holds:
 * All letters in this word are capitals, like "USA".
 *  All letters in this word are not capitals, like "leetcode".
 *  Only the first letter in this word is capital, like "Google".
 *  Otherwise, we define that this word doesn't use capitals in a right way.
 * */

class Solution {
public:
    bool detectCapitalUse(string word) {
        int lowercase = 0;
        int uppercase = 0;
        
        for(auto &c: word) {
            if(c <= 'Z') ++uppercase;
            else ++lowercase;
        }
        
        int size = word.size();
        
        if(size == lowercase)
            return true;
        if(size == uppercase)
            return true;
        if(uppercase == 1 && word[0] <= 'Z')
            return true;
        return false;
        
    }
};
