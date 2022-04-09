#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_profit(vec![7, 1, 5, 3, 6, 4]), 5);
        assert_eq!(Solution::max_profit(vec![7, 6, 4, 3, 1]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.len() < 2 {
            return 0
        }
        let mut ans = 0;
        let mut min = i32::MAX;

        for price in prices {
            if price < min {
                min = price
            } else if price - min > ans  {
                ans = price - min
            }
        }

        ans
    }
}