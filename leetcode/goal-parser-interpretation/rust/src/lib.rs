#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::interpret("G()(al)".to_string()), "Goal".to_string());
    }
}

struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn interpret(command: String) -> String {
        let mut ans: Vec<String> = vec![];
        let mut i = 0;
        let chars: Vec<char> = command.chars().collect();

        while i < command.len() {
            if chars[i] == 'G' {
                ans.push("G".to_string());
                i += 1;
            } else if chars[i] == '(' {
                if chars[i+1] == ')' {
                    ans.push("o".to_string());
                    i += 2;
                } else  {
                    ans.push("al".to_string());
                    i += 4;
                }
            }
        }

        ans.into_iter().collect()
    }
}