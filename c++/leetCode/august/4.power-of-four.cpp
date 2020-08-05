/**
*   Given an integer (signed 32 bits), write a function
*   to check whether it is a power of 4.
*
*   Follow up: Could you solve it without loops/recursion?
*/

class Solution {
public:
    bool isPowerOfFour(int num) {
        if( num <= 0 )
            return false;

        int mask = 0b1010101010101010101010101010101;
        int lsb = num & (-num);
        
        if( num - lsb )
            return false;

        return lsb & mask;
    }
};
