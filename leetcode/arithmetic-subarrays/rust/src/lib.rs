#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::check_arithmetic_subarrays(
                vec![4, 6, 5, 9, 3, 7],
                vec![0, 0, 2],
                vec![2, 3, 5]
            ),
            vec![true, false, true]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn check_arithmetic_subarrays(nums: Vec<i32>, l: Vec<i32>, r: Vec<i32>) -> Vec<bool> {
        let mut ans = vec![];
        for i in 0..l.len() {
            let mut min = std::i32::MAX;
            let mut max = std::i32::MIN;
            let len = r[i] - l[i] + 1;

            for j in l[i]..=r[i] {
                min = std::cmp::min(min, nums[j as usize]);
                max = std::cmp::max(max, nums[j as usize]);
            }

            if min == max {
                ans.push(true);
            } else if (max - min) % (len - 1) != 0 {
                ans.push(false);
            } else {
                let mut diffs = vec![false; len as usize];
                let step = (max - min) / (len - 1);
                let mut j = l[i];
                while j <= r[i] {
                    if (nums[j as usize] - min) % step != 0
                        || diffs[((nums[j as usize] - min) / step) as usize]
                    {
                        break;
                    }
                    diffs[((nums[j as usize] - min) / step) as usize] = true;
                    j += 1;
                }
                ans.push(j > r[i]);
            }
        }
        ans
    }
}
