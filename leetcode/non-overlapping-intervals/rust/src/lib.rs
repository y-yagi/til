#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::erase_overlap_intervals(vec![vec![1, 2], vec![2, 3], vec![3, 4], vec![1, 3]]),
            1
        );
    }
}

struct Solution {}

impl Solution {
    pub fn erase_overlap_intervals(intervals: Vec<Vec<i32>>) -> i32 {
        let mut end = std::i32::MIN;
        let mut erased = 0;
        let mut intervals = intervals;

        intervals.sort_by_key(|i| i[1]);

        for i in intervals {
            if i[0] >= end {
                end = i[1];
            } else {
                erased += 1;
            }
        }
        erased
    }
}
