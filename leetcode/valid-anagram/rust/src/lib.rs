#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_anagram("anagram".to_string(), "mnagraa".to_string()), true);
        assert_eq!(Solution::is_anagram("rat".to_string(), "car".to_string()), false);
    }
}

struct Solution {}

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() { return false }

        let mut counter = vec![0; 26];

        for c in s.chars() {
            counter[(c as u32 - 'a' as u32) as usize] += 1;
        }

        for c in t.chars() {
            counter[(c as u32 - 'a' as u32) as usize] -= 1;
        }

        for v in counter {
            if v != 0 { return false }
        }

        true
    }
}