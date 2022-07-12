#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::frequency_sort("tree".to_string()), "eert".to_string());
        assert_eq!(Solution::frequency_sort("cccaaa".to_string()), "cccaaa".to_string());
        assert_eq!(Solution::frequency_sort("Aabb".to_string()), "bbaA".to_string());
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn frequency_sort(s: String) -> String {
        let mut map: HashMap<char, i32> = HashMap::new();

        for c in s.chars() {
            *map.entry(c).or_insert(0) += 1;
        }

        let mut freqs: Vec<(char, i32)> = map
            .into_iter()
            .map(|(a, b)| (a, b))
            .collect::<Vec<(char, i32)>>();
        freqs.sort_by(|a, b| a.1.cmp(&b.1).reverse());

        let mut v = vec![];
        for (c, count) in freqs {
            for _ in 0..count {
                v.push(c);
            }
        }

        v.into_iter().collect()
    }
}