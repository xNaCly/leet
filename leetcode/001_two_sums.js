let r = (n, t) => {
	let m = {};
    // insert all elements in n array in hashmap m
	for (let i = 0; i < n.length; i++) {
		m[n[i]] = i;
	}

	for (let i = 0; i < n.length; i++) {
        // calc complement
		let c = t - n[i];
        // if c is in keys of hashmap and doesnt have a value of i return indexes
		if (
			Object.keys(m)
				.map((x) => parseInt(x))
				.includes(c) &&
			m[c] != i
		) {
			return [i, m[c]];
		}
	}

	return null;
};

// console.log(r([2, 7, 11, 15], 9));
console.log(r([3, 2, 4], 6));
// console.log(r([3, 3], 6));
// console.log(r([3, 2, 3], 6));
