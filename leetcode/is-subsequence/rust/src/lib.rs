#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_subsequence(String::from("abc"), String::from("ahbgdc")), true);
    }
}

struct Solution {}

impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        if s.len() == 0 { return true }

        let mut index_for_s = 0;
        for index_for_t in 0..t.len() {
            if t.as_bytes()[index_for_t] == s.as_bytes()[index_for_s] {
                index_for_s += 1;
            }
            if index_for_s == s.len() {
                return true
            }
        }

        false
    }
}