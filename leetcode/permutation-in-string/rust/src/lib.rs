#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check_inclusion("ab".to_string(), "eidbaooo".to_string()), true);
        assert_eq!(Solution::check_inclusion("ab".to_string(), "eidboaoo".to_string()), false);
    }
}

struct Solution {}

impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        let u8_to_idx = |c: u8| { (c - b'a') as usize };

        let mut s1_freq  = [0u8; 26];
        s1.bytes().for_each(|b| s1_freq[u8_to_idx(b)] += 1);

        let mut s2_freq  = [0u8; 26];
        for i in 0..s2.len() {
            let cur = u8_to_idx(s2.as_bytes()[i]);
            s2_freq[cur] += 1;

            if i >= s1.len() {
                let decrement = u8_to_idx(s2.as_bytes()[i - s1.len()]);
                s2_freq[decrement] -= 1;
            }

            if s1_freq == s2_freq { return true }
        }

        false
    }
}