#include <assert.h>
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

// the easy way, lol - compiler builtin, supported in llvm & gcc
int hammingWeight(uint32_t n) { return __builtin_popcount(n); }

int hammingWeight1(uint64_t x) {
  int count;
  for (count = 0; x; count++)
    x &= x - 1;
  return count;
}

int main(void) {
  assert(hammingWeight(12) == hammingWeight1(12));
  return EXIT_SUCCESS;
}
