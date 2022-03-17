#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(length_of_longest_substring(String::from("abcabcbb")), 3);
        assert_eq!(length_of_longest_substring(String::from("bbbbb")), 1);
        assert_eq!(length_of_longest_substring(String::from("pwwkew")), 3);
    }
}

use std::collections::HashMap;
use std::cmp::max;

pub fn length_of_longest_substring(s: String) -> i32 {
    let mut m : HashMap<char, i32> = HashMap::new();
    let mut ans = 0;
    let mut before = -1;
    let mut current = 0;

    for c in s.chars() {
        if let Some(last) = m.insert(c, current) {
            before = max(before, last);
        }
        ans = max(ans, current - before) ;
        current += 1;
    }

    ans
}