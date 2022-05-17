#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::letter_case_permutation("a1b2".to_string()),
            vec!["a1b2".to_string(), "a1B2".to_string(), "A1b2".to_string(), "A1B2".to_string()]
        )
    }
}

struct Solution {}

impl Solution {
    pub fn letter_case_permutation(s: String) -> Vec<String> {
        let mut ans: Vec<String> = vec![];
        Self::backtrack("".to_string(), s.clone(), &mut ans);
        ans
    }

    fn backtrack(mut curr: String, s: String, res: &mut Vec<String>) {
        if s == "".to_string() {
            res.push(curr);
            return;
        }

        let letter = s.chars().nth(0).unwrap();
        let new_s = &s[1..];
        let new_s = new_s.to_string();

        if letter.is_alphabetic() {
            let upper = letter.to_uppercase().next().unwrap();
            let lower = letter.to_lowercase().next().unwrap();

            let mut upper_s = curr.clone();
            let mut lower_s = curr.clone();

            upper_s.push(upper);
            lower_s.push(lower);

            Self::backtrack(upper_s, new_s.clone(), res);
            Self::backtrack(lower_s, new_s.clone(), res);
        } else {
            curr.push(letter);
            Self::backtrack(curr, new_s.clone(), res);
        }
    }
}