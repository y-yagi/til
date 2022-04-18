#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::height_checker(vec![1, 1, 4, 2, 1, 3]), 3);
        assert_eq!(Solution::height_checker(vec![5, 1, 2, 3, 4]), 5);
        assert_eq!(Solution::height_checker(vec![1, 2, 3]), 0);
    }
}

struct Solution {}

impl Solution {
    pub fn height_checker(heights: Vec<i32>) -> i32 {
        let mut dup = heights.clone();
        dup.sort();

        heights.iter().zip(&dup).map(|(i, j)| if i != j { 1 } else { 0 }).sum()
    }
}