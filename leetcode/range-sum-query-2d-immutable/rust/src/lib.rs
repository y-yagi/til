struct NumMatrix {
    matrix: Vec<Vec<i32>>,
}

impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        let rows: usize = matrix.len();
        let cols: usize = if rows > 0 { matrix[0].len() } else { 0 };

        let mut sums = vec![vec![0; cols + 1]; rows + 1];
        for r in 0..rows {
            for c in 0..cols {
                // Need to remove doubly added sum
                sums[r + 1][c + 1] = matrix[r][c] + sums[r][c + 1] + sums[r + 1][c] - sums[r][c];
            }
        }

        NumMatrix { matrix: sums }
    }
    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        let row1 = row1 as usize;
        let row2 = row2 as usize;
        let col1 = col1 as usize;
        let col2 = col2 as usize;

        return self.matrix[row2 + 1][col2 + 1]
            - self.matrix[row2 + 1][col1]
            - self.matrix[row1][col2 + 1]
            + self.matrix[row1][col1];
    }
}
