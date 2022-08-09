#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::plus_one(vec![1, 2, 4]), vec![1, 2, 5]);
        assert_eq!(Solution::plus_one(vec![9]), vec![1, 0]);
    }
}

struct Solution {}

impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut ans = digits.clone();

        for i in (0..digits.len()).rev() {
            if digits[i] != 9 {
                ans[i] = ans[i] + 1;
                return ans;
            }
            ans[i] = 0;
        }

        [&vec![1], &ans[..]].concat()
    }
}