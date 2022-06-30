#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::longest_palindrome("abccccdd".to_string()), 7);
        assert_eq!(Solution::longest_palindrome("1".to_string()), 1);
    }
}

struct Solution {}

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut counter = vec![0; 128];
        for b in s.as_bytes() {
            counter[(b - b'0') as usize] += 1;
        }

        let mut ans = 0;

        for v in counter {
            ans += v / 2 * 2;
            if ans % 2 == 0 && v % 2 == 1 {
                ans += 1;
            }
        }

        ans
    }
}
