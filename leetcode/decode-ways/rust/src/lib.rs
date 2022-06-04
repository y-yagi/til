#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::num_decodings("12".to_string()), 2);
        assert_eq!(Solution::num_decodings("226".to_string()), 3);
        assert_eq!(Solution::num_decodings("06".to_string()), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        if s.len() == 0 {
            return 0;
        }

        let mut dp = vec![0; s.len()+1];
        dp[0] = 1;

        let bytes = s.as_bytes();

        if bytes[0] == b'0' {
            dp[1] = 0;
        } else {
            dp[1] = 1;
        }

        for i in 2..dp.len() {
            if bytes[i-1] != b'0' {
                dp[i] += dp[i-1];
            }
            let sum = (bytes[i-2] - b'0') * 10 + bytes[i-1] - b'0';
            if sum >= 10 && sum <= 26 {
                dp[i] += dp[i-2];
            }
        }

        dp[s.len()]
    }
}