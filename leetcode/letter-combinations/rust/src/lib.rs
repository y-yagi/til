#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::letter_combinations("2".to_string()), vec!["a".to_string(), "b".to_string(), "c".to_string()]);
    }
}

struct Solution {}

impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        if digits.len() == 0 {
            return vec![];
        }

        let digits = digits.as_bytes();
        let mut ans = vec![];
        Self::backtrace(digits, 0, &mut vec![], &mut ans);
        ans
    }

    fn backtrace(digits: &[u8], pos: usize, curr: &mut Vec<u8>, res: &mut Vec<String>) {
        if pos == digits.len() {
            res.push(String::from_utf8(curr.clone()).unwrap());
            return;
        }

        let letters = match digits[pos] {
            b'2' => [b'a', b'b', b'c'].iter(),
            b'3' => [b'd', b'e', b'f'].iter(),
            b'4' => [b'g', b'h', b'i'].iter(),
            b'5' => [b'j', b'k', b'l'].iter(),
            b'6' => [b'm', b'n', b'o'].iter(),
            b'7' => [b'p', b'q', b'r', b's'].iter(),
            b'8' => [b't', b'u', b'v'].iter(),
            b'9' => [b'w', b'x', b'y', b'z'].iter(),
            _ => panic!(),
        };

        for &letter in letters {
            curr.push(letter);
            Self::backtrace(digits, pos + 1, curr, res);
            curr.pop();
        }
    }
}