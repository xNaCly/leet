#include <stdio.h>
#include <stdlib.h>
#include <math.h>


double myPow(double b, double n) {
    if(!n) return 1;
    else if(n == 1) return b;

    n = (double) n;
    b = (double) b;
    double r = b;
    if(n < 0){
        for(int i = n; i < 1; i++){
            r /= b;
        }
    } else {
        for(int i = n - 1; i > 0; i--){
            r *= b;
        }
    }

    return r;
}

int main(void){
    printf("%f\n", myPow(2.0, 10));
    printf("%f\n", myPow(2.1, 3));
    printf("%f\n", myPow(2.0, -2));
    printf("%f\n", myPow(0.44528, 0));
    return EXIT_SUCCESS;
}
