#include <bits/stdc++.h>
using namespace std;

int calls;

int fibonacci(int n) {
    if (n < 2)
        return n;

    calls += 2;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

int main(void) {

    int test, number;
    cin >> test;

    while (test--) {
        cin >> number;

        calls = 0;
        int result = fibonacci(number);

        cout << "fib(" << number << ") = " << calls << " calls = " << result << endl;
    }

}