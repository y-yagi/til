struct MyQueue {
    sin: Vec<i32>,
    sout: Vec<i32>,
}

impl MyQueue {
    fn new() -> Self {
        Self {
            sin: vec![],
            sout: vec![],
        }
    }

    fn push(&mut self, x: i32) {
        self.sin.push(x);
    }

    fn pop(&mut self) -> i32 {
        if self.sout.is_empty() {
            while !self.sin.is_empty() {
                self.sout.push(self.sin.pop().unwrap());
            }
        }
        self.sout.pop().unwrap()
    }

    fn peek(&mut self) -> i32 {
        if self.sout.is_empty() {
            while !self.sin.is_empty() {
                self.sout.push(self.sin.pop().unwrap());
            }
        }
        *self.sout.last().unwrap()
    }

    fn empty(&self) -> bool {
        self.sin.is_empty() && self.sout.is_empty()
    }
}
