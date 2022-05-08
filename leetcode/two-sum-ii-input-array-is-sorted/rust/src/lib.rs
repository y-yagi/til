#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![2, 3, 4], 6), vec![1, 3]);
        assert_eq!(Solution::two_sum(vec![-1, 0], -1), vec![1, 2]);
    }
}

struct Solution {}

impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut left: i32 = 0;
        let mut right: i32 = (numbers.len() - 1) as i32;

        while left < right {
            let v = numbers[left as usize] + numbers[right as usize];
            if v == target {
                return vec![left + 1, right + 1];
            } else if v < target {
                left += 1;
            } else {
                right -= 1;
            }
        }

        vec![]
    }
}