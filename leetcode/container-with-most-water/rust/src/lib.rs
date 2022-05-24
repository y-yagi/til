#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]), 49);
    }
}

struct Solution {}

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut left = 0;
        let mut right = height.len() - 1;

        while left < right {
            let mut area = 0;
            if height[left] > height[right] {
                area = height[right] * (right - left) as i32;
                right -= 1;
            } else {
                area = height[left] * (right - left) as i32;
                left += 1;
            }
            ans = std::cmp::max(area, ans);
        }

        ans
    }
}