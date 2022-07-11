#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2), 5);
        assert_eq!(
            Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
            4
        );
    }
}

struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        Self::select(nums, k as usize)
    }

    fn select(nums: Vec<i32>, k: usize) -> i32 {
        let mut left = vec![];
        let mut right = vec![];
        let mut mid = vec![];
        let pivot = nums[nums.len() / 2];

        for num in nums {
            if num > pivot {
                left.push(num);
            } else if num < pivot {
                right.push(num);
            } else {
                mid.push(num);
            }
        }

        if k <= left.len() {
            return Self::select(left, k);
        } else if k > left.len() + mid.len() {
            return Self::select(right, k - left.len() - mid.len());
        } else {
            return mid[0];
        }
    }
}
