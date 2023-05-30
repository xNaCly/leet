let runningSum = (n: number[]) => {
  let t = 0;
  for (let i = 0; i < n.length; i++) {
    t += n[i];
    n[i] = t;
  }
  return n;
};

console.log(runningSum([1, 2, 3, 4]));
console.log(runningSum([1, 1, 1, 1, 1]));
console.log(runningSum([3, 1, 2, 10, 1]));
