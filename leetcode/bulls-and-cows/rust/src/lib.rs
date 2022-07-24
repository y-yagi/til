#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::get_hint("1807".to_string(), "7810".to_string()),
            "1A3B".to_string()
        );
        assert_eq!(
            Solution::get_hint("1123".to_string(), "0111".to_string()),
            "1A1B".to_string()
        );
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn get_hint(secret: String, guess: String) -> String {
        let mut dict: HashMap<char, i32> = HashMap::new();
        let secret_chars: Vec<char> = secret.chars().collect();
        let guess_chars: Vec<char> = guess.chars().collect();
        let mut bulls = 0;
        let mut cows = 0;
        let mut remain: Vec<char> = vec![];

        for i in 0..secret.len() {
            if secret_chars[i] == guess_chars[i] {
                bulls += 1;
            } else {
                *dict.entry(secret_chars[i]).or_insert(0) += 1;
                remain.push(guess_chars[i]);
            }
        }

        for i in 0..remain.len() {
            if let Some(v) = dict.get(&remain[i]){
                if v > &0 {
                    cows += 1;
                    *dict.entry(remain[i]).or_insert(0) -= 1;
                }
            }
        }

        format!("{}A{}B", bulls, cows)
    }
}
