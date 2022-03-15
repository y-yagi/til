#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_valid_test() {
        assert_eq!(num_unique_emails(vec!["test.email+alex@leetcode.com".to_string(), "test.e.mail+bob.cathy@leetcode.com".to_string(),"testemail+david@lee.tcode.com".to_string()]), 2);
        assert_eq!(num_unique_emails(vec!["a@leetcode.com".to_string(), "b@leetcode.com".to_string(), "c@leetcode.com".to_string()]), 3);
    }
}

use std::collections::HashSet;

pub fn num_unique_emails(emails: Vec<String>) -> i32 {

    let mut unique: HashSet<String> = HashSet::new();

    for email in emails.iter() {
        let mut parts = email.split("@");
        let local = parts.next().unwrap();
        let domain = parts.next().unwrap();

        let local = match local.find("+") {
            Some(pos) => &local[..pos],
            None => local,
        };

        let local: String = local.replace(".", "");

        let email = format!("{}@{}", local, domain);
        unique.insert(email);
    }

    return unique.len() as i32
}