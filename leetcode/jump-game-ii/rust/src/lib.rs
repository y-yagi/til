#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::jump(vec![2, 3, 1, 1, 4]), 2);
        assert_eq!(Solution::jump(vec![3, 2, 0, 1, 4]), 2);
    }
}

struct Solution {}

impl Solution {
    pub fn jump(nums: Vec<i32>) -> i32 {
        let mut jumps = 0;
        let mut cur_end = 0;
        let mut cur_farthest = 0;

        for i in 0..nums.len() - 1 {
            cur_farthest = std::cmp::max(cur_farthest, i + nums[i] as usize);
            if i == cur_end {
                jumps += 1;
                cur_end = cur_farthest
            }
        }

        jumps
    }
}
