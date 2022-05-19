#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::hammingWeight(1011), 3);
        assert_eq!(Solution::hammingWeight(10000000), 1);
    }
}

struct Solution {}

impl Solution {
    pub fn hammingWeight (n: u32) -> i32 {
        let (mut n, mut ans) = (n, 0);
        while n > 0 {
            if n & 1 == 1 {
                ans += 1;
            }
            n >>= 1;
        }

        ans
    }
}