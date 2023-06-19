// TODO: Time Limit exceeded
trait Solution {}
impl dyn Solution {
    pub fn shifting_letters(s: String, shifts: Vec<i32>) -> String {
        let mut r: Vec<char> = Vec::new();
        for (i, c) in s.chars().enumerate() {
            let shift = shifts.get(i).unwrap();
            r.push(c);
            r = r
                .iter()
                .map(|x| {
                    let chr = *x as u32;
                    let offset = if chr < 97 { 97 } else { chr - 97 };
                    char::from_u32(((offset + *shift as u32) % 26) + 97).unwrap()
                })
                .collect::<Vec<char>>();
        }
        r.iter().collect()
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
        if sol != expected.get(i).unwrap().to_string() {
            println!("{sol}");
        }
    }
}
