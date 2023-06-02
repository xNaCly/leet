#include <stdlib.h>
#include <stdio.h>

int lengthOfLastWord(char * s) {
    int c = 0;
    char l;
    for(int i = 0;;i++){
        // if at end of the string, exit loop
        if(s[i] == '\0') break;
        // if last char is a space and current char isnt a space, reset counter
        else if(l == ' ' && s[i] != ' ') c = 0;
        // if current char is not space increment counter
        if(s[i] != ' ') c++;
        // set last char to current char;
        l = s[i];
    }
    return c;
}

int main(void){
    printf("%d\n", lengthOfLastWord("Hello World"));
    printf("%d\n", lengthOfLastWord("fly me   to   the moon  "));
    printf("%d\n", lengthOfLastWord("luffy is still joyboy"));
    return EXIT_SUCCESS;
}
