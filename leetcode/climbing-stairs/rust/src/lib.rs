#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::climb_stairs(2), 2);
    }
}

struct Solution {}

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        if n <= 2 { return n }

        let mut one_step_before = 2;
        let mut two_steps_before = 1;
        let mut ans = 0;

        for i in 3..=n {
            ans = one_step_before + two_steps_before;
            two_steps_before = one_step_before;
            one_step_before = ans;
        }

        ans
    }
}