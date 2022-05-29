#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::combination_sum2(vec![10, 1, 2, 7, 6, 1, 5], 8),
            vec![vec![1, 1, 6], vec![1, 2, 5], vec![1, 7], vec![2, 6]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn combination_sum2(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut candidates = candidates;
        candidates.sort();
        let mut ans = Vec::new();
        let mut curr = Vec::new();

        Self::backtrace(0, target, &candidates, &mut curr, &mut ans);
        ans
    }


    fn backtrace(i: usize, t: i32, candidates: &Vec<i32>, curr: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>) {
        if t == 0 {
            ans.push(curr.clone());
            return;
        }

        let start = i;
        let mut i = i;

        while i < candidates.len() && t >= candidates[i] {
            if i == start || candidates[i] != candidates[i - 1] {
                curr.push(candidates[i]);
                Self::backtrace(i + 1, t - candidates[i], candidates, curr, ans);
                curr.pop();
            }
            i += 1;
        }
    }
}