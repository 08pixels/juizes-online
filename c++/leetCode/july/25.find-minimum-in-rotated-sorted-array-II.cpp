/*
    Suppose an array sorted in ascending order is rotated
    at some pivot unknown to you beforehand.
    (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
    Find the minimum element.
    The array may contain duplicates.
*/

class Solution {
public:
    int findMin(vector<int>& nums) {
        int i = 0;
        int j = nums.size();

        while(++i < j && nums[i-1] >= nums[i]);
        while(--j > i && nums[j-1] <= nums[j]);
        
        return nums[i-1] < nums[j] ? nums[i-1] : nums[j];
    }
};
