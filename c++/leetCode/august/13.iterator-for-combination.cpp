class CombinationIterator {
public:
    string characters;
    int length;
    int combinationLength;
    vector<string> answer;
    string tmp;
    int top;
    
    
    void solve(int ind, int amount) {

        if ( amount == combinationLength ) {
            answer.push_back(tmp);
            return ;
        }
            
        for(int i = ind + 1; i < length; ++i) {
            tmp[amount] = characters[i];
            solve(i, amount + 1);
        }
    }
 
    CombinationIterator(string characters, int combinationLength) {
        this->characters = characters;
        this->length = characters.size();
        this->combinationLength = combinationLength;
        this->tmp.resize(combinationLength);
        this->top = -1;
        
        for(int i=0; i<=length-combinationLength; ++i) {
            tmp[0] = characters[i];
            solve(i, 1);
        }
        
    }
    
    string next() {
        return answer[++top];
    }
    
    bool hasNext() {
        return (this->top+1) < answer.size();
    }
};

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * CombinationIterator* obj = new CombinationIterator(characters, combinationLength);
 * string param_1 = obj->next();
 * bool param_2 = obj->hasNext();
 */
