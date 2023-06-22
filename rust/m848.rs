trait Solution {}
impl dyn Solution {
    pub fn shifting_letters(s: String, shifts: Vec<i32>) -> String {
        s.bytes()
            .zip(&shifts)
            .scan(
                shifts.iter().fold(0, |acc, &x| (acc + x) % 26),
                |state, (u, shift)| {
                    let s = *state as u8;
                    *state = (*state - shift).rem_euclid(26);
                    Some((b'a' + (u - b'a' + s) % 26) as char)
                },
            )
            .collect()
    }
}

fn main() {
    let inputs = vec![
        ("abc", vec![3, 5, 9]),
        ("aaa", vec![1, 2, 3]),
        ("bad", vec![10, 20, 30]),
    ];
    let expected = vec!["rpl", "gfd", "jyh"];

    for (i, v) in inputs.iter().enumerate() {
        let sol = <dyn Solution>::shifting_letters(String::from(v.0), v.1.to_vec());
        let exp = expected.get(i).unwrap().to_string();
        if sol != exp {
            panic!("{} != {}", sol, exp);
        }
    }
}
