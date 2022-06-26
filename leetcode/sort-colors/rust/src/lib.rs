#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut v = vec![2];
        Solution::sort_colors(&mut v);
        assert_eq!(v, vec![2]);

        let mut v = vec![2, 0, 2, 1, 1, 0];
        Solution::sort_colors(&mut v);
        assert_eq!(v, vec![0, 0, 1, 1, 2, 2]);
    }
}

struct Solution {}

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut l = 0;
        let mut r = nums.len() - 1;
        let mut curr = 0;

        while curr <= r && r > 0 {
            if nums[curr] == 0 {
                nums.swap(l, curr);
                l += 1;
                curr += 1;
            } else if nums[curr] == 2 {
                nums.swap(l, r);
                r -= 1;
            } else {
                curr += 1;
            }
        }
    }
}
