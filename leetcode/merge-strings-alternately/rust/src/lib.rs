#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::merge_alternately("abc".to_string(), "pqr".to_string()), "apbqcr".to_string());
        assert_eq!(Solution::merge_alternately("ab".to_string(), "pqrs".to_string()), "apbqrs".to_string());
    }
}

struct Solution {}

impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {
        let mut ans: Vec<char> = vec![];
        let chars_1: Vec<char> = word1.chars().collect();
        let chars_2: Vec<char> = word2.chars().collect();
        let (mut i, mut j) = (0, 0);

        while i < chars_1.len() || j < chars_2.len() {
            if i < chars_1.len() {
                ans.push(chars_1[i]);
            }

            if j < chars_2.len() {
                ans.push(chars_2[j]);
            }
            i += 1;
            j += 1;
        }

        ans.into_iter().collect()
    }
}