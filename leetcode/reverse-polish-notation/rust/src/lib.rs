#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::eval_rpn(vec![
                "2".to_string(),
                "1".to_string(),
                "+".to_string(),
                "3".to_string(),
                "*".to_string()
            ]),
            9
        );
        assert_eq!(
            Solution::eval_rpn(vec![
                "4".to_string(),
                "13".to_string(),
                "5".to_string(),
                "/".to_string(),
                "+".to_string()
            ]),
            6
        );
    }
}

struct Solution {}

impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack: Vec<i32> = vec![];

        for token in tokens {
            match token.as_str() {
                "+" => {
                    let t = stack.pop().unwrap() + stack.pop().unwrap();
                    stack.push(t);
                },
                "-" => {
                    let t1 = stack.pop().unwrap();
                    let t2 = stack.pop().unwrap();
                    stack.push(t2 - t1);
                }
                "*" => {
                    let t = stack.pop().unwrap() * stack.pop().unwrap();
                    stack.push(t);
                }
                "/" => {
                    let t1 = stack.pop().unwrap();
                    let t2 = stack.pop().unwrap();
                    stack.push(t2 / t1);
                }
                _ => { stack.push(token.parse().unwrap()) },
            }
        }

        stack[0]
    }
}
