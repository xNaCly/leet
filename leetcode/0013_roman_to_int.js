let m = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
}

let romanToInt = (s) => {
    let r = m[s[s.length-1]];
    let l = s[s.length-1];
    for(let i = s.length-2; i > -1; i--){
        let cc = s[i];
        console.log({i: i, cc: cc, l: l, r: r});
        if((l == "X" || l == "V") && cc == "I") r -= 1;
        else if((l == "L" || l == "C") && cc == "X") r -= 10;
        else if((l == "D" || l == "M") && cc == "C") r -= 100;
        else r += m[cc];
        l = cc;
    }
    console.log({l: l, r: r});
    return r;
}

console.log(romanToInt("III"));
console.log(romanToInt("LVIII"));
console.log(romanToInt("MCMXCIV"));
