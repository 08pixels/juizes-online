// https://www.urionlinejudge.com.br/judge/pt/problems/view/1023

#include<stdio.h>

int main() {

  int city = 0;
  int storage[205];
  int aHouses, aPeople, consumption;

  while(1) {
    scanf("%d", &aHouses);
    
    if(aHouses == 0) break;
    int i;
    for(i=0; i<=200; ++i)
      storage[i] = 0;
    
    double sum = 0.0;


    for(i=0; i<aHouses; ++i) {
      scanf("%d %d", &aPeople, &consumption);
      storage[consumption/aPeople] += aPeople;
      sum += consumption;
    }

    int people = 0;

    if(city) printf("\n");
    printf("Cidade# %d:\n", ++city);


    for(i=0; i<=200; ++i) {
      if(storage[i] != 0) {
        if(people != 0) printf(" %d-%d", storage[i], i);
        else printf("%d-%d", storage[i], i);

        people += storage[i];
      }
    }
    printf("\n");

    float average = sum/(float)people;
    long long toInteger = ((long long)(average*100L));

    printf("Consumo medio: %.2f m3.\n", toInteger/100.00);
  }
  return 0;
}
