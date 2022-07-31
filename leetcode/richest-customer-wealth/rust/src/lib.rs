#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::maximum_wealth(vec![vec![1, 2, 3], vec![3, 2, 1]]), 6);
    }
}

struct Solution {}

impl Solution {
    pub fn maximum_wealth(accounts: Vec<Vec<i32>>) -> i32 {
        let mut ans = i32::MIN;
        for account in accounts {
            ans = std::cmp::max(ans, account.iter().sum());
        }
        ans
    }
}