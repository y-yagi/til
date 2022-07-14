#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_judge(2, vec![vec![1, 2]]), 2);
        assert_eq!(Solution::find_judge(3, vec![vec![1, 3], vec![2, 3]]), 3);
        assert_eq!(
            Solution::find_judge(3, vec![vec![1, 3], vec![2, 3], vec![3, 1]]),
            -1
        );
    }
}

struct Solution {}

impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut trusts_out = vec![0; n as usize];
        let mut trusts_in = vec![0; n as usize];

        for v in trust.iter() {
            trusts_out[(v[0] - 1) as usize] += 1;
            trusts_in[(v[1] - 1) as usize] += 1;
        }

        for i in 0..n {
            if trusts_in[i as usize] == (n - 1) as usize && trusts_out[i as usize] == 0 {
                return i + 1;
            }
        }

        -1
    }
}
