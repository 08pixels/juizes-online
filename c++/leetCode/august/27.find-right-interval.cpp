class Solution {
public:
    vector<int> findRightInterval(vector<vector<int>>& intervals) {
        int size = intervals.size();
        map<int, int> index;
        
        for(int i=0; i<size; ++i) {
            int curr = intervals[i][0];
            
            if ( index.count(curr) )
                index[curr] = min( index[curr], i );
            else
                index[curr] = i;
        }
        
        vector<int> answer(size, -1);
        for(int i=0; i<size; ++i) {
            auto it = index.lower_bound(intervals[i][1]);
            
            if ( it != index.end() )
                answer[i] = it->second;
        }
        
        return answer;
    }
};
