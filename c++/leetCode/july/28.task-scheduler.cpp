class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        if(n == 0) return tasks.size();
            
        vector<int> frequency(30, 0);
        
        for(auto &task: tasks)
            ++frequency[task - 'A'];

        while(true) {
            int unique = 0;

            for(int i=0; i<26; ++i) {
                if(frequency[i]) {
                    unique |= (1 << i);
                    --frequency[i];
                    
                    if(__builtin_popcount(unique) == (n + 1))
                        break;
                }
            }
            
            if(__builtin_popcount(unique) != (n + 1)) {
                
                for(int i=0; i<26; ++i)
                    if( ( n & ( 1 << i ) ) )
                        ++frequency[i];
                
                break;
            }
            
        }
        
        
        int worst = 1;

        for(auto &amount: frequency)
            worst = max(worst, amount);
        
        return tasks.size() + (worst - 1);
    }
};

