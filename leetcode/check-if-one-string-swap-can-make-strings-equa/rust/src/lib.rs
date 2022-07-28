#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::are_almost_equal("kelb".to_string(), "kelb".to_string()),
            true
        );
        assert_eq!(
            Solution::are_almost_equal("aa".to_string(), "ac".to_string()),
            false
        );
        assert_eq!(
            Solution::are_almost_equal("bank".to_string(), "kanb".to_string()),
            true
        );
        assert_eq!(
            Solution::are_almost_equal("attack".to_string(), "defend".to_string()),
            false
        );
    }
}

struct Solution {}

impl Solution {
    pub fn are_almost_equal(s1: String, s2: String) -> bool {
        if s1.len() != s2.len() {
            return false;
        }

        if s1 == s2 {
            return true;
        }

        let chars1: Vec<char> = s1.chars().collect();
        let chars2: Vec<char> = s2.chars().collect();

        let mut swapped = false;
        let mut index: i32 = -1;

        for i in 0..chars1.len() {
            if chars1[i] == chars2[i] {
                continue;
            }

            if index == -1 {
                index = i as i32;
                continue;
            }

            if swapped {
                return false;
            }

            if chars1[i] == chars2[index as usize] && chars1[index as usize] == chars2[i] {
                swapped = true;
                continue;
            }

            return false;
        }

        if index != -1 && !swapped {
            return false;
        }

        true
    }
}
