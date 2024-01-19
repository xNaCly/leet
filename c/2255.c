#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int countPrefixes(char **words, int wordsSize, char *s) {
  int c = 0;
  for (int i = 0; i < wordsSize; i++) {
    strncmp(words[i], s, strlen(words[i])) == 0 && c++;
  }
  return c;
}

int main(void) {
  typedef struct {
    char *in;
    char **words;
    int wordCount;
    int count;
  } TestCase;
  TestCase testcases[] = {
      {
          .in = "abc",
          .words = (char *[]){"a", "b", "c", "ab", "bc", "abc"},
          .wordCount = 6,
          .count = 3,
      },
      {
          .in = "aa",
          .words = (char *[]){"a", "a"},
          .wordCount = 2,
          .count = 2,
      },
      {
          .in = "w",
          .words = (char *[]){"feh",  "w",  "w", "lwd",  "c",   "s", "vk",
                              "zwlv", "n",  "w", "sw",   "qrd", "w", "w",
                              "mqe",  "w",  "w", "w",    "gb",  "w", "qy",
                              "xs",   "br", "w", "rypg", "wh",  "g", "w",
                              "w",    "fh", "w", "w",    "sccy"},
          .wordCount = 33,
          .count = 14,
      },
      {
          .in = "vomy",
          .words = (char *[]){"ycwj", "hak",  "v",    "kjg", "zw",   "vtes",
                              "vom",  "ii",   "as",   "v",   "vo",   "v",
                              "w",    "vomy", "loa",  "fbm", "j",    "v",
                              "vo",   "e",    "y",    "t",   "eh",   "yk",
                              "bt",   "x",    "vomy", "vom", "yab",  "v",
                              "sydi", "wnb",  "z",    "v",   "iygp", "vomy"},
          .wordCount = 36,
          .count = 13,
      },
  };
  for (int i = 0; i < sizeof(testcases) / sizeof(TestCase); i++) {
    TestCase c = testcases[i];
    int count = countPrefixes(c.words, c.wordCount, c.in);
    if (count != c.count) {
      printf("Expected %d, got %d for '%s'\n", c.count, count, c.in);
      exit(1);
    }
  }
  return EXIT_SUCCESS;
}
