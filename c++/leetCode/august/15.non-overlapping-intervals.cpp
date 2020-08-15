
bool compare(vector<int> &a, vector<int> &b) {
    if ( a[0] ^ b[0] )
        return a[0] < b[0];

    return a[1] <= b[1];
}

class Solution {
public:

    int eraseOverlapIntervals(vector<vector<int>>& intervals) {
        int n = intervals.size();
        sort(intervals.begin(), intervals.end(), compare);
        
        int answer = 0;
        int l      = 0;
        
        for(int r = 1; r < n; ++r) {
            
            if ( intervals[l][1] <= intervals[r][0] ) {
                l = r;
                continue;
            }
            
            if ( intervals[l][1] > intervals[r][1] )
                l = r;
            
            ++answer;
        }
        
        return answer;
    }
};
