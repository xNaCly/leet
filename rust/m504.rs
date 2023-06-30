trait Solution {}
impl dyn Solution {
    pub fn convert_to_base7(num: i32) -> String {
        if num == 0 {
            return String::from("0");
        }
        let is_neg = num.is_negative();
        let mut res = vec![];
        let mut n = num;
        if is_neg {
            n = num.abs();
        }
        while n > 0 {
            let r = n % 7;
            n = n / 7;
            res.push(r)
        }
        let str = res.iter().map(|e| e.to_string()).rev().collect::<String>();
        if is_neg {
            (str.parse::<i32>().unwrap() * -1).to_string()
        } else {
            return str;
        }
    }
}

fn main() {
    assert!(vec![(100, "202"), (-7, "-10"), (0, "0"), (-8, "-11")]
        .into_iter()
        .all(|e| {
            let sol = <dyn Solution>::convert_to_base7(e.0);
            println!("got={},sol={}", sol, e.1);
            return sol == e.1;
        }));
}
