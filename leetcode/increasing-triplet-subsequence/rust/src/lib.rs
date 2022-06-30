#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::increasing_triplet(vec![1, 2, 3, 4, 5]), true);
        assert_eq!(Solution::increasing_triplet(vec![2, 1, 5, 0, 4, 6]), true);
    }
}

struct Solution {}

impl Solution {
    pub fn increasing_triplet(nums: Vec<i32>) -> bool {
        let mut small = i32::MAX;
        let mut big = i32::MAX;

        for n in nums {
            if n <= small {
                small = n;
            } else if n <= big {
                big = n
            } else {
                return true
            }
        }

        false
    }
}