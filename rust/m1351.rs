use std::collections::HashMap;

trait Solution {}
impl dyn Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        grid.iter()
            .flat_map(|a| a.iter())
            .filter(|x| x.is_negative())
            .collect::<Vec<&i32>>()
            .len() as i32
    }
}

fn main() {
    let tests = HashMap::from([
        (
            8,
            vec![
                vec![4, 3, 2, -1],
                vec![3, 2, 1, -1],
                vec![1, 1, -1, -2],
                vec![-1, -1, -2, -3],
            ],
        ),
        (0, vec![vec![3, 2], vec![1, 0]]),
    ]);
    assert!(tests
        .iter()
        .all(|x| *(x.0) == <dyn Solution>::count_negatives(x.1.to_vec())))
}
