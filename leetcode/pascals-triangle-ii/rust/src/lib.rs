#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::get_row(3), vec![1, 3, 3, 1]);
    }
}

struct Solution {}

impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let mut data: Vec<Vec<i32>> = vec![vec![1]];

        for i in 1..=row_index {
            let mut t = vec![1; i as usize + 1];
            let before = &data[(i - 1) as usize];
            for j in 1..i {
                t[j as usize] = before[(j - 1) as usize] + before[j as usize];
            }

            if i == row_index {
                return t;
            }

            data.push(t);
        }

        vec![1]
    }
}