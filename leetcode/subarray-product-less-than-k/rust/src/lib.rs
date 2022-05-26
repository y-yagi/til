#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::num_subarray_product_less_than_k(vec![10, 5, 2, 6], 100), 8);
    }
}

struct Solution {}

impl Solution {
    pub fn num_subarray_product_less_than_k(nums: Vec<i32>, k: i32) -> i32 {
        if k <= 1 { return 0 };

        let mut ans = 0;
        let mut prod = 1;
        let mut left = 0;

        for right in 0..nums.len() {
            prod *= nums[right];
            while prod >= k  {
                prod /= nums[left];
                left += 1
            }
            ans += right - left + 1;
        }

        ans as i32
    }
}