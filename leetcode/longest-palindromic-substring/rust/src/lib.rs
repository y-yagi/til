#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::longest_palindrome("babad".to_string()),
            "bad".to_string()
        );
        assert_eq!(
            Solution::longest_palindrome("cbbd".to_string()),
            "bb".to_string()
        );
    }
}

struct Solution {}

impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let bytes = s.as_bytes();
        let (mut start, mut end) = (0, 0);

        for i in 0..bytes.len() {
            let (l1, r1) = Self::expand_around_center(bytes, i, i);
            let (l2, r2) = Self::expand_around_center(bytes, i, i + 1);
            let len1 = (r1 - l1 + 1);
            let len2 = (r2 - l2 + 1);
            let cur_len = end - start + 1;

            if len1 >= len2 && len1 > cur_len {
                start = l1;
                end = r1;
            } else if len2 > len1 && len2 > cur_len {
                start = l2;
                end = r2;
            }
        }

        String::from_utf8(bytes[start..=end].to_vec()).unwrap()
    }

    fn expand_around_center(s: &[u8], left: usize, right: usize) -> (usize, usize) {
        let (mut l, mut r) = (left, right);
        let (mut start, mut end) = (0, 0);

        while r < s.len() && s[l] == s[r] {
            start = l;
            end = r;
            if l == 0 {
                break;
            }
            l -= 1;
            r += 1;
        }
        (start, end)
    }
}
