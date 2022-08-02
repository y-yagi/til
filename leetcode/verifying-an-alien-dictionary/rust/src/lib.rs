#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::is_alien_sorted(
                vec!["hello".to_string(), "leetcode".to_string()],
                "hlabcdefgijkmnopqrstuvwxyz".to_string()
            ),
            true
        );
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn is_alien_sorted(words: Vec<String>, order: String) -> bool {
        let dict: HashMap<char, i32> = order
            .chars()
            .enumerate()
            .map(|(i, c)| (c, i as i32))
            .collect();

        for i in 0..words.len() - 1 {
            let curr_chars: Vec<char> = words[i].chars().collect();
            let next_chars: Vec<char> = words[i + 1].chars().collect();
            for i in 0..curr_chars.len() {
                if i == next_chars.len() || dict[&next_chars[i]] < dict[&curr_chars[i]] {
                    return false;
                }
                if dict[&curr_chars[i]] < dict[&next_chars[i]] {
                    break;
                }
            }
        }
        true
    }
}
