use std::collections::HashMap;

trait Solution {}
impl dyn Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack: Vec<char> = Vec::new();
        let o = HashMap::from([(']', '['), (')', '('), ('}', '{')]);
        for c in s.chars() {
            match c {
                '(' => stack.push(c),
                '[' => stack.push(c),
                '{' => stack.push(c),
                _ => {
                    if stack.iter().last() == o.get(&c) {
                        stack.pop();
                    } else {
                        return false;
                    }
                }
            }
        }
        stack.is_empty()
    }
}

fn main() {
    let tests = HashMap::from([("()", true), ("()[]{}", true), ("(]", false)]);

    assert!(tests
        .values()
        .into_iter()
        .all(|s| <dyn Solution>::is_valid(s.to_string())));
}
