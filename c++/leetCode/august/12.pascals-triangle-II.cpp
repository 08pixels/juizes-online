class Solution {
public:
    vector<int> getRow(int rowIndex) {        
        vector<int> line(rowIndex + 1);
        vector<int> tmp(rowIndex + 1);
        
        line[0] = tmp[0] = 1;
        
        for(int i = 1; i <= rowIndex; ++i) {
            for(int j = 1; j <= i; ++j)
                tmp[j] = line[j] + line[j - 1];

            line = tmp;
        }
        
        return line;
    }
};
