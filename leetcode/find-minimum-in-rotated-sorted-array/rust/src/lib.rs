#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_min(vec![3, 4, 5, 1, 2]), 1);
        assert_eq!(Solution::find_min(vec![4, 5, 6, 7, 0, 1, 2]), 0);
        assert_eq!(Solution::find_min(vec![11, 13, 15, 17]), 11);
    }
}

struct Solution {}

impl Solution {
    pub fn find_min(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 { return nums[0] }

        let mut left = 0;
        let mut right = nums.len() - 1;

        if nums[right] > nums[left] { return nums[left ]}

        while right > left {
            let mid = left + (right - left) / 2;

            if nums[mid] > nums[mid+1] { return nums[mid+1]; }
            if nums[mid-1] > nums[mid] { return nums[mid]; }

            if nums[mid] > nums[0] {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        -1
    }
}