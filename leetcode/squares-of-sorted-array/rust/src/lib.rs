#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::sorted_squares(vec![-4, -1, 0, 3, 10]), vec![0, 1, 9, 16, 100]);
        assert_eq!(Solution::sorted_squares(vec![-7, -3, 2, 3, 11]), vec![4, 9, 9, 49, 121]);
    }
}

struct Solution {}

impl Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![0; nums.len()];
        let mut i = 0;
        let mut j = nums.len() - 1;
        let mut p: i32 = (nums.len() - 1) as i32;

        while p >= 0 {
            if nums[i].abs() > nums[j].abs() {
                ans[p as usize] = nums[i] * nums[i];
                i += 1;
            } else {
                ans[p as usize] = nums[j] * nums[j];
                j -= 1;
            }
            p -= 1;
        }

        ans
    }
}