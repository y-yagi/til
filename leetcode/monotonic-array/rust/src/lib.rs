#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_monotonic(vec![1, 2, 2, 3]), true);
        assert_eq!(Solution::is_monotonic(vec![1, 3, 2]), false);
    }
}

struct Solution {}

impl Solution {
    pub fn is_monotonic(nums: Vec<i32>) -> bool {
        let mut increasing = true;
        let mut decreasing = true;

        for i in 0..nums.len() - 1 {
            if nums[i] > nums[i+1] {
                increasing = false;
            }
            if nums[i] < nums[i+1] {
                decreasing = false;
            }
        }

        increasing || decreasing
    }
}