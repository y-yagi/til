#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::rob(vec![1, 2, 3, 1]), 4);
    }
}

struct Solution { }

use std::cmp::max;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut prev_max = 0;
        let mut cur_max = 0;

        for num in nums {
            let t = cur_max;
            cur_max = max(prev_max + num, cur_max);
            prev_max = t;
        }

        cur_max
    }
}