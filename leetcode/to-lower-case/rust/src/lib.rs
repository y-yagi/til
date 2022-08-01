#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::to_lower_case(("Hello").to_string()), "hello".to_string());
    }
}

struct Solution {}

impl Solution {
    pub fn to_lower_case(s: String) -> String {
        let mut chars: Vec<char> = s.chars().collect();

        for i in 0..chars.len() {
            if chars[i] >= 'A' && chars[i] <= 'Z' {
                chars[i] = ((chars[i] as u8) + 32) as char
            }
        }

        chars.into_iter().collect()
    }
}