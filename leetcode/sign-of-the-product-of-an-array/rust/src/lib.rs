#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::array_sign(vec![-1, -2, -3, -4, 3, 2, 1]), 1);
        assert_eq!(Solution::array_sign(vec![-1, 5, 0, 2, -3]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn array_sign(nums: Vec<i32>) -> i32 {
        let mut ans = 1;
        for num in nums {
            if num == 0 {
                return 0;
            }
            if num < 0 {
                ans *= -1;
            }
        }

        ans
    }
}
