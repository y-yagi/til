#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::combination_sum(vec![2, 3, 6, 7], 7),
            vec![vec![2, 2, 3], vec![7]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ans = Vec::new();
        let mut vec = Vec::new();
        Self::backtrace(&candidates, target, vec, &mut ans, 0);

        ans
    }

    fn backtrace(
        candidates: &Vec<i32>,
        target: i32,
        curr: Vec<i32>,
        ans: &mut Vec<Vec<i32>>,
        start_idex: usize,
    ) {
        for i in start_idex..candidates.len() {
            let item = candidates[i];
            if target - item < 0 {
                continue;
            }

            let mut vec = curr.clone();
            vec.push(item);
            if target == item {
                ans.push(vec);
            } else {
                Self::backtrace(candidates, target - item, vec, ans, i);
            }
        }
    }
}
