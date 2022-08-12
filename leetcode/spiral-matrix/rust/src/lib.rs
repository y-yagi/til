#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::spiral_order(vec![vec![1]]),
            vec![1]
        );
        assert_eq!(
            Solution::spiral_order(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]]),
            vec![1, 2, 3, 6, 9, 8, 7, 4, 5]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![];
        if matrix.len() == 0 {
            return ans;
        }

        let mut row_begin: i32 = 0;
        let mut row_end: i32 = (matrix.len() - 1) as i32;
        let mut col_begin: i32 = 0;
        let mut col_end: i32 = (matrix[0].len() - 1) as i32;

        while row_begin <= row_end && col_begin <= col_end {
            // Right
            for i in col_begin..=col_end {
                ans.push(matrix[row_begin as usize][i as usize]);
            }
            row_begin += 1;

            // Down
            for i in row_begin..=row_end{
                ans.push(matrix[i as usize][col_end as usize]);
            }
            col_end -= 1;

            // Left
            if row_begin <= row_end {
                for i in (col_begin..=col_end).rev() {
                    ans.push(matrix[row_end as usize][i as usize]);
                }
            }
            row_end -= 1;

            // Up
            if col_begin <= col_end {
                for i in (row_begin..=row_end).rev() {
                    ans.push(matrix[i as usize][col_begin as usize]);
                }
            }
            col_begin += 1;
        }

        ans
    }
}
