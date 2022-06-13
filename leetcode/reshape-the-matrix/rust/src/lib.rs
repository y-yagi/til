#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::matrix_reshape(vec![vec![1, 2]], 1, 0),
            vec![vec![1, 2]]
        );
        assert_eq!(
            Solution::matrix_reshape(vec![vec![1, 2], vec![3, 4]], 1, 4),
            vec![vec![1, 2, 3, 4]]
        );
        assert_eq!(
            Solution::matrix_reshape(vec![vec![1, 2], vec![3, 4]], 2, 4),
            vec![vec![1, 2], vec![3, 4]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn matrix_reshape(mat: Vec<Vec<i32>>, r: i32, c: i32) -> Vec<Vec<i32>> {
        let n = mat.len();
        let m = mat[0].len();

        if r * c != (n * m) as i32 {
            return mat;
        };

        let r = r as usize;
        let c = c as usize;

        let mut ans = vec![vec![0; c]; r];
        for i in 0..r * c {
            ans[i / c][i % c] = mat[i / m][i % m];
        }

        ans
    }
}
