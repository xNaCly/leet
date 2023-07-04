#include <stdio.h>

// niven numbers
s,j;main(i){for(;i<101;i++){for(j=i,s=0;j>0;j/=10)s+=j%10;i%s==0&&printf("%d\n",i);}}
// niven numbers long
s,j;main(i){for(;i<1e4+1;i++){for(j=i,s=0;j>0;j/=10)s+=j%10;i%s==0&&printf("%d\n",i);}}
