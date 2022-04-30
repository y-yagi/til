#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::min_sub_array_len(7, vec![2, 3, 1, 2, 4, 3]), 2);
        assert_eq!(Solution::min_sub_array_len(4, vec![1, 4, 4]), 1);
        assert_eq!(Solution::min_sub_array_len(11, vec![1, 1, 1, 1]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        let mut ans = std::i32::MAX;
        let mut left = 0;
        let mut sum = 0;

        for i in 0..nums.len() {
            sum += nums[i];
            while sum >= target {
                ans = std::cmp::min(ans, (i+1-left) as i32);
                sum -= nums[left];
                left += 1;
            }
        }

        if ans == std::i32::MAX { ans = 0 }
        ans
    }
}