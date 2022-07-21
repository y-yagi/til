#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::decode_string("3[a]2[bc]".to_string()),
            "aaabcbc".to_string()
        );
        assert_eq!(
            Solution::decode_string("3[a2[c]]".to_string()),
            "accaccacc".to_string()
        );
    }
}

struct Solution {}

impl Solution {
    pub fn decode_string(s: String) -> String {
        let mut stack = vec![];
        let mut num: u32 = 0;
        let mut ans = String::new();

        for c in s.chars() {
            if c == '[' {
                stack.push((num, ans));
                num = 0;
                ans = String::new();
            } else if c == ']' {
                let r = stack.pop().unwrap();
                ans = format!("{}{}", r.1, ans.repeat(r.0 as usize));
            } else if let Some(d) = c.to_digit(10) {
                num = num * 10 + d;
            } else {
                ans = format!("{}{}", ans, c);
            }
        }

        ans
    }
}
