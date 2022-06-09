#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_happy(19), true);
        assert_eq!(Solution::is_happy(2), false);
    }
}

struct Solution {}

use std::collections::HashSet;

impl Solution {
    pub fn is_happy(mut n: i32) -> bool {
        let mut cache = HashSet::new();

        while n != 1 {
            let mut current = n;
            let mut tmp = 0;
            while current != 0 {
                tmp += (current % 10).pow(2);
                current = current / 10;
            }

            n = tmp;
            if !cache.insert(n) {
                return false;
            }
        }

        return true;
    }
}