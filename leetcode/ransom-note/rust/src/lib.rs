#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::can_construct("a".to_string(), "b".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aa".to_string(), "aab".to_string()),
            true
        );
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut map: HashMap<char, i32> = HashMap::new();

        magazine
            .chars()
            .for_each(|c| *map.entry(c).or_insert(0) += 1);

        for c in ransom_note.chars() {
            if let Some(v) = map.get_mut(&c) {
                if *v == 0 { return false };
                *v -= 1;
            } else {
                return false;
            }
        }

        true
    }
}
