#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::str_str("hello".to_string(), "ll".to_string()), 2);
        assert_eq!(Solution::str_str("aaaaa".to_string(), "bba".to_string()), -1);
    }
}

struct Solution {}

impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        let (hlen, nlen) = (haystack.len(), needle.len());

        if nlen == 0 {
            return 0;
        }

        for i in 0..=(hlen - nlen) {
            if &haystack[i..i+nlen] == needle {
                return i as i32;
            }
        }

        -1
    }
}