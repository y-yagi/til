#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::word_pattern("abba".to_string(), "dog dog dog dog".to_string()),
            false
        );
        assert_eq!(
            Solution::word_pattern("abba".to_string(), "dog cat cat dog".to_string()),
            true
        );
        assert_eq!(
            Solution::word_pattern("abba".to_string(), "dog cat cat fish".to_string()),
            false
        );
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn word_pattern(pattern: String, s: String) -> bool {
        let mut dict = HashMap::new();
        let mut rdict = HashMap::new();
        let words: Vec<&str> = s.split(" ").collect();

        if pattern.len() != words.len() {
            return false;
        }

        for (i, c) in pattern.chars().enumerate() {
            if let Some(word) = dict.get(&c) {
                if word != &words[i] {
                    return false;
                }
            } else {
                if let Some(v) = rdict.get(words[i]) {
                    if v != &c {
                        return false;
                    }
                }
                dict.insert(c, words[i].clone());
                rdict.insert(words[i].clone(), c);
            }
        }
        return true;
    }
}
