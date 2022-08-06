#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_rotation(vec![vec![0, 1], vec![1, 0]], vec![vec![1,0], vec![0, 1]]), true);
    }
}

struct Solution {}

impl Solution {
    pub fn find_rotation(mat: Vec<Vec<i32>>, target: Vec<Vec<i32>>) -> bool {
        if mat == target {
            return true;
        }

        let l = mat.len();
        let mut mat = mat;

        for _ in 0..3 {
            let buf = mat.clone();

            for i in 0..l {
                for j in 0..l {
                    mat[i][j] = buf[l - j - 1][i];
                }
            }

            if target == mat {
                return true;
            }
        }

        false
    }
}