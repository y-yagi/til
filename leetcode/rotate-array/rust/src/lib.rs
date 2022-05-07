struct Solution {}

impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let k = k as usize % nums.len();
        // Reverse entire array
        nums.reverse();
        // // Reverse first k elements
        nums[..k].reverse();
        // //  Reverse from k to the end
        nums[k..].reverse();
    }
}