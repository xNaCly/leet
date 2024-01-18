// TODO: im too fucking dumb to solve this WTF how the fuck is this easy
// "medium" - fuck this shit
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int mod(int a, int b) {
  int r = a % b;
  if (r < 0) {
    return r * -1;
  }
  return r;
}

char *revStr(char *str) {
  char tmp, *src, *dst;
  size_t len = strnlen(str, 10e9);
  src = str;
  dst = src + len - 1;
  while (src < dst) {
    tmp = *src;
    *src++ = *dst;
    *dst-- = tmp;
  }
  return str;
}

char *baseNeg2(int n) {
  char *str = malloc((sizeof(char) * 10e9));
  int r = 0;
  int i = 0;
  for (; n != 0; i++) {
    r = mod(n, -2);
    n = n / -2;
    str[i] = r + 48;
  }
  return revStr(str);
}

int main(void) {
  char *r = baseNeg2(2);
  puts(r);
  assert(strncmp(r, "110", 2) == 0);
  free(r);
  r = baseNeg2(3);
  assert(strncmp(baseNeg2(3), "111", 3) == 0);
  free(r);
  r = baseNeg2(4);
  assert(strncmp(baseNeg2(4), "100", 3) == 0);
  free(r);
  return EXIT_SUCCESS;
}
