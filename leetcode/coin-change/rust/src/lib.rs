#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::coin_change(vec![1, 2, 5], 11), 3);
        assert_eq!(Solution::coin_change(vec![2], 3), -1);
    }
}

struct Solution { }

use std::cmp;

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let max = amount + 1;
        let mut dp = vec![max; amount as usize + 1];
        dp[0] = 0;

        for i in 0..=(amount as usize) {
            for j in 0..coins.len() {
                if coins[j] as usize <= i {
                    dp[i] = cmp::min(dp[i], dp[i-coins[j] as usize] + 1);
                }
            }
        }

        if dp[amount as usize] > amount {
            return -1;
        }

        dp[amount as usize]
    }
}