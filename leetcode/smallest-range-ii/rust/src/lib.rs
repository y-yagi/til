#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::smallest_range_ii(vec![1], 0), 0);
        assert_eq!(Solution::smallest_range_ii(vec![0, 10], 2), 6);
        assert_eq!(Solution::smallest_range_ii(vec![1, 3, 6], 3), 3);
    }
}

struct Solution {}

impl Solution {
    pub fn smallest_range_ii(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        nums.sort();
        let mut ans = nums[nums.len() - 1] - nums[0];

        for i in 0..nums.len() - 1 {
            let high = std::cmp::max(nums[nums.len() - 1] - k, nums[i] + k);
            let low = std::cmp::min(nums[0] + k, nums[i + 1] - k);
            ans = std::cmp::min(ans, high - low);
        }

        ans
    }
}
