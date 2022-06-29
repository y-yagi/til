struct Solution {}

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut col: i32 = (matrix[0].len() - 1) as i32;
        let mut row = 0;

        while col >= 0 && row <= matrix.len() - 1 {
            if target == matrix[row][col as usize] {
                return true;
            } else if target < matrix[row][col as usize] {
                col -= 1;
            } else {
                row += 1;
            }
        }

        false
    }
}