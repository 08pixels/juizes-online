int dis[12][12];

int x[]={0,0,1,-1};
int y[]={1,-1,0,0};

class Solution {
public:
    int row, col;
    
    bool isValid(int u, int v) {
        return u >= 0 && v >= 0 && u < row && v < col;
    }
    
    void traverse(int u, int v, vector<vector<int>>& grid) {

        for(int i = 0; i < 4; ++i ) {
            int a = u + x[i];
            int b = v + y[i];
            
            if ( isValid(a, b) && abs(grid[a][b]) == 1) {
                if ( dis[a][b] > 1 + dis[u][v] ) {
                    grid[a][b] = -1;
                    dis[a][b] = 1 + dis[u][v];
                    traverse(a, b, grid);
                }
            }
        }
    }
    
    int orangesRotting(vector<vector<int>>& grid) {
        row = grid.size();
        col = grid[0].size();
        
        for(int i = 0; i < row; ++i)
            for(int j = 0; j < col; ++j)
                dis[i][j] = grid[i][j] != 2 ? (1 << 7) : 0;
        
        for(int i = 0; i < row; ++i) {
            for(int j = 0; j < col; ++j) {
                if ( grid[i][j] == 2 )
                    traverse(i, j, grid);
            }
        }
        
        int answer = 0;
        
        for(int i=0; i<row; ++i) {
            for(int j=0; j<col; ++j) {
                if ( grid[i][j] ==  1 ) return -1;
                if ( grid[i][j] == -1 ) answer = max(answer, dis[i][j]);
            }
        }
        
        return answer;
        
    }
};
