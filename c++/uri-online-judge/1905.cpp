#include <bits/stdc++.h>
#define MAXN 5
#define WALL 1
#define VISITED -1

using namespace std;

int x_moves[] = {0,  0, 1, -1};
int y_moves[] = {1, -1, 0,  0};


int grid[MAXN][MAXN];

bool isValid(int coord) {
    return coord >= 0 and coord < MAXN;
}

string solve(int x, int y) {
    if (x == (MAXN - 1) and y == (MAXN - 1))
        return "COPS";

    grid[x][y] = VISITED;

    for (int i=0; i<4; ++i) {
        int guess_x = x + x_moves[i];
        int guess_y = y + y_moves[i];

        if (!isValid(guess_x) or !isValid(guess_y))
            continue;
        
        if (grid[guess_x][guess_y] == WALL)
            continue;

        if (grid[guess_x][guess_y] == VISITED)
            continue;

        string result = solve(guess_x, guess_y);

        if (result == "COPS")
            return result;
    }

    return "ROBBERS";
}

int main(void) {

    int test;
    cin >> test;

    while (test--) {

        for (int i=0; i<MAXN; ++i) {
            for (int j=0; j<MAXN; ++j)
                cin >> grid[i][j];
        }

        cout << solve(0, 0) << endl;
    }

    return 0;
}