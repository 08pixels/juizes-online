#include <bits/stdc++.h>
#define MAX 30010
using namespace std;

int v[MAX];
vector<int> tree[3 * MAX];

void build(int ind, int start, int end)
{
	if(start == end) {
		tree[ind].push_back(v[start]);
	} else {
		int mid = (start + end) >> 1;

		build(ind << 1, start, mid);
		build(ind << 1 | 1, mid + 1, end);

		tree[ind].resize(tree[ind << 1].size() + tree[ind << 1 | 1].size());

		merge(	tree[ind << 1].begin(), tree[ind << 1].end(), 
				tree[ind << 1 | 1].begin(), tree[ind << 1 | 1].end(),
				tree[ind].begin()
				);
	}
}

int query(int ind, int start, int end, int q0, int qf, int key)
{
	if(start > qf || end < q0)
		return 0;

	if(start >= q0 && end <= qf) {
		vector<int>::iterator current = upper_bound(tree[ind].begin(), tree[ind].end(), key);

		if(current != tree[ind].end())
			return tree[ind].size() - (current - tree[ind].begin());
		return 0;
	}

	int mid = (start + end) >> 1;

	int l = query(ind << 1, start, mid, q0, qf, key);
	int r = query(ind << 1 | 1, mid + 1, end, q0, qf, key);

	return l + r;
}

int main()
{
	int N, Q;
	int x,y,z;

	scanf("%d", &N);

	for(int i=1; i<=N; ++i)
		scanf("%d", v + i);

	build(1, 1, N);
	scanf("%d", &Q); ++Q;

	int last = 0;
	while(--Q) {
		scanf("%d %d %d", &x, &y, &z);
		last = query(1, 1, N, x^last, y^last, z^last);
		printf("%d\n", last);
	}

	return 0;
}