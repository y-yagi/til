#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_anagrams("cbaebabacd".to_string(), "abc".to_string()), vec![0, 6]);
        assert_eq!(Solution::find_anagrams("abab".to_string(), "ab".to_string()), vec![0, 1, 2]);

    }
}

struct Solution {}

impl Solution {
    pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![];

        if s.len() < p.len() {
            return ans;
        }

        let usizeof = |c: u8| { (c - 'a' as u8) as usize };
        let s = s.as_bytes();
        let p = p.as_bytes();
        let mut s_table = [0u8; 26];
        let mut p_table = [0u8; 26];

        for i in 0..p.len() {
            p_table[usizeof(p[i])] += 1;
            s_table[usizeof(s[i])] += 1;
        }

        for i in (p.len() - 1)..s.len() {
            if s_table == p_table {
                ans.push((i+1-p.len()) as i32);
            }
            if i+1 < s.len() {
                s_table[usizeof(s[i+1-p.len()])] -= 1;
                s_table[usizeof(s[i+1])] += 1;
            }
        }


        ans
    }
}