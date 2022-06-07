#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::integer_break(2), 1);
        assert_eq!(Solution::integer_break(10), 36);
    }
}

struct Solution {}

impl Solution {
    pub fn integer_break(n: i32) -> i32 {
        if n == 2 {
            return 1;
        }

        if n == 3 {
            return 2;
        }

        let mut product = 1;
        let mut n = n;

        while n > 4 {
            product *= 3;
            n -= 3;
        }
        product *= n;

        product
    }
}
