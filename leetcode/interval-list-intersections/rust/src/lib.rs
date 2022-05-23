#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::interval_intersection(
                vec![vec![0, 2], vec![5, 10], vec![13, 23], vec![24, 25]],
                vec![vec![1, 5], vec![8, 12], vec![15, 24], vec![25, 26]]
            ),
            vec![
                vec![1, 2],
                vec![5, 5],
                vec![8, 10],
                vec![15, 23],
                vec![24, 24],
                vec![25, 25]
            ]
        );
    }
}

struct Solution {}

use std::cmp::{max, min};

impl Solution {
    pub fn interval_intersection(first_list: Vec<Vec<i32>>, second_list: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut i = 0;
        let mut j = 0;
        let mut ans = vec![];

        while i < first_list.len() && j < second_list.len() {
            let lo = max(first_list[i][0], second_list[j][0]);
            let hi = min(first_list[i][1], second_list[j][1]);

            if lo <= hi  {
                ans.push(vec![lo, hi]);
            }
            if first_list[i][1] < second_list[j][1] {
                i += 1;
            } else {
                j += 1;
            }
        }

        ans
    }
}
