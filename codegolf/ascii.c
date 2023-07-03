i;main(j,c){puts("   2 3 4 5 6 7
-------------");for(;i<16;i++){printf("%X: ",i);for(j=0;j<6;j++)c=i+32+(16*j),c==127?puts("DEL"):printf("%c ",c);puts("");}}
