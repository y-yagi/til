#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::my_pow(2.0, 10), 1024.0);
    }
}

struct Solution {}

impl Solution {
    pub fn my_pow(x: f64, n: i32) -> f64 {
        if n == 0 { return 1.0; }
        if n < 0 { return 1.0 / Self::my_pow(x, -1) }
        if n % 2 == 0 { return x * Self::my_pow(x, n - 1) }

        return Self::my_pow(x*x, n / 2);
    }
}