trait Solution {}
impl dyn Solution {
    pub fn shifting_letters(s: String, shifts: Vec<i32>) -> String {
        let mut r: Vec<char> = Vec::new();
        // TODO: none of this fucking shit works, why the fuck does it fail to iterate over the
        // chars and convert them to digits wtf
        for (i, c) in s.chars().enumerate() {
            let shift = shifts.get(i).unwrap();
            r.push(c);
            println!("{:?}", r);
            r = r
                .iter()
                .map(|x| {
                    char::from_digit(*x as u32 + *shift as u32, 10)
                        .expect("failed to convert int to char")
                })
                .collect::<Vec<char>>();
        }
        println!("{:?}", r);
        r.iter().collect()
    }
}

fn main() {
    let inputs = vec![("abc", vec![3, 5, 9]), ("aaa", vec![1, 2, 3])];
    let expected = vec!["rpl", "gfd"];

    for (i, v) in inputs.iter().enumerate() {
        assert!(
            <dyn Solution>::shifting_letters(String::from(v.0), v.1.to_vec())
                == expected.get(i).unwrap().to_string()
        );
    }
}
