#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::majority_element(vec![3, 3, 4]), 3);
        assert_eq!(Solution::majority_element(vec![3, 2, 3]), 3);
        assert_eq!(Solution::majority_element(vec![2, 2, 1, 1, 1, 2, 2]), 2);
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut ans = -1;

        for n in &nums {
            let counter = map.entry(*n).or_insert(0);
            *counter += 1;

            if *counter > (nums.len() / 2) as i32 {
                ans = *n
            }
        }

        ans
    }
}