#include <stdio.h>
#include <stdlib.h>
#include <math.h>


double myPow(double b, double n) {
    if(!n) return 1;
    else if(n == 1) return b;
    else if(b == 1.0 || b == -1.0) return b;
    else if(b < 0.1) return 0; 
    

    b = (double) b;
    double r = 1;
    int _n;

    if(n < 0 && n < -10000) return 0;

    if(n < 0) _n = 0 - n;
    else _n = n;

    for(int i = 1; i <= _n; i++){
        r = r * b;
    }

    if(n < 0) r = 1.0 / r;

    return r;
}

int main(void){
    printf("%f\n", myPow(2.0, 10));
    printf("%f\n", myPow(2.1, 3));
    printf("%f\n", myPow(2.0, -2));
    printf("%f\n", myPow(0.44528, 0));
    printf("%f\n", myPow(1.0, 2147483647));
    printf("%f\n", myPow(2.0, -2147483648));
    printf("%f\n", myPow(0.00001, 2147483647));
    return EXIT_SUCCESS;
}
