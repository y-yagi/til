#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::k_closest(vec![vec![1, 3], vec![-2, 2]], 1),
            vec![vec![-2, 2]]
        );
        assert_eq!(
            Solution::k_closest(vec![vec![3, 3], vec![5, -1], vec![-2, 4]], 2),
            vec![vec![3, 3], vec![-2, 4]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn k_closest(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let distances: Vec<i32>  = points.iter().map(|p| p[0] * p[0] + p[1] * p[1]).collect();
        let kth = Self::quick_select(distances, k as usize);
        points.into_iter().filter(|p| p[0] * p[0] + p[1] * p[1] <= kth).collect::<Vec<_>>()
    }

    fn quick_select(distances: Vec<i32>, k: usize) -> i32 {
        if k == 1 {
            return distances.into_iter().min().unwrap_or(-1);
        }

        if distances.len() == k {
            return distances.into_iter().max().unwrap_or(-1);
        }

        let m = (distances[0] + distances[distances.len() - 1]) / 2;

        let (smaller, bigger): (Vec<i32>, Vec<i32>) = distances.into_iter().partition(|&d| d <= m);

        if smaller.len() >= k {
            return Self::quick_select(smaller, k);
        }

        Self::quick_select(bigger, k - smaller.len())
    }
}
