#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_circle_num(vec![vec![1, 1, 0], vec![1, 1, 0], vec![0, 0, 1]]),
            2
        );
        assert_eq!(
            Solution::find_circle_num(vec![vec![1, 0, 0], vec![0, 1, 0], vec![0, 0, 1]]),
            3
        );
        assert_eq!(
            Solution::find_circle_num(vec![
                vec![1, 0, 0, 1],
                vec![0, 1, 1, 0],
                vec![0, 1, 1, 1],
                vec![1, 0, 1, 1]
            ]),
            1
        );
    }
}

struct Solution {}

impl Solution {
    pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {
        let mut visited: Vec<bool> = vec![false; is_connected.len()];
        let mut count: i32 = 0;
        for i in 0..is_connected.len() {
            if visited[i] == false {
                Self::dfs(&is_connected, &mut visited, i);
                count += 1;
            }
        }

        count
    }

    fn dfs(is_connected: &Vec<Vec<i32>>, visited: &mut Vec<bool>, i: usize) {
        for j in 0..is_connected.len() {
            if is_connected[i][j] == 1 && visited[j] == false {
                visited[j] = true;
                Self::dfs(is_connected, visited, j);
            }
        }
    }
}
