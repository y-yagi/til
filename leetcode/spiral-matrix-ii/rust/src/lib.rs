#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::generate_matrix(3),
            vec![vec![1, 2, 3], vec![8, 9, 4], vec![7, 6, 5]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let mut ans = vec![vec![0; n as usize]; n as usize];
        let mut cnt = 1;

        for layer in 0..(n+1) / 2 {
            //  travese from left to right
            for ptr in layer..n - layer {
                ans[layer as usize][ptr as usize] = cnt;
                cnt += 1;
            }

            //  travese from top to bottom
            for ptr in (layer + 1)..(n-layer) {
                ans[ptr as usize][(n - layer - 1)as usize] = cnt;
                cnt += 1;
            }

            //  travese from right to left
            for ptr in (layer + 1)..(n-layer) {
                ans[(n - layer - 1)as usize][(n - ptr - 1)as usize] = cnt;
                cnt += 1;
            }

            //  travese from bottom to top
            for ptr in (layer + 1)..(n-layer - 1) {
                ans[(n - ptr - 1)as usize][layer as usize] = cnt;
                cnt += 1;
            }
        }

        ans
    }
}
