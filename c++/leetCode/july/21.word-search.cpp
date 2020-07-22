
int dx[] = {0, 0, 1,-1};
int dy[] = {1,-1, 0, 0};

class Solution {
public:
    int seen[401][401];
    vector<vector<char>> board;
    
    int height;
    int width;
    int it;

    
    bool solve(int x, int y, int curr, string& word) {
        if(curr == word.size())
            return true;
        
        for(int i=0; i<4; ++i) {
            int px = x + dx[i];
            int py = y + dy[i];
            
            if(px >= 0 && px < height && py >= 0 && py < width && seen[px][py] != it && word[curr] == board[px][py]) {
                seen[x][y] = it;                
                if(solve(px, py, curr + 1, word)) return true;
                seen[x][y] = it-1;
            }
        }
        
        return false;
        
    }
    
    bool exist(vector<vector<char>>& board, string word) {
        this->board  = board;
        this->height = board.size();
        this->width  = board[0].size();
        this->it     = 0;
        
        for(int i=0; i<400; ++i)
            for(int j=0; j<400; ++j)
                seen[i][j] = 0;
        
        for(int i=0; i<height; ++i) {
            for(int j=0; j<width; ++j) {
                if(board[i][j] != word[0])
                    continue;
                
                ++it;
                if(solve(i,j, 1, word))
                    return true;
            }
        }
        
        return false;
    }
};