#include <stdio.h>
#include <math.h>

int mySqrt(int x){
    return (int)sqrt(x);
}

int main(void) {
    printf("%d", mySqrt(8));
    return 0;
}