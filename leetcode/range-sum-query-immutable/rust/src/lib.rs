#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

struct NumArray {
    nums: Vec<i32>,
    sums: Vec<i32>,
}

impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        let mut sums: Vec<i32> = vec![0; nums.len()+1];
        for i in 0..nums.len() {
            sums[i+1] = nums[i] + sums[i];
        }

        Self {
            nums: nums,
            sums: sums,
        }
    }

    fn sum_range(&self, left: i32, right: i32) -> i32 {
        self.sums[right as usize + 1] - self.sums[left as usize]
    }
}
