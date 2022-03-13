#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_valid_test() {
        assert_eq!(is_valid(String::from("()")), true);
        assert_eq!(is_valid(String::from("()[]{}")), true);
        assert_eq!(is_valid(String::from("(]")), false);
    }
}

pub fn is_valid(s: String) -> bool {
    let mut stack = Vec::new();
    for i in s.chars() {
        match i {
            '{' => stack.push('}'),
            '(' => stack.push(')'),
            '[' => stack.push(']'),
            '}'|')'|']' if i != stack.pop() => return false,
            _ => (),
        }
    }

    return stack.is_empty();
}