let maxVowels = (s: string, k: number): number => {
  let m = 0;
  for (let i = 0; i < s.length; i++) {
    let l = s
      .slice(i, k + i < s.length ? k + i : s.length)
      .split("")
      .filter((x) => {
        switch (x) {
          case "a":
          case "e":
          case "i":
          case "o":
          case "u":
            return true;
          default:
            return false;
        }
      }).length;

    if (l) m = l > m ? l : m;
  }
  return m;
};

console.log(maxVowels("abciiidef", 3) === 3);
console.log(maxVowels("aeiou", 2) === 2);
console.log(maxVowels("leetcode", 3) === 2);
console.log(maxVowels("rhythms", 4) === 0);
console.log(maxVowels("weallloveyou", 7) === 4);
