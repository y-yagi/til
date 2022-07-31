#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::diagonal_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]]),
            25
        );
    }
}

struct Solution {}

impl Solution {
    pub fn diagonal_sum(mat: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        let n = mat.len();

        for i in 0..n {
            ans += mat[i][i];
            ans += mat[n - 1 - i][i];
        }

        if n % 2 == 0 {
            return ans;
        }

        ans - mat[n / 2][n / 2]
    }
}
