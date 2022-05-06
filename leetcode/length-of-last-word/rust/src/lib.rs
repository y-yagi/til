#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::length_of_last_word("Hello World".to_string()), 5);
        assert_eq!(Solution::length_of_last_word("   fly me   to   the moon ".to_string()), 4);
    }
}

struct Solution {}

impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let mut ans = 0;
        let mut i = (s.len() - 1) as i32;

        while i >= 0 && s.as_bytes()[i as usize] as char == ' ' {
            i -= 1;
        }

        while i >= 0 && s.as_bytes()[i as usize] as char != ' ' {
            ans += 1;
            i -= 1;
        }

        ans
    }
}