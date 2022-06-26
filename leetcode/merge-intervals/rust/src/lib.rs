#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::merge(vec![vec![1, 3], vec![2, 6], vec![8, 10], vec![15, 18]]),
            vec![vec![1, 6], vec![8, 10], vec![15, 18]]
        );
    }
}

struct Solution {}

use std::cmp::max;

impl Solution {
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        intervals.sort_by(|x, y| x[0].cmp(&y[0]));

        let mut ans = vec![];

        for interval in intervals {
            let prev_interval = match ans.last_mut() {
                Some(a) => a,
                None => {
                    ans.push(interval);
                    continue;
                }
            };

            if prev_interval[1] >= interval[0] {
                *prev_interval = vec![prev_interval[0], max(interval[1], prev_interval[1])];
            } else {
                ans.push(interval);
            }
        }

        ans
    }
}
