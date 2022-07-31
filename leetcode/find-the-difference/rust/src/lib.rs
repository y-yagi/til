#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_the_difference("abcd".to_string(), "abcde".to_string()), 'e');
        assert_eq!(Solution::find_the_difference("".to_string(), "y".to_string()), 'y');
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn find_the_difference(s: String, t: String) -> char {
        let mut map: HashMap<char, i32> = HashMap::new();

        for c in t.chars() {
            *map.entry(c).or_insert(0) += 1;
        }

        for c in s.chars() {
            *map.entry(c).or_insert(0) -= 1;
        }

        for v in map {
            if v.1 != 0 {
                return v.0;
            }
        }

        ' '
    }
}