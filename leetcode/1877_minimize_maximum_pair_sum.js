let io = {
	i: [
		[3, 5, 2, 3],
		// 0 1 2 3
		[3, 5, 4, 2, 4, 6],
	],
	o: [7, 8],
};

let miniPairSum = (n) => {
	// use this after minimising the maximum pair sum
	let m = 0;
	for (let i = 0; i < n.length; i++) {
		if ((i + 1) % 2 == 0 && i != 0) {
			let t = n[i] + n[i - 1];
			m = t > m ? t : m;
			console.log(m);
		}
	}
	return m;
};

for (let i = 0; i < 1; i++) {
	let o = miniPairSum(io.i[i]);
	console.log({ o, r: o === io.o[i] });
}
