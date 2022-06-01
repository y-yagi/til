#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::can_jump(vec![2, 0, 0]), true);
        assert_eq!(Solution::can_jump(vec![2, 3, 1, 1, 4]), true);
        assert_eq!(Solution::can_jump(vec![3, 2, 1, 0, 4]), false);
    }
}

struct Solution {}

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut m = 0;
        for (i, val) in nums.iter().enumerate() {
            if i > m {
                return false;
            }

            let val = *val as usize;
            m = std::cmp::max(m, i + val);
        }
        return true;
    }
}
