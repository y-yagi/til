#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::daily_temperatures(vec![73, 74, 75, 71, 69, 72, 76, 73]),
            vec![1, 1, 4, 2, 1, 1, 0, 0]
        );
        assert_eq!(
            Solution::daily_temperatures(vec![30, 40, 50, 60]),
            vec![1, 1, 1, 0]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![0; temperatures.len()];
        let mut stack = vec![];

        for i in 0..temperatures.len() {
            while stack.len() != 0 && temperatures[i] > temperatures[stack[stack.len()-1]] {
                let idx = stack.pop().unwrap();
                ans[idx] = (i - idx) as i32;
            }

            stack.push(i);
        }

        ans
    }
}
