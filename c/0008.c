#include <assert.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

int myAtoi(char *s) {
  if (*s == ' ') {
    return myAtoi(s + 1);
  }

  int isNegative = *s == '-';
  if (*s == '-' || *s == '+') {
    s++;
  }

  while (*s == '0') {
    s++;
  }

  int n = 0;
  for (char *p = s; *p; ++p) {
    if (*p < '0' || *p > '9') {
      break;
    }
    n++;
  }

  if (n == 0) {
    return 0;
  } else if (n > 10) {
    return isNegative ? INT_MIN : INT_MAX;
  } else {
    long long int result = 0LL;
    for (int i = 0; i < n; i++) {
      // "42" as input
      // 4 -> 0*10 + 4, result = 4
      // 2 -> 4*10 + 2, result = 2

      // "4193" as input
      // 4 -> 0*10 + 4, result = 4
      // 1 -> 4*10 + 1, result = 41
      // 9 -> 41*10 + 9, result = 419
      // 3 -> 419*10 + 9, result = 4193
      result = result * 10 + (s[i] - 48);
      if (n == 10) {
        if (!isNegative && (result > INT_MAX)) {
          return INT_MAX;
        }
        if (isNegative && (-result < INT_MIN)) {
          return INT_MIN;
        }
      }
    }
    return isNegative ? (-result) : result;
  }
}

int main(void) {
  assert(42 == myAtoi("42"));
  assert(-42 == myAtoi("-42"));
  assert(4193 == myAtoi("4193 with words"));
  assert(-42 == myAtoi("  -42"));
  return EXIT_SUCCESS;
}
