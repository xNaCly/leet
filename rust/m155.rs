struct MinStack {
    c: Vec<i32>,
    t: i32,
}

impl MinStack {
    fn new() -> Self {
        MinStack { c: vec![], t: 0 }
    }

    fn push(&mut self, val: i32) {
        self.t = val;
        self.c.push(val);
    }

    fn pop(&mut self) {
        self.c.pop();
        self.t = self.c.last().unwrap_or(&0).clone();
    }

    fn top(&self) -> i32 {
        self.t
    }

    fn get_min(&self) -> i32 {
        let mut m: i32 = *(self.c.get(0).unwrap());
        for i in self.c.to_vec() {
            if i < m {
                m = i;
            }
        }
        return m;
    }
}

fn main() {
    let mut m = MinStack::new();
    m.push(-2);
    m.push(0);
    m.push(-3);
    assert_eq!(m.get_min(), -3); // return -3
    m.pop();
    assert_eq!(m.top(), 0); // return 0
    assert_eq!(m.get_min(), -2); // return 0
}
