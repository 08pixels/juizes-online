https://www.hackerearth.com/practice/codemonk/?purpose=login&source=code_monk&update=google

#include <stdio.h>
#include <string.h>

int main(){
	int num;
	scanf("%d", &num);

	char v[num  + 1][15];

	for(int i=0; i<num; ++i) {
		char key[15]; scanf("%s", key);
		
		int k = i - 1;

		while(k >= 0 && strcmp(v[k], key) >= 0) {
            strcpy(v[k + 1], v[k]);
			--k;
		}

        strcpy(v[k+1], key);
		printf("%d\n", k + 1);
	}
}
