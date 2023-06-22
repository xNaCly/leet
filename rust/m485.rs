// INFO:
// runtime: 3ms -> beats 89.5%
// mem: 2mb -> beats 94%
trait Solution {}
impl dyn Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let mut max = 0;
        let mut cc = 1;
        for (i, num) in nums.iter().enumerate() {
            if *num == 0 {
                if max < cc - 1 {
                    max = cc - 1;
                }
                cc = 1;
            } else if i == nums.len() - 1 {
                if max < cc {
                    max = cc;
                }
                cc = 1;
            } else {
                cc += 1;
            }
        }
        max
    }
}

fn main() {
    let tests = Vec::from([
        (3, vec![1, 1, 0, 1, 1, 1]),
        (2, vec![1, 0, 1, 1, 0, 1]),
        (2, vec![1, 1, 0, 1]),
    ]);
    assert!(tests.iter().all(|e| {
        let sol = <dyn Solution>::find_max_consecutive_ones(e.1.to_vec());
        let exp = e.0;
        println!("got={}, wanted={}", sol, exp);
        sol == exp
    }));
}
