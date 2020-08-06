/**
*   Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
*   some elements appear twice and others appear once.
*
*   Find all the elements that appear twice in this array.
*   Could you do it without extra space and in O(n) runtime?
**/

#define modInd abs(nums[i]) - 1

class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> result;
        
        for(int i = 0; i < nums.size(); ++i) {
            if( nums[modInd] < 0 )
              result.push_back(modInd + 1);
            else
              nums[modInd] = -nums[modInd];
        }
        
        return result;
    }
};
