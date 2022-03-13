#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_valid_test() {
        assert_eq!(search_insert(vec![1, 3, 5, 6], 5), 2);
        assert_eq!(search_insert(vec![1, 3, 5, 6], 2), 1);
        assert_eq!(search_insert(vec![1, 3, 5, 6], 7), 4);
    }
}

pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    if target < nums[0] {
        return 0
    }

    let mut left = 0;
    let mut right = nums.len() - 1;

    while  left <= right {
        let mid = left + (right-left)/2;
        if nums[mid] == target {
            return mid as i32
        }

        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return left as i32
}