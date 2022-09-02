import re

class Solution:
  def __init__(self) -> None:
    self.r = re.compile(r'\s+')

  def reverseWords(self, s: str) -> str:
    s = (" ".join(s.split(" ")[::-1])).strip()
    return self.r.sub(" ", s)

S = Solution();
print(S.reverseWords("the sky is blue"))
print(S.reverseWords("  hello world  "))
print(S.reverseWords("a good   example"))