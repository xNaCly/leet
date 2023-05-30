let myPow = (b: number, n: number): number => {
	if (n == 0) return 1;
	else if (n == 1) return b;
	else if (n < 0) {
		// todo: implement newton's method here:
		return -1;
	}

	let r = b;
	for (let i = n - 1; i > 0; i--) {
		r *= b;
	}
	return r;
};

// console.log(myPow(2.0, 10));
// console.log(myPow(2.1, 3));
console.log(myPow(2.0, -2));
// console.log(myPow(0.44528, 0));
