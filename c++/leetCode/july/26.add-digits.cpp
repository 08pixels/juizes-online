/*
    Given a non-negative integer num, repeatedly
    add all its digits until the result has only one digit.
*/


class Solution {
public:
    int addDigits(int num) {
        
        while(num > 9) {
            int sum = 0;
            while(num) {
                sum += num % 10;
                num /= 10;
            }
            num = sum;
        }
        
        return num;
    }
};
