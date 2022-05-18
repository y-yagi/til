#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::minimum_total(
            vec![vec![2], vec![3, 4], vec![6, 5, 7], vec![4, 1, 8, 3]]),
            11
        );
        assert_eq!(Solution::minimum_total(vec![vec![-10]]), -10);
    }
}

struct Solution {}

impl Solution {
    pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
        let mut dp = triangle[triangle.len() - 1].to_vec();
        for i in (0..triangle.len() - 1).rev() {
            for j in 0..i + 1 {
                dp[j] = triangle[i][j] + std::cmp::min(dp[j], dp[j + 1])
            }
        }
        dp[0]
    }
}