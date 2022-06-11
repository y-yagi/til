#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        Solution::merge(&mut vec![1, 2, 3, 0, 0, 0], 3, &mut vec![2, 5, 6], 3);
    }
}

struct Solution {}

impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut i = m - 1;
        let mut j = n - 1;
        let mut pos = (m + n - 1) as usize;

        while i >= 0 && j >= 0 {
            if nums1[i as usize] > nums2[j as usize] {
                nums1[pos] = nums1[i as usize];
                i -= 1;
            } else {
                nums1[pos] = nums2[j as usize];
                j -= 1;
            }
            pos -= 1;
        }

        while j >= 0  {
            nums1[pos] = nums2[j as usize];
            pos -= 1;
            j -= 1;
        }
    }
}