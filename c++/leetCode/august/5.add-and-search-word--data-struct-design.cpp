/*
*   Design a data structure that supports the following two operations:
*       void addWord(word)
*       bool search(word)

*   search(word) can search a literal word or a regular expression string
*   containing only letters a-z or .. A . means it can represent any one letter.
*/


typedef struct node {
    bool isEnd;
    unordered_map<char, struct node *> childs;
} node;


class WordDictionary {
public:
    node *root;

    WordDictionary() {
        root = new node;
        root->isEnd = false;
    }
    
    void addWord(string word) {
        node *trie = root;
        
        for(auto &w: word) {
            
            if(trie->childs.count(w) == 0) {
                trie->childs[w] = new node;
                trie->childs[w]->isEnd = false;
            }

            trie = trie->childs[w];
        }
        
        trie->isEnd = true;

    }
    
    bool solve(int ind, string word, node *tmp) {
        if(tmp == NULL)
            return false;
        
        if(ind == word.size())
            return tmp->isEnd;
        
        if(word[ind] == '.') {
            for(auto &k: tmp->childs)
                if(solve(ind + 1, word, k.second))
                    return true;
            
        } else {
            return solve(ind + 1, word, tmp->childs[word[ind]]);
        }
        
        return false;

    }
    
    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    bool search(string word) {
        node *tmp = root;
        return solve(0, word, tmp);
    }
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */
 