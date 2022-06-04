#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::number_of_arithmetic_slices(vec![1, 2, 3, 4]), 3);
    }
}

struct Solution {}

impl Solution {
    pub fn number_of_arithmetic_slices(nums: Vec<i32>) -> i32 {
        if nums.len() < 3 { return 0 }

        let mut ans: i32 = 0;
        let mut dp = vec![0; nums.len()];

        if nums[2] - nums[1] == nums[1] - nums[0] {
            dp[2] = 1;
        }
        ans = dp[2] as i32;

        for i in 3..nums.len() {
            if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
                dp[i] = dp[i-1] + 1;
            }
            ans += dp[i] as i32;
        }

        ans
    }
}