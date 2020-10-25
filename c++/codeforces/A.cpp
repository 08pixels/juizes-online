#include <bits/stdc++.h>
#define MAX 2100000
using namespace std;

typedef struct {
	int x,y,z;
}node;

char str[1000006];
int o[MAX], c[MAX], v[MAX];

void build(int ind, int i, int j)
{
	if(j == i) {
		if(str[i] == '(')
			o[ind] = 1;
		else
			c[ind] = 1;

	} else {
		int mid = (i+j) >> 1;

		build(2 * ind, i, mid);
		build(2 * ind + 1, mid + 1, j);

		int current = min(o[2 * ind], c[2 * ind + 1]);

		v[ind] = v[2 * ind] + v[2 * ind + 1] + current;
		o[ind] = o[2 * ind] + o[2 * ind + 1] - current;
		c[ind] = c[2 * ind] + c[2 * ind + 1] - current;
	}
}

node query(int ind, int i, int j, int q0, int qf)
{
	if(j < q0 || i > qf) {
		node s;
		s.x = s.y = s.z = 0;
		return s;
	}

    if(i >= q0 && j <= qf) {
    	node s;
		s.x = v[ind];
		s.y = o[ind];
		s.z = c[ind];
		return s;
    }

	int mid = (i+j) >> 1;

	node l = query(2 * ind, i, mid, q0, qf);
	node r = query(2 * ind + 1, mid + 1, j, q0, qf);

	int current = min(l.y, r.z);

	node answer;
	answer.x = l.x + r.x + current;
	answer.y = l.y + r.y - current;
	answer.z = l.z + r.z - current;
	return answer;
}

int main()
{
	int Q, q0,qf;

	scanf("%s %d", str, &Q);
	int size = strlen(str) - 1;
	build(1, 0, size);

	while(Q--) {
		scanf("%d %d", &q0, &qf);
		printf("%d\n", query(1, 0, size, q0-1, qf-1).x * 2);
	}

	return 0;
}