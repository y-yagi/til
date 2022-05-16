#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::combine(1, 1), vec![vec![1]]);
    }
}

struct Solution {}

impl Solution {
    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        let mut ans = vec![];
        Self::backtrace(n, k, &mut vec![], 1, &mut ans);
        ans
    }

    fn backtrace(n: i32, k: i32, curr: &mut Vec<i32>, v: i32, ans: &mut Vec<Vec<i32>>) {
        if curr.len() == k as usize {
            ans.push(curr.clone());
            return;
        }

        for i in v..=n {
            curr.push(i);
            Self::backtrace(n, k, curr, i + 1, ans);
            curr.pop();
        }
    }
}