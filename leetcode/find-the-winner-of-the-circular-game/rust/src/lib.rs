#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_the_winner(5, 2), 3);
        assert_eq!(Solution::find_the_winner(6, 5), 1);
    }
}

struct Solution {}

impl Solution {
    pub fn find_the_winner(n: i32, k: i32) -> i32 {
        if n == 1 {
            return 1
        }

        (k + Self::find_the_winner(n-1, k) - 1) % n + 1
    }
}