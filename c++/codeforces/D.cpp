// https://www.urionlinejudge.com.br/judge/en/problems/view/1301

#include <bits/stdc++.h>

int N, K, x,y, current;
int v[100005];
int tree[200005];

int check(int x) {
    return x > 0 ? 1 : x ? -1 : 0;
}

void build(int node, int i, int j)
{
    if(i == j) {
        tree[node] = check(v[i]);
    } else {

        int mid = (i+j) / 2;

        build(2 * node, i, mid);
        build(2 * node + 1, mid + 1, j);

        tree[node] = tree[2 * node] * tree[2 * node + 1];
    }
}

void update(int node, int i, int j, int index, int value)
{
    if(index > j || index < i) {
        return;
    } else if(i == j) {
        tree[node] = check(value);
    } else {
        int mid = (i+j) / 2;

        update(2 * node, i, mid, index, value);
        update(2 * node + 1, mid+1, j, index, value);

        tree[node] = tree[2 * node] * tree[2 * node + 1];
    }
}

int query(int node, int i, int j, int A, int B)
{
    if(i > B || j < A)
        return 1;

    if(A <= i && B >= j)
        return tree[node];

    int mid = (i+j) / 2;

    int side1 = query(2 * node, i, mid, A, B);
    int side2 = query(2 * node + 1, mid + 1, j, A, B);

    return side1 * side2;
}

int main()
{
    char c;

    while(scanf("%d %d", &N, &K) != EOF)
    {
        for(int i=1; i<=N; ++i)
            scanf("%d", v+i);

        build(1, 1, N);

        while(K--)
        {
            getchar();
            scanf("%c %d %d", &c, &x, &y);

            if(c == 'C')
                update(1, 1, N, x, y);
            else {
                int answer = query(1, 1, N, x, y);
                putchar(answer > 0 ? '+' : (answer ? '-' : '0'));
            }
        }
        putchar('\n');
    }

    return 0;
}