struct Solution {}

impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        let m = s.len() / 2;
        let l = s.len() - 1;

        for i in 0..m {
            let t  = s[l-i];
            s[l-i] = s[i];
            s[i] = t;
        }
    }
}