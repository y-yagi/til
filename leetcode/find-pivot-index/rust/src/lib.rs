#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::pivot_index(vec![1, 7, 3, 6, 5, 6]), 3);
        assert_eq!(Solution::pivot_index(vec![1, 2, 3]), -1);
        assert_eq!(Solution::pivot_index(vec![2, 1, -1]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        let (mut sum, mut leftsum) = (0, 0);

        for num in &nums {
            sum += num;
        }

        for i in 0..nums.len() {
            if leftsum == sum - leftsum - nums[i] {
                return i as i32;
            }
            leftsum += nums[i];
        }

        -1
    }
}
