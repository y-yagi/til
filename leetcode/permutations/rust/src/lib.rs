#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::permute(vec![0, 1]), vec![vec![0, 1], vec![1, 0]]);
    }
}

struct Solution { }

impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result: Vec<Vec<i32>> = vec![];
        Self::backtrace(&mut result, nums, 0);
        result
    }

    pub fn backtrace(result: &mut Vec<Vec<i32>>, nums: Vec<i32>, begin: usize) {
        if begin >= nums.len() {
            result.push(nums);
            return
        }

        for i in begin..nums.len() {
            let mut nums = nums.clone();
            nums.swap(begin, i);
            Self::backtrace(result, nums, begin + 1);
        }
    }
}