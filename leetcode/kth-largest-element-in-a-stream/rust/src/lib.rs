use std::collections::BinaryHeap;

struct KthLargest {
    k: i32,
    heap: BinaryHeap<i32>,
}

impl KthLargest {
    fn new(k: i32, nums: Vec<i32>) -> Self {
        let mut kthLargest = KthLargest {
            k: k,
            heap: BinaryHeap::with_capacity(k as usize + 1),
        };
        for n in nums {
            kthLargest.add(n);
        }
        kthLargest
    }

    fn add(&mut self, val: i32) -> i32 {
        self.heap.push(-val);
        if self.heap.len() > self.k as usize {
            self.heap.pop();
        }
        -self.heap.peek().unwrap()
    }
}