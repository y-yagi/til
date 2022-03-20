#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::group_anagrams(vec![String::from("eat"), String::from("tea"), String::from("tan"), String::from("ate"), String::from("nat"), String::from("bat")]),
            vec![vec![String::from("bat")], vec![String::from("nat"), String::from("tan")], vec![String::from("ate")], vec![String::from("eat"), String::from("tea")]]);
        assert_eq!(Solution::group_anagrams(vec![String::from("")]), vec![vec![String::from("")]]);
    }
}

struct Solution { }

use std::collections::HashMap;

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut dict: HashMap<String, Vec<String>> = HashMap::new();
        let mut ans = vec![];

        for str in strs.iter() {
            let mut chars: Vec<_> = str.chars().collect();
            chars.sort();
            let sorted_str: String = chars.into_iter().collect();
            let entry = dict.entry(sorted_str).or_insert(vec![]);
            entry.push(str.to_string());
        }

        for v in dict.values() {
            ans.push(v.to_owned());
        }

        ans
    }
}