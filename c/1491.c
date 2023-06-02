#include <stdio.h>
#include <stdlib.h>

double average(int *s, int l) {
  int max = s[0], min = s[0], sum, nl;
  for (int i = 0; i < l; i++) {
    if (s[i] > max) {
      max = s[i];
    }
    if (s[i] < min) {
      min = s[i];
    }
  }

  for (int i = 0; i < l; i++) {
    if (s[i] == min || s[i] == max)
      continue;
    sum += s[i];
    nl++;
  }

  return sum / (double)nl;
}

int main(void) {
  int i[] = {4000, 3000, 1000, 2000};
  printf("%f\n", average(i, sizeof(i) / sizeof(int)));
  int i1[] = {1000, 2000, 3000};
  printf("%f\n", average(i1, sizeof(i1) / sizeof(int)));
  int i2[] = {48000, 59000, 99000, 13000, 78000, 45000, 31000,
              17000, 39000, 37000, 93000, 77000, 33000, 28000,
              4000,  54000, 67000, 6000,  1000,  11000};
  printf("%f\n", average(i2, sizeof(i2) / sizeof(int)));
  return EXIT_SUCCESS;
}
