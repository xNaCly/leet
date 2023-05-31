use std::collections::HashMap;

trait Solution {}
impl dyn Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut n: HashMap<i32, usize> = HashMap::new();
        for i in nums {
            n.entry(i).and_modify(|val| *val += 1).or_insert(1);
        }
        for (key, value) in n {
            if value == 1 {
                return key;
            }
        }
        panic!("cant hit this")
    }
}

fn main() {
    let tests = HashMap::from([(1, vec![2, 2, 1]), (4, vec![4, 1, 2, 1, 2]), (1, vec![1])]);

    for (expected, input) in &tests {
        let res = <dyn Solution>::single_number(input.to_vec());
        if res != *expected {
            panic!("{} not equal to {}", res, expected)
        }
    }
}
