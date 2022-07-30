#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::check_straight_line(vec![vec![2, 1], vec![4, 2], vec![6, 3]]),
            true
        );
        assert_eq!(
            Solution::check_straight_line(vec![vec![0, 0], vec![0, 1], vec![0, -1]]),
            true
        );
        assert_eq!(
            Solution::check_straight_line(vec![vec![1, 2], vec![2, 3]]),
            true
        );
    }
}

struct Solution {}

impl Solution {
    pub fn check_straight_line(coordinates: Vec<Vec<i32>>) -> bool {
        let (x0, y0, x1, y1) = (
            coordinates[0][0],
            coordinates[0][1],
            coordinates[1][0],
            coordinates[1][1],
        );
        let (dx, dy) = (x1 - x0, y1 - y0);

        for coordinate in coordinates {
            let (x, y) = (coordinate[0], coordinate[1]);
            if dx * (y - y1) != dy * (x - x1) {
                return false;
            }
        }

        true
    }
}
