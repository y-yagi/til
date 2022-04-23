#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::unique_paths_with_obstacles(vec![vec![0, 0, 0], vec![0, 1, 0], vec![0, 0, 0]]),  2);
    }
}

struct Solution {}

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let m: usize = obstacle_grid.len();
        let n: usize = obstacle_grid[0].len();

        let mut dp: Vec<Vec<i32>> = vec![vec![0; n]; m];

        for i in 0..m {
            for j in 0..n {
                if obstacle_grid[i][j] == 1 {
                    dp[i][j] = 0;
                    continue
                }

                if i == 0 && j == 0 {
                    dp[i][j] = 1;
                    continue
                }

                if i == 0 {
                    dp[i][j] = dp[i][j-1];
                } else if j == 0 {
                    dp[i][j] = dp[i-1][j];
                } else {
                    dp[i][j] = dp[i-1][j] + dp[i][j-1];
                }
            }
        }

        dp[m-1][n-1]
    }
}
