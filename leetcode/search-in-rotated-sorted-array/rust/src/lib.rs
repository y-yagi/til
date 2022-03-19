#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(search(vec![4, 5, 6, 7, 0, 1, 2], 0), 4);
        assert_eq!(search(vec![4, 5, 6, 7, 0, 1, 2], 3), -1);
        assert_eq!(search(vec![1], 0), -1);
    }
}

pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let mut start: usize = 0;
    let mut end: usize = nums.len() - 1;

    while start <= end {
        let mid = start + (end-start)/2;
        if nums[mid] == target {
            return mid as i32
        }

        if nums[mid] >= nums[start] {
            if target >= nums[start] && target < nums[mid] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        } else {
            if target <= nums[end] && target > nums[mid] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }

    -1
}
