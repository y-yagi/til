struct Solution {}

impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut index_of_zero = 0;
        for i in 0..nums.len() {
            if nums[i] != 0 {
                nums.swap(i, index_of_zero);
                index_of_zero += 1;
            }
        }
    }
}