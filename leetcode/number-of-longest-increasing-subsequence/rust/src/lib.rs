#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_number_of_lis(vec![1, 3, 5, 4, 7]), 2);
        assert_eq!(Solution::find_number_of_lis(vec![2, 2, 2, 2, 2]), 5);
    }
}

struct Solution {}

impl Solution {
    pub fn find_number_of_lis(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n <= 1 {
            return n as i32;
        }

        let mut lengths = vec![0; n];
        let mut counts = vec![1; n];

        for i in 0..n {
            for j in 0..i {
                if nums[j] < nums[i] {
                    if lengths[j] >= lengths[i] {
                        lengths[i] = lengths[j] + 1;
                        lengths[i] = counts[j]
                    } else if lengths[j] + 1 == lengths[i] {
                        counts[i] += counts[j]
                    }
                }
            }
        }

        let mut longest = 0;
        let mut ans = 0;

        for length in &lengths {
            longest = std::cmp::max(longest, *length);
        }

        for i in 0..n {
            if lengths[i] == longest {
                ans += counts[i];
            }
        }

        ans
    }
}