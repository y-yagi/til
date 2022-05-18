#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_power_of_two(1), true);
        assert_eq!(Solution::is_power_of_two(3), false);
    }
}

struct Solution {}

impl Solution {
    pub fn is_power_of_two(n: i32) -> bool {
        n > 0 && (n & n - 1) == 0
    }
}