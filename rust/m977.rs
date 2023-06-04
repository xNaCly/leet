trait Solution {}
impl dyn Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let mut p = nums
            .clone()
            .into_iter()
            .map(|x| x.pow(2))
            .collect::<Vec<i32>>();
        p.sort();
        return p;
    }
}

fn compare_vecs(a: Vec<i32>, b: Vec<i32>) {
    assert_eq!(a.len(), b.len());
    for i in 0..a.len() {
        assert_eq!(a.get(i), b.get(i));
    }
}

fn main() {
    let tests = vec![
        vec![vec![-4, -1, 0, 3, 10], vec![0, 1, 9, 16, 100]],
        vec![vec![-7, -3, 2, 3, 11], vec![4, 9, 9, 49, 121]],
    ];
    for i in tests {
        let o = <dyn Solution>::sorted_squares(i.get(0).unwrap().to_vec());
        compare_vecs(o, i.get(1).unwrap().to_vec())
    }
}
