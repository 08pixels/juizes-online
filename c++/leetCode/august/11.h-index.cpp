class Solution {
public:
    int hIndex(vector<int>& citations) {
        int n = citations.size();
        if ( n == 0 )
            return 0;
        
        sort(citations.begin(), citations.end(), greater());
        
        int i = -1;
        while (1+i < n && citations[1 + i] > (1 + i))
            ++i;

        return 1 + i;
    }
};
