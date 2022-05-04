#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::fib(2), 1);
        assert_eq!(Solution::fib(3), 2);
        assert_eq!(Solution::fib(4), 3);
    }
}

struct Solution {}

impl Solution {
    pub fn fib(n: i32) -> i32 {
        let mut tup = (0, 1);
        match n {
            0 => 0,
            1 => 1,
            _ => {
                for i in 2..n+1 {
                    let t = tup.1;
                    tup.1 = tup.0 + tup.1;
                    tup.0 = t;
                }
                tup.1
            }
        }
    }
}