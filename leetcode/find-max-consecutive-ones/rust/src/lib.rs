#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_valid_test() {
        assert_eq!(Solution::find_max_consecutive_ones(vec![1, 0, 0, 1, 1, 0, 1]), 2);
        assert_eq!(Solution::find_max_consecutive_ones(vec![1, 1, 0, 1, 1, 1]), 3);
    }
}

struct Solution {}

use std::cmp::max;

impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let (mut ans, mut t) = (0, 0);

        for num in nums {
            if num == 0 {
                ans = max(ans, t);
                t = 0;
            } else {
                t += 1;
            }
        }

        max(ans, t)
    }
}