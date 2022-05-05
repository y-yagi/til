#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::unique_morse_representations(vec!["a".to_string()]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn unique_morse_representations(words: Vec<String>) -> i32 {
        const MORSE: [&str; 26] = [
            ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
            "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
        ];

        words
            .iter()
            .map(|w| w.chars().map(|c| MORSE[c as usize - 'a' as usize]).collect::<String>())
            .collect::<std::collections::HashSet<_>>()
            .len() as i32
    }
}