trait Solution {}
impl dyn Solution {
    pub fn average(mut salary: Vec<i32>) -> f64 {
        salary.sort();
        salary = salary
            .split_first()
            .unwrap()
            .1 // remove first element
            .split_last()
            .unwrap()
            .1 // remove last element
            .to_vec();
        let l = salary.len() as f64;
        return (salary.iter().map(|v| *v as f64).sum::<f64>() / l).into();
    }
}

fn main() {
    let tests = vec![
        (vec![4000, 3000, 1000, 2000], 2500.00),
        (vec![1000, 2000, 3000], 2000.0),
    ];

    for (input, expected) in &tests {
        let res = <dyn Solution>::average(input.to_vec());
        if res != *expected {
            panic!("{} not equal to {}", res, expected)
        }
    }
}
