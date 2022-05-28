#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::subsets_with_dup(vec![0]), vec![vec![], vec![0]]);
        assert_eq!(
            Solution::subsets_with_dup(vec![1, 2, 2]),
            vec![
                vec![],
                vec![1],
                vec![1, 2],
                vec![1, 2, 2],
                vec![2],
                vec![2, 2]
            ]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = vec![];
        let mut nums = nums;
        nums.sort();
        Self::backtrace(&mut ans, &nums[..], &mut vec![], 0);
        ans
    }

    pub fn backtrace(ans: &mut Vec<Vec<i32>>, nums: &[i32] , cur: &mut Vec<i32>, begin: usize) {
        ans.push(cur.clone());

        for i in begin..nums.len() {
            if i == begin || nums[i] != nums[i - 1] {
              cur.push(nums[i]);
              Self::backtrace(ans, nums, cur, i + 1);
              cur.pop();
            }
        }
    }
}
