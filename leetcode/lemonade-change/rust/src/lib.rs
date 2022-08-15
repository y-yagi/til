#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::lemonade_change(vec![5, 5, 5, 10, 20]), true);
        assert_eq!(Solution::lemonade_change(vec![5, 5, 10, 10, 20]), false);
    }
}

struct Solution {}

impl Solution {
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        let mut five = 0;
        let mut ten = 0;

        for bill in bills {
            if bill == 5  {
                five += 1;
            } else if bill == 10 {
                if five == 0 {
                    return false;
                }
                five -= 1;
                ten += 1;
            } else {
                if five > 0 && ten > 0 {
                    five -= 1;
                    ten -= 1;
                } else if five >= 3 {
                    five -= 3;
                } else {
                    return false;
                }
            }
        }

        true
    }
}