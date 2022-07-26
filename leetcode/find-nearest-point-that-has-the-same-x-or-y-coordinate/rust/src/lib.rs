#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::nearest_valid_point(
                3,
                4,
                vec![vec![1, 2], vec![3, 1], vec![2, 4], vec![2, 3], vec![4, 4]]
            ),
            2
        );
    }
}

struct Solution {}

impl Solution {
    pub fn nearest_valid_point(x: i32, y: i32, points: Vec<Vec<i32>>) -> i32 {
        let mut ans: i32 = -1;
        let mut smallest_distance = i32::MAX;

        for i in 0..points.len() {
            let point = &points[i];
            if point[0] != x && point[1] != y {
                continue
            }

            let distance = (point[0] - x).abs() + (point[1] - y ).abs();
            if distance < smallest_distance {
                smallest_distance = distance;
                ans = i as i32;
            }

        }

        ans
    }
}
