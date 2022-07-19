#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::top_k_frequent(
                vec![
                    "i".to_string(),
                    "love".to_string(),
                    "leetcode".to_string(),
                    "i".to_string(),
                    "love".to_string(),
                    "coding".to_string()
                ],
                2
            ),
            vec!["i".to_string(), "love".to_string()]
        );
        assert_eq!(
            Solution::top_k_frequent(
                vec![
                    "the".to_string(),
                    "day".to_string(),
                    "is".to_string(),
                    "sunny".to_string(),
                    "the".to_string(),
                    "the".to_string(),
                    "the".to_string(),
                    "sunny".to_string(),
                    "is".to_string(),
                    "is".to_string(),
                ],
                4
            ),
            vec![
                "the".to_string(),
                "is".to_string(),
                "sunny".to_string(),
                "day".to_string()
            ]
        );
    }
}

struct Solution {}

use std::collections::{BinaryHeap, HashMap};

impl Solution {
    pub fn top_k_frequent(words: Vec<String>, k: i32) -> Vec<String> {
        let mut map = HashMap::<String, i32>::new();

        for word in words{
            *map.entry(word).or_default() += 1;
        }

        let mut heap = BinaryHeap::<(i32, String)>::new();

        for (word, freq) in map{
            heap.push((-freq, word));
            if heap.len() > (k as usize){
                heap.pop();
            }
        }

        let mut ans: Vec<String> = vec![];
        while let Some((_freq, word)) = heap.pop(){
            ans.push(word);
        }
        ans.reverse();
        ans
    }
}