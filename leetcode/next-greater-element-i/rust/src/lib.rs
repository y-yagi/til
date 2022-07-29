#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::next_greater_element(vec![4, 1, 2], vec![1, 3, 4, 2]), vec![-1, 3, -1]);
        assert_eq!(Solution::next_greater_element(vec![2, 4], vec![1, 2, 3, 4]), vec![3, -1]);
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut map: HashMap<i32, usize> = HashMap::new();
        for i in 0..nums2.len() {
            map.insert(nums2[i],i);
        }

        let mut ans = vec![];

        for num in nums1 {
            let nums2_index = map.get(&num).unwrap();
            let mut x = -1;
            for i in *nums2_index..nums2.len() {
                if nums2[i] > num {
                    x = nums2[i];
                    break;
                }
            }
            ans.push(x);
        }

        ans
    }
}