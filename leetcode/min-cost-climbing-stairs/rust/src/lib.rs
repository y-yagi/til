#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::min_cost_climbing_stairs(vec![10, 15, 20]), 15);
    }
}

struct Solution {}

use std::cmp::min;

impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let mut count: Vec<i32> = cost.clone();
        count.push(0);

        for i in 2..cost.len()+1 {
            count[i] = min(count[i-1], count[i-2]) + count[i];
        }

        count[cost.len()]
    }
}