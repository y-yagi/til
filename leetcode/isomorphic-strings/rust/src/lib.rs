#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::is_isomorphic("foo".to_string(), "bar".to_string()),
            false
        );
        assert_eq!(
            Solution::is_isomorphic("egg".to_string(), "add".to_string()),
            true
        );
        assert_eq!(
            Solution::is_isomorphic("paper".to_string(), "title".to_string()),
            true
        );
        assert_eq!(
            Solution::is_isomorphic("badc".to_string(), "baba".to_string()),
            false
        );
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn is_isomorphic(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let s_bytes: Vec<_> = s.chars().collect();
        let t_bytes: Vec<_> = t.chars().collect();
        let mut dict1: HashMap<char, char> = HashMap::new();
        let mut dict2: HashMap<char, char> = HashMap::new();

        for i in 0..s.len() {
            if let Some(c) = dict1.get(&s_bytes[i]) {
                if *c != t_bytes[i] {
                    return false;
                }
            }
            if let Some(c) = dict2.get(&t_bytes[i]) {
                if *c != s_bytes[i] {
                    return false;
                }
            }
            dict1.insert(s_bytes[i], t_bytes[i]);
            dict2.insert(t_bytes[i], s_bytes[i]);
        }

        true
    }
}
