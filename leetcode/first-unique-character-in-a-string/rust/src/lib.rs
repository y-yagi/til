#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::first_uniq_char(String::from("leetcode")), 0);
        assert_eq!(Solution::first_uniq_char(String::from("loveleetcode")), 2);
        assert_eq!(Solution::first_uniq_char(String::from("aabb")), -1);
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn first_uniq_char(s: String) -> i32 {
        let mut map: HashMap<char, i32> = HashMap::new();

        s.chars().for_each(|c| *map.entry(c).or_insert(0) += 1);
        for (i, c) in s.chars().enumerate() {
            let v = *map.get(&c).unwrap();
            if v == 1 { return i as i32 }
        }

        -1
    }
}