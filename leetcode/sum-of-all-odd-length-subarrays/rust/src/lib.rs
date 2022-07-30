#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::sum_odd_length_subarrays(vec![1, 4, 2, 5, 3]), 58);
        assert_eq!(Solution::sum_odd_length_subarrays(vec![1, 2]), 3);
    }
}

struct Solution {}

impl Solution {
    pub fn sum_odd_length_subarrays(arr: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = arr.len();

        for i in 0..n {
            let ending_here = i + 1;
            let starting_here = n - i;
            let total_subarrays = ending_here * starting_here;
            let mut odd_subarrays = total_subarrays / 2;
            if total_subarrays %2 == 1 {
                odd_subarrays += 1;
            }
            ans += odd_subarrays * arr[i] as usize;
        }

        ans as i32
    }
}