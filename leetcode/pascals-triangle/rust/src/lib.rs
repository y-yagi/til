#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::generate(5),
            vec![
                vec![1],
                vec![1, 1],
                vec![1, 2, 1],
                vec![1, 3, 3, 1],
                vec![1, 4, 6, 4, 1]
            ]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        if num_rows == 0 {
            return vec![];
        }
        let mut ans: Vec<Vec<i32>> = vec![vec![1]];

        for i in 1..num_rows {
            let mut t = vec![1; i as usize + 1];
            let before = &ans[(i - 1) as usize];
            for j in 1..i {
                t[j as usize] = before[(j - 1) as usize] + before[j as usize];
            }
            ans.push(t);
        }

        ans
    }
}
