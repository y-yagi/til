struct MyCircularQueue {
    queue: Vec<i32>,
    capacity: usize,
    len: usize,
    head: usize,
    tail: usize,
}

impl MyCircularQueue {
    fn new(k: i32) -> Self {
        Self {
            queue: vec![0; k as usize],
            capacity: k as usize,
            len: 0,
            head: 0,
            tail: 0,
        }
    }

    fn en_queue(&mut self, value: i32) -> bool {
        if self.is_full() {
            return false;
        }

        if self.len > 0 {
            self.tail = (self.tail + 1) % self.capacity;
        }

        self.len += 1;
        self.queue[self.tail] = value;
        true
    }

    fn de_queue(&mut self) -> bool {
        if self.is_empty() {
            return false;
        }
        self.len -= 1;
        if self.len == 0 {
            self.head = 0;
            self.tail = 0;
        } else {
            self.head = (self.head + 1) % self.capacity;
        }
        true
    }

    fn front(&self) -> i32 {
        if self.is_empty() {
            -1
        } else {
            self.queue[self.head]
        }
    }

    fn rear(&self) -> i32 {
        if self.is_empty() {
            -1
        } else {
            self.queue[self.tail]
        }
    }

    fn is_empty(&self) -> bool {
        return self.len == 0;
    }

    fn is_full(&self) -> bool {
        return self.len == self.capacity;
    }
}
