#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::largest_perimeter(vec![2, 1, 2]), 5);
        assert_eq!(Solution::largest_perimeter(vec![1, 2, 1]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn largest_perimeter(nums: Vec<i32>) -> i32 {
        let mut nums = nums;
        nums.sort();
        nums.reverse();

        for i in 1..nums.len() - 1{
            if nums[i-1] + nums[i] > nums[i+1] && nums[i-1] + nums[i+1] > nums[i] && nums[i] + nums[i+1] > nums[i-1] {
                return nums[i-1] + nums[i] + nums[i+1];
            }
        }

        0
    }
}