#include <bits/stdc++.h>
#define MAX 330010
using namespace std;

typedef struct {
	int l,r,key, index, type, answer;
}data;

data v[MAX];
int tree[30004], answer[200005], N, Q;

bool compare(data s1, data s2) {
	return (s1.key == s2.key) ? (s1.type > s2.type) : (s1.key > s2.key);
}

void complete(int i, int type, int index) {
	v[i].type = type;
	v[i].index = index;
}

void update(int x) {
	while(x <= N) {
		++tree[x];
		x += x&(-x);
	}
}

int query(int x) {
	int current = 0;

	while(x) {
		current += tree[x];
		x -= x&(-x);
	}
	return current;
}

int main()
{
	scanf("%d", &N);

	for(int i=1; i<=N; ++i) {
		scanf("%d", &v[i].key);
		complete(i, 0, i);
	}
	scanf("%d", &Q);

	for(int i=N+1; i<=N+Q; ++i) {
		scanf("%d %d %d", &v[i].l, &v[i].r, &v[i].key);
		complete(i, 1, i - (N+1));
	}

	sort(v + 1, (v + N + Q + 1), compare);
	for(int i=1; i<=N+Q; ++i)
	{
		if(v[i].type)
			answer[ v[i].index ] = query(v[i].r) - query(v[i].l - 1);
		else
			update(v[i].index);
	}

	for(int i=0; i<Q; ++i)
		printf("%d\n", answer[i]);
	return 0;
}