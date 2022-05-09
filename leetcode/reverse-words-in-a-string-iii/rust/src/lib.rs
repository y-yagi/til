#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::reverse_words("Let's take LeetCode contest".to_string()), "s'teL ekat edoCteeL tsetnoc".to_string());
        assert_eq!(Solution::reverse_words("God Ding".to_string()), "doG gniD".to_string());
    }
}

struct Solution {}

impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut ans = String::with_capacity(s.len());

        s.split_ascii_whitespace().for_each(|str| {
            ans.extend(str.chars().rev());
            ans.push(' ');
        });

        ans.truncate(ans.len()-1);
        ans
    }
}