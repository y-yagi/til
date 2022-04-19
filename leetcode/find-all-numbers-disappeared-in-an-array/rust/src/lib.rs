#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_disappeared_numbers(vec![4, 3, 2, 7, 8, 2, 3, 1]), vec![5, 6]);
        assert_eq!(Solution::find_disappeared_numbers(vec![1, 1]), vec![2]);
    }
}

struct Solution {}

impl Solution {
    pub fn find_disappeared_numbers(nums: Vec<i32>) -> Vec<i32> {
        let mut dup = nums.clone();

        for i in 0..dup.len() {
            let idx = (dup[i].abs() - 1) as usize;
            dup[idx] = -dup[idx].abs();
        }

        let mut ans: Vec<i32> = vec![];
        for (idx, &n) in dup.iter().enumerate() {
            if n > 0 { ans.push((idx + 1) as i32); }
        }

        ans
    }
}
