#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::can_make_arithmetic_progression(vec![3, 5, 1]), true);
        assert_eq!(Solution::can_make_arithmetic_progression(vec![1, 2, 4]), false);
    }
}

struct Solution {}

impl Solution {
    pub fn can_make_arithmetic_progression(arr: Vec<i32>) -> bool {
        if arr.len() < 3 {
            return true;
        }

        let mut arr = arr;
        arr.sort();

        let differ = arr[1] - arr[0];

        for i in 2..arr.len() {
            if arr[i] - arr[i-1] != differ {
                return false;
            }
        }

        true
    }
}