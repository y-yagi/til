#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::backspace_compare("ab#c".to_string(), "ad#c".to_string()),
            true
        );
    }
}

struct Solution {}

impl Solution {
    pub fn backspace_compare(s: String, t: String) -> bool {
        Self::backspace(s) == Self::backspace(t)
    }

    fn  backspace(s: String) -> Vec<char> {
        let mut stack: Vec<char> = vec![];
        for c in s.chars() {
            if c == '#' {
                stack.pop();
            } else {
                stack.push(c);
            }
        }
        stack
    }
}
