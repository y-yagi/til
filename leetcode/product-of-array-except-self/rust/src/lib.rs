#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::product_except_self(vec![1, 2, 3, 4]), vec![24, 12, 8, 6]);
    }
}

struct Solution {}

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![1; nums.len()];

        let mut left = 1;
        for i in 0..nums.len() {
            ans[i] *= left;
            left *= nums[i];
        }

        let mut right = 1;
        for i in (0..=nums.len()-1).rev() {
            ans[i] *= right;
            right *= nums[i];
        }

        ans
    }
}