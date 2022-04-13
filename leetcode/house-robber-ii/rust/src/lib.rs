#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::rob(vec![2, 3, 2]), 3);
        assert_eq!(Solution::rob(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob(vec![1, 2, 3]), 3);
    }
}

struct Solution { }

use std::cmp::max;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n == 0 {
            return 0;
        }
        if n == 1 {
            return nums[0];
        }

        max(Self::robber(&nums, 0, n-2), Self::robber(&nums, 1, n-1))
    }

    fn robber(nums: &Vec<i32>, l: usize, r: usize) -> i32 {
        let mut pre = 0;
        let mut cur = 0;

        for i in l..r+1 {
            let tmp = max(pre+nums[i], cur);
            pre = cur;
            cur = tmp;
        }

        cur
    }
}
