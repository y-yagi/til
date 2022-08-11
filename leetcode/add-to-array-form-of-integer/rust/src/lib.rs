#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::add_to_array_form(vec![1, 2, 0, 0], 34),
            vec![1, 2, 3, 4]
        );
        assert_eq!(
            Solution::add_to_array_form(vec![2, 1, 5], 806),
            vec![1, 0, 2, 1]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn add_to_array_form(num: Vec<i32>, k: i32) -> Vec<i32> {
        let mut  ans = vec![];
        let mut i: i32 = (num.len() as i32) - 1;
        let mut cur = k;

        while i >= 0 || cur > 0 {
            if i >= 0 {
                cur += num[i as usize];
            }
            ans.push(cur % 10);
            cur /= 10;
            i -= 1;
        }

        ans.reverse();
        ans
    }
}