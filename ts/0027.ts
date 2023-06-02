let removeElement = (n: Array<number>, v: number) => {
  let c = 0;
  for (let i = 0; i < n.length; i++) {
    if (n[i] !== v) {
      n[c++] = n[i];
    }
  }
  return c;
};

console.log(removeElement([3, 2, 2, 3], 3));
console.log(removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2));
