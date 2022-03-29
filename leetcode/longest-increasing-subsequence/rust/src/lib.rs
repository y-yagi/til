#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::length_of_lis(vec![0]), 1);
        assert_eq!(Solution::length_of_lis(vec![0,1,0,3,2,3]), 4);
        assert_eq!(Solution::length_of_lis(vec![10,9,2,5,3,7,101,18]), 4);
    }
}

struct Solution {}

use std::cmp;

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut ans = 1;
        let mut dp = vec![1; nums.len()];

        for i in 1..nums.len() {
            let mut maxval  = 0;
            for j in 0..i {
                if nums[i] > nums[j] {
                    maxval = cmp::max(maxval, dp[j])
                }
            }

            dp[i] = maxval + 1;
            ans = cmp::max(ans, dp[i]);
        }

        ans
    }
}