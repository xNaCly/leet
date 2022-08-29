let myPow = (b, n) => {
	switch (n) {
		case 1:
			return b;
		case 0:
			return 1;
		default:
			return Math.pow(b, n);
	}
};

console.log(myPow(2.0, 10));
console.log(myPow(2.1, 3));
console.log(myPow(2.0, -2));
console.log(myPow(0.44528, 0));
