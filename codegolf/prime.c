#include <stdio.h>
main(n,j,r){for(n=2;n<101;n++){for(j=2,r=0;j<n;j++)n%j==0&&(r=1);r||printf("%d\n",n);}}
