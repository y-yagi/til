#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(generate_parenthesis(3), vec!["((()))".to_string(), "(()())".to_string(),  "(())()".to_string(),  "()(())".to_string(), "()()()".to_string()]);
        assert_eq!(generate_parenthesis(1), vec!["()".to_string()]);
    }
}

pub fn generate_parenthesis(n: i32) -> Vec<String> {
    fn backtrace(s: String, open: i32, close: i32) -> Vec<String> {
        let mut res = vec![];
        if open == 0 && close == 0 {
            return vec![s];
        }

        if open > 0 {
            res.append(&mut backtrace(s.clone() + "(", open-1, close+1));
        }
        if close > 0 {
            res.append(&mut backtrace(s.clone() + ")", open, close-1));
        }
        res
    }

    backtrace("".to_string(), n, 0)
}