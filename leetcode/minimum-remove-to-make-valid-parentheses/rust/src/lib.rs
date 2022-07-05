#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::min_remove_to_make_valid("lee(t(c)o)de)".to_string()),
            "lee(t(c)o)de".to_string()
        );
        assert_eq!(
            Solution::min_remove_to_make_valid("a)b(c)d".to_string()),
            "ab(c)d".to_string()
        );
    }
}

struct Solution {}

impl Solution {
    pub fn min_remove_to_make_valid(s: String) -> String {
        let mut stack = Vec::new();
        let mut chars: Vec<char> = s.chars().collect();
        let invalid = '*';

        for i in 0..chars.len() {
            match chars[i] {
                '(' => stack.push(i),
                ')' => {
                    if let Some(_) = stack.pop() {
                        continue;
                    } else {
                        chars[i] = invalid;
                    }
                }
                _ => continue,
            }
        }

        while let Some(i) = stack.pop() {
            chars[i] = '*';
        }

        chars
            .iter()
            .filter_map(|&x| if x == invalid { None } else { Some(x) })
            .collect::<String>()
    }
}
