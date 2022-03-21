#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::subsets(vec![1, 2, 3]),
            vec![
                vec![], vec![1], vec![2], vec![1, 2],
                vec![3], vec![1, 3], vec![2, 3], vec![1, 2, 3],
            ]
        );
        assert_eq!(Solution::subsets(vec![]), vec![vec![], vec![0]]);
    }
}

struct Solution { }

impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = vec![];
        Self::backtrace(&mut ans, &nums[..], &mut vec![], 0);
        ans
    }

    pub fn backtrace(ans: &mut Vec<Vec<i32>>, nums: &[i32] , cur: &mut Vec<i32>, begin: usize) {
        ans.push(cur.clone());

        for i in begin..nums.len() {
            cur.push(nums[i]);
            Self::backtrace(ans, nums, cur, i + 1);
            cur.pop();
        }
    }
}