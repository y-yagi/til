#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::next_greater_elements(vec![1, 8, -1, -100, -1, 222, 1111111, -111111]),
            vec![8, 222, 222, -1, 222, 1111111, -1, 1]
        );
        assert_eq!(
            Solution::next_greater_elements(vec![1, 2, 1]),
            vec![2, -1, 2]
        );
        assert_eq!(
            Solution::next_greater_elements(vec![1, 2, 3, 4, 3]),
            vec![2, 3, 4, -1, 4]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let mut max = std::i32::MAX;
        let mut ans = vec![-1; nums.len()];
        let mut stack = vec![];

        for i in 0..nums.len() {
            if nums[i] > max {
                max = nums[i];
            }
            stack.push(i);
        }

        while stack.len() > 0 {
            let mut found = false;
            let index = stack.pop().unwrap();
            if nums[index] == max {
                continue;
            }

            for i in index..nums.len() {
                if nums[i] > nums[index] {
                    ans[index] = nums[i];
                    found = true;
                    break;
                }
            }

            if !found {
                for i in 0..index {
                    if nums[i] > nums[index] {
                        ans[index] = nums[i];
                        break;
                    }
                }
            }
        }

        ans
    }
}
