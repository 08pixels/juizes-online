/**
*    You are given a char array representing tasks CPU need to do.
*    It contains capital letters A to Z where each letter represents a different task.
*    Tasks could be done without the original order of the array.
*    Each task is done in one unit of time. For each unit of time,
*    the CPU could complete either one task or just be idle.
*    However, there is a non-negative integer n that represents the cooldown period
*    between two same tasks (the same letter in the array), that is that there must be
*    at least n units of time between any two same tasks.
*    You need to return the least number of units of times that
*    the CPU will take to finish all the given tasks.
**/

class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        if(n == 0) return tasks.size();
            
        vector<int> frequency(30, 0);
        vector<int> priority;
        
        for( auto &task: tasks ) {
            ++frequency[task - 'A'];
        }
        
        
        for( auto &amount: frequency ) {
            if(amount)
                priority.push_back(amount);
        }
        
        int idle = 0;
        int size = priority.size();

        sort( priority.begin(), priority.end(), greater<int>() );
        
        while( size ) {

            int amount = min(n + 1, size);
            
            for( int i = 0; i < amount ; ++i ) {
                --priority[i];
            }
            
            
            sort( priority.begin(), priority.end(), greater<int>() );
            
            while(priority.size() && priority[priority.size() - 1] == 0 ){
                priority.pop_back();
            }
            
            size = priority.size();
            
            if( size && amount < (n + 1) ) {
                idle += (n + 1) - amount;
            }
        }

        
        return tasks.size() + idle;
    }
};