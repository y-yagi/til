#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::top_k_frequent(vec![1, 1, 2, 2, 3], 2), vec![1, 2]);
        assert_eq!(Solution::top_k_frequent(vec![1], 1), vec![1]);
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();

        for num in nums {
            *map.entry(num).or_insert(0) += 1;
        }

        let mut freqs: Vec<(i32, i32)> = map
            .into_iter()
            .map(|(a, b)| (a, b))
            .collect::<Vec<(i32, i32)>>();
        freqs.sort_by(|a, b| a.1.cmp(&b.1).reverse());

        let mut ans = vec![];
        for i in 0..k {
            ans.push(freqs[i as usize].0);
        }
        ans
    }
}