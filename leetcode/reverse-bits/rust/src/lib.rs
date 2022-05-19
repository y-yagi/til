#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

struct Solution {}

impl Solution {
    pub fn reverse_bits(x: u32) -> u32 {
        let mut ans: u32 = 0u32;
        let mut x = x.clone();

        for _ in 0..32 {
            ans = (ans << 1) | (x & 1);
            x >>= 1;
        }

        ans
    }
}