
class Solution {
public:
    vector<int> distributeCandies(int candies, int num_people) {
        
        vector<int> answer(num_people, 0);
        
        int curr  = 0;
        
        while(candies) {
            int value = min(1 + curr, candies);

            answer[curr % num_people] += value;
            candies -= value;
            
            ++curr;
        }
        
        return answer;
    }
};
