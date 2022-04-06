#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::subarray_sum(vec![1, 1, 1], 2), 2);
        assert_eq!(Solution::subarray_sum(vec![1, 2, 3], 3), 2);
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
        let mut count: i32 = 0;
        let mut sum: i32 = 0;
        let mut map: HashMap<i32, i32> = HashMap::new();
        map.insert(0, 1);

        for num in nums {
            sum += num;
            if let Some(v) = map.get(&(sum - k)) {
                count += v;
            }

            *map.entry(sum).or_insert(0) += 1;
        }

        count
    }
}