#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_repeated_dna_sequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT".to_string()),
            vec!["AAAAACCCCC".to_string(), "CCCCCAAAAA".to_string()]
        );
    }
}

struct Solution {}

use std::collections::HashSet;

impl Solution {
    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        if s.len() <= 10 {
            return vec![];
        }

        let mut seen: HashSet<String> = HashSet::new();
        let mut repeated: HashSet<String> = HashSet::new();

        for i in 0..s.len() - 9 {
            let ten = &s.to_string()[i..i + 10];
            if !seen.insert(ten.to_string()) {
                repeated.insert(ten.to_string());
            }
        }

        repeated.into_iter().collect()
    }
}
