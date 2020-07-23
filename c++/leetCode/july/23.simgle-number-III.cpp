class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        vector<int> v;
        unordered_map<int, int> frequency;
        
        for(int i=0; i<nums.size(); ++i) {
            if(frequency[nums[i]])
                frequency.erase(nums[i]);
            else 
                ++frequency[nums[i]];
        }
        
        for(auto &e: frequency)
            v.push_back(e.first);
        
        return v;
    }
};
