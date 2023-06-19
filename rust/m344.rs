use std::collections::HashMap;

trait Solution {}
impl dyn Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        s.reverse();
    }
}

fn main() {
    let tests: HashMap<&str, &str> = HashMap::from([("hello", "olleh"), ("Hannah", "hannaH")]);
    assert!(tests.iter().all(|x| {
        let mut key = x.0.to_owned().chars().collect::<Vec<char>>();
        <dyn Solution>::reverse_string(&mut key);
        key.iter().collect::<String>() == x.1.to_owned()
    }));
}
