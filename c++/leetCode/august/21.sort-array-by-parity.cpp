class Solution {
public:
    vector<int> sortArrayByParity(vector<int>& A) {
        int l = 0;
        int r = A.size() - 1;

        while ( l < r ) {
            while ( l < r && !(A[l] & 1) )
                ++l;
            while ( l < r && (A[r] & 1) )
                --r;
            
            if ( l < r )
                A[l] ^= A[r], A[r] ^= A[l], A[l++] ^= A[r--];
        }
        
        return A;
    }
};
