#include <stdio.h>
#include <stdlib.h>

void x_swap(int *a, int *b) {
  int t = *a;
  *a = *b;
  *b = t;
}

int x_partition(int *arr, int low, int high) {
  int p = arr[high];
  int i = low - 1;
  for (int j = low; j < high; j++) {
    if (arr[j] <= p) {
      i++;
      x_swap(&arr[i], &arr[j]);
    }
  }
  x_swap(&arr[i + 1], &arr[high]);
  return i + 1;
}

void x_sort(int *n, int l, int h) {
  if (l < h) {
    int p = x_partition(n, l, h);
    x_sort(n, l, p - 1);
    x_sort(n, p + 1, h);
  }
}

void sortColors(int *n, int l) { x_sort(n, 0, l - 1); }

void x_printArray(int *n, int l) {
  printf("[");
  for (int i = 0; i < l; i++) {
    printf("%d%c", n[i], i + 1 == l ? 0 : ',');
  }
  printf("]\n");
}

int main(void) {
  int i[] = {2, 0, 2, 1, 1, 0};
  int i1[] = {2, 0, 1};
  x_printArray(i, sizeof(i) / sizeof(i[0]));
  sortColors(i, sizeof(i) / sizeof(i[0]));
  x_printArray(i, sizeof(i) / sizeof(i[0]));

  x_printArray(i1, sizeof(i1) / sizeof(i1[0]));
  sortColors(i1, sizeof(i1) / sizeof(i1[0]));
  x_printArray(i1, sizeof(i1) / sizeof(i1[0]));
  return EXIT_SUCCESS;
}
