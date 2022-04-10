#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::word_break(String::from("leetcode"), vec![String::from("leet"), String::from("code")]), true);
        assert_eq!(Solution::word_break(String::from("applepenapple"), vec![String::from("apple"), String::from("pen")]), true);
    }
}

struct Solution { }

use std::collections::HashSet;

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let dict: HashSet<String> = word_dict.into_iter().collect();
        let mut dp: Vec<bool> = vec![false; s.len() + 1];
        dp[0] = true;

        for i in 1..=s.len() {
            for j in (0..=i-1).rev() {
                if dp[j] && dict.contains(&s[j..=i-1]) {
                    dp[i] = true;
                    break;
                }
            }
        }

        dp[s.len()]
    }
}