#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::exist(
                vec![
                    vec!['A', 'B', 'C', 'E'],
                    vec!['S', 'F', 'C', 'S'],
                    vec!['A', 'D', 'E', 'E']
                ],
                "ABCCED".to_string()
            ),
            true
        );
        assert_eq!(
            Solution::exist(
                vec![
                    vec!['A', 'B', 'C', 'E'],
                    vec!['S', 'F', 'C', 'S'],
                    vec!['A', 'D', 'E', 'E']
                ],
                "SEE".to_string()
            ),
            true
        );
        assert_eq!(
            Solution::exist(
                vec![
                    vec!['A', 'B', 'C', 'E'],
                    vec!['S', 'F', 'C', 'S'],
                    vec!['A', 'D', 'E', 'E']
                ],
                "ABCB".to_string()
            ),
            false
        );
        assert_eq!(
            Solution::exist(
                vec![
                    vec!['a', 'b'],
                ],
                "ba".to_string()
            ),
            true
        );
    }
}

struct Solution {}

impl Solution {
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        let mut board = board;

        for i in 0..board.len() {
            for j in 0..board[0].len() {
                if Solution::helper(&mut board, &word, 0, i, j) {
                    return true;
                }
            }
        }

        false
    }

    fn helper(board: &mut Vec<Vec<char>>, word: &str, pos: usize, i: usize, j: usize) -> bool {
        if i == board.len() || j == board[0].len() {
            return false;
        }

        if word.chars().nth(pos) != Some(board[i][j]) {
            return false;
        }

        if pos == word.len() - 1 {
            return true;
        }

        let cur = board[i][j];
        board[i][j] = '0';

        let mut found = Self::helper(board, word, pos + 1, i + 1, j);
        if !found {
            found = Self::helper(board, word, pos + 1, i, j + 1);
        }
        if !found && i > 0 {
            found = Self::helper(board, word, pos + 1, i - 1, j);
        }
        if !found && j > 0 {
            found = Self::helper(board, word, pos + 1, i, j - 1);
        }

        board[i][j] = cur;
        found
    }
}
