#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::partition_labels("ababcbacadefegdehijhklij".to_string()), vec![9, 7, 8]);
        assert_eq!(Solution::partition_labels("eccbbbbdec".to_string()), vec![10]);
    }
}

struct Solution {}

impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let (mut v, mut l, mut r) = (vec![], 0, 0);
        for (i, c) in s.chars().enumerate() {
            r = std::cmp::max(r, s.rfind(c).unwrap());
            if i == r {
                v.push((r - l + 1) as i32);
                l = r + 1
            }
        }
        v
    }
}
