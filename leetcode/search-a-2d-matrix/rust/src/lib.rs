#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::search_matrix(vec![vec![1]], 1), true);
        assert_eq!(
            Solution::search_matrix(
                vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]],
                3
            ),
            true
        );
        assert_eq!(
            Solution::search_matrix(
                vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]],
                13
            ),
            false
        );
    }
}

struct Solution {}

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut x = 0;
        let mut y = 0;

        while x < matrix.len() && y < matrix[0].len() {
            if matrix[x][y] == target {
                return true;
            }

            if x < matrix.len() - 1 && target >= matrix[x + 1][y] {
                x += 1
            } else {
                y += 1
            }
        }

        false
    }
}
