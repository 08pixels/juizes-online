#include <bits/stdc++.h>
using namespace std;

bool seen[8][8];
int move_x[] = { 1, -1, -1, 1, -2, -2,  2, 2};
int move_y[] = {-2, -2,  2, 2, -1,  1, -1, 1};


typedef struct {
    int x;
    int y;
} coordinate;

bool isValid(int coord) {
    return (coord >= 0 && coord <= 7);
}

void cleanSeenVector() {
    for (int i=0; i<8; ++i)
        for (int j=0; j<8; ++j)
            seen[i][j] = false;
}

int solve(coordinate source, coordinate target) {

    cleanSeenVector();

    queue<pair<int, coordinate> > bfsQueue;
    bfsQueue.push({0, source});

    seen[source.x][source.y] = true;

    while (!bfsQueue.empty()) {
        int level = bfsQueue.front().first;
        coordinate current = bfsQueue.front().second;

        bfsQueue.pop();

        for (int i=0; i<8; ++i) {
            int x = current.x + move_x[i];
            int y = current.y + move_y[i];

            if (isValid(x) && isValid(y) && !seen[x][y]) {

                if (x == target.x && y == target.y) {
                    while (!bfsQueue.empty())
                        bfsQueue.pop();

                    return 1 + level;
                }

                coordinate guess;
                guess.x = x;
                guess.y = y;

                bfsQueue.push({1 + level, guess});
                seen[x][y] = true;
            }
        }
    }

    return 0;
}

int main() {
    string op1, op2;

    while (cin >> op1 >> op2) {

        coordinate source, target;

        source.x = op1[0]-'a', source.y = op1[1]-'1';
        target.x = op2[0]-'a', target.y = op2[1]-'1';

        int result = solve(source, target);
        string output = "To get from " + op1 + " to " + op2 + " takes " + to_string(result) + " knight moves.";

        cout << output << endl;
    }


    return 0;
}