#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::last_stone_weight(vec![2, 7, 4, 1, 8, 1]), 1);
        assert_eq!(Solution::last_stone_weight(vec![1]), 1);
    }
}

struct Solution {}

use std::collections::BinaryHeap;

impl Solution {
    pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
        let mut heap = BinaryHeap::<i32>::new();

        for stone in stones {
            heap.push(stone);
        }

        while heap.len() > 1 {
            let x = heap.pop().unwrap();
            let y = heap.pop().unwrap();

            if x != y {
                heap.push(x - y);
            }
        }

        if heap.len() == 1 {
            return heap.pop().unwrap();
        }

        0
    }
}