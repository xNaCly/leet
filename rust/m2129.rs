use std::collections::HashMap;

trait Solution {}
impl dyn Solution {
    pub fn capitalize_title(title: String) -> String {
        title
            .split_whitespace()
            .map(|x| x.to_ascii_lowercase())
            .map(|x| match &x.len() {
                1..=2 => x,
                _ => x
                    .to_string()
                    .chars()
                    .enumerate()
                    .map(|(i, c)| if i == 0 { c.to_ascii_uppercase() } else { c })
                    .collect(),
            })
            .collect::<Vec<String>>()
            .join(" ")
    }
}

fn main() {
    let tests = HashMap::from([
        ("capiTalIze tHe titLe", "Capitalize The Title"),
        ("First leTTeR of EACH Word", "First Letter of Each Word"),
        ("i lOve leetcode", "i Love Leetcode"),
    ]);
    assert!(tests.iter().all(|e| {
        let got = <dyn Solution>::capitalize_title(e.0.to_string());
        let expected = *e.1;
        let cond = got == expected;
        if !cond {
            println!("got='{}', expected='{}'", got, expected);
        }
        cond
    }))
}
