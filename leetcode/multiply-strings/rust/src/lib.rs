#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::multiply("2".to_string(), "3".to_string()),
            "6".to_string()
        );
    }
}

struct Solution {}

impl Solution {
    pub fn multiply(num1: String, num2: String) -> String {
        if num1 == "0".to_string() || num2 == "0".to_string() {
            return "0".to_string();
        }

        let len = num1.len() + num2.len();
        let mut res: Vec<u32> = vec![0; len];

        for i in (0..num1.len()).rev() {
            for j in (0..num2.len()).rev() {
                res[i + j + 1] +=
                    ((num1.as_bytes()[i] - b'0') as u32) * ((num2.as_bytes()[j] - b'0') as u32);
            }
        }

        for i in (1..len).rev() {
            res[i - 1] += res[i] / 10;
            res[i] %= 10;
        }

        let mut ans: Vec<char> = Vec::new();
        if res[0] != 0 {
            ans.push(char::from_digit(res[0], 10).unwrap());
        }

        for i in 1..len {
            ans.push(char::from_digit(res[i], 10).unwrap());
        }
        ans.into_iter().collect()
    }
}
