Array.prototype.last = function () {
  if (!this.length) return -1;
  return this[this.length - 1];
};

let arr = [1, 2, 3];
console.log(arr.last());
arr = [null, {}, 3, {}];
console.log(arr.last());
arr = [];
console.log(arr.last());
