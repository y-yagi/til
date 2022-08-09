#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::repeated_substring_pattern("abab".to_string()),
            true
        );
        assert_eq!(
            Solution::repeated_substring_pattern("abcabcabcabc".to_string()),
            true
        );
    }
}

struct Solution {}

impl Solution {
    pub fn repeated_substring_pattern(s: String) -> bool {
        for i in 1..(s.len() / 2 + 1) {
            if s.len() % i != 0 {
                continue;
            }
            let substring = (&s[..i]).clone();
            if s.eq(&substring.repeat(s.len() / i)) {
                return true;
            }
        }
        false
    }
}
