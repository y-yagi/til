#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::count_odds(3, 7), 3);
        assert_eq!(Solution::count_odds(8, 10), 1);
    }
}

struct Solution {}

impl Solution {
    pub fn count_odds(low: i32, high: i32) -> i32 {
        let nums = high - low + 1;
        if low % 2 != 0 && high %2 != 0 {
            return nums / 2 + 1
        }
        nums / 2
    }
}