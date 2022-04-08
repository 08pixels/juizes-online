#include <bits/stdc++.h>
#define MAXN 1024
#define SEEN 'o'

using namespace std;

int move_x[] = {0,  0, 1, -1};
int move_y[] = {1, -1, 0,  0};

char grid[MAXN][MAXN];


bool is_valid(int coord, int length) {
    return coord >= 0 and coord < length;
}

void flood_fill(int x, int y, int height, int width) {
    grid[x][y] = SEEN;

    for (int i=0; i<4; ++i) {
        int new_x = x + move_x[i];
        int new_y = y + move_y[i];

        if (is_valid(new_x, height) and is_valid(new_y, width) and grid[new_x][new_y] != SEEN)
            flood_fill(new_x, new_y, height, width);
    }

    return ;
}

int main(void) {

    int answer = 0;
    int height, width;

    cin >> height >> width;

    for (int i=0; i<height; ++i)
        cin >> grid[i];

    for (int i=0; i<height; ++i) {
        for (int j=0; j<width; ++j) {

            if (grid[i][j] != SEEN) {
                ++answer;
                flood_fill(i, j, height, width);
            }
        }
    }

    cout << answer << endl;
    cout << (int)'o' << endl;

    return 0;
}