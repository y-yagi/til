#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check_if_exist(vec![10, 2, 5, 3]),  true);
        assert_eq!(Solution::check_if_exist(vec![3, 1, 7, 11]),  false);
        assert_eq!(Solution::check_if_exist(vec![7, 1, 14, 11]),  true);
    }
}

struct Solution {}

use std::collections::HashSet;

impl Solution {
    pub fn check_if_exist(arr: Vec<i32>) -> bool {
        let set : HashSet<_> = arr.iter().cloned().collect();
        let mut has_zero = false;

        for v in arr {
            if v == 0 {
                if has_zero {
                    return true;
                }
                has_zero = true;
            } else {
              if set.contains(&(v*2)) {
                  return true;
              }
            }
        }

        false
    }
}