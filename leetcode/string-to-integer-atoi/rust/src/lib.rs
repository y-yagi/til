#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_valid_test() {
        assert_eq!(Solution::my_atoi("42".to_string()), 42);
        assert_eq!(Solution::my_atoi("4193 with words".to_string()), 4193);
    }
}

struct Solution {}

impl Solution {
    pub fn my_atoi(str: String) -> i32 {
        let s = str.trim().to_string();
        let mut ans: i64 = 0;

        if s.len() == 0 {
            return 0;
        }

        let mut sign = 1;
        let mut start = 0;
        if s.chars().nth(0) == Some('-') {
            sign = -1;
            start = 1;
        }

        if s.chars().nth(0) == Some('+') {
            start = 1;
        }

        for c in s.chars().skip(start) {
            if !c.is_digit(10) {
                break;
            }

            ans = ans * 10 + c.to_digit(10).unwrap_or(0) as i64;
            if ans * sign > std::i32::MAX as i64 {
                return std::i32::MAX;
            }

            if ans * sign < std::i32::MIN as i64 {
                return std::i32::MIN;
            }
        }

        return (ans * sign) as i32;
    }
}