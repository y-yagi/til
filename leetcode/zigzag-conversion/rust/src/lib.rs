#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::convert(String::from("PAYPALISHIRING"), 3), "PAHNAPLSIIGYIR");
        assert_eq!(Solution::convert(String::from("PAYPALISHIRING"), 4), "PINALSIGYAHRPI");
    }
}

struct Solution {}

impl Solution {
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows == 1{
            return s;
        }

        let mut strings = vec![String::new(); num_rows as usize];
        let mut i = 0;
        let mut down = true;

        for c in s.chars() {
            strings[i].push(c);
            i = if down { i + 1 } else { i - 1 };
            down =  down == (i > 0 && i < num_rows as usize - 1);
        }

        strings.concat()
    }
}