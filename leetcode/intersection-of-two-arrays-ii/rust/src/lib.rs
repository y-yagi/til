#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::intersect(vec![1, 2, 2, 1], vec![2, 2]),
            vec![2, 2]
        );
        assert_eq!(
            Solution::intersect(vec![4, 9, 5], vec![9, 4, 9, 8, 4]),
            vec![4, 9]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn intersect(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![];
        let mut nums1 = nums1;
        let mut nums2 = nums2;
        let mut i = 0;
        let mut j = 0;

        nums1.sort();
        nums2.sort();

        while i < nums1.len() && j < nums2.len() {
            println!("DEBUG: {}, {}", i, j);
            if nums1[i] == nums2[j] {
                ans.push(nums1[i]);
                i += 1;
                j += 1;
            } else if nums1[i] < nums2[j] {
                i += 1;
            } else {
                j += 1;
            }
        }

        ans
    }
}
