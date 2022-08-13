#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::num_of_minutes(1, 0, vec![-1], vec![0]), 0);
        assert_eq!(
            Solution::num_of_minutes(6, 2, vec![2, 2, -1, 2, 2, 2], vec![0, 0, 1, 0, 0, 0]),
            1
        );
    }
}

struct Solution {}

impl Solution {
    pub fn num_of_minutes(n: i32, head_id: i32, manager: Vec<i32>, inform_time: Vec<i32>) -> i32 {
        let mut manager = manager;
        let mut inform_time = inform_time;
        (0..n).map(|i| Self::dfs(i as usize, &mut manager, &mut inform_time))
        .max()
        .unwrap()
    }

    fn dfs(i: usize, manager: &mut Vec<i32>, inform_time: &mut Vec<i32>) -> i32 {
        if manager[i] != -1 {
            inform_time[i] += Self::dfs(manager[i] as usize, manager, inform_time);
            manager[i] = -1;
        }
        inform_time[i]
    }
}
