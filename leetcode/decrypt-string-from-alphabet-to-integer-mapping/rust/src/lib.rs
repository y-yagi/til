#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::freq_alphabets("10#11#12".to_string()),
            "jkab".to_string()
        );
    }
}

struct Solution {}

impl Solution {
    pub fn freq_alphabets(s: String) -> String {
        let mut chars = s.chars().rev();
        let mut ans: Vec<char> = vec![];

        while let Some(c) = chars.next() {
            let d = match c {
                '#' => {
                    chars.next().unwrap() as u8 - b'0' + (chars.next().unwrap() as u8 - b'0') * 10
                }
                _ => c as u8 - b'0',
            } + b'a' - 1;

            ans.insert(0, d as char)
        }

        ans.into_iter().collect()
    }
}
