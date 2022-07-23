#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::character_replacement("ABBB".to_string(), 2), 4);
        assert_eq!(Solution::character_replacement("ABAB".to_string(), 2), 4);
        assert_eq!(Solution::character_replacement("AABABBA".to_string(), 1), 4);
    }
}

struct Solution {}

use std::cmp::max;

impl Solution {
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let mut ans = 0;
        let mut left = 0;
        let mut max_freq_count = 0;
        let mut freqs = [0; 26];
        let chars: Vec<char> = s.chars().collect();

        for right in 0..chars.len() {
            freqs[(chars[right] as u32 - 'A' as u32) as usize] += 1;
            max_freq_count = max(max_freq_count, freqs[(chars[right] as u32 - 'A' as u32) as usize]);

            while right - left + 1 - max_freq_count > k as usize {
                left += 1;
            }

            ans = max(ans, right - left + 1);
        }

        ans as i32
    }
}
