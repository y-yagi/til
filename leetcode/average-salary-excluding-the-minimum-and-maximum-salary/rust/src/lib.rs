#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::average(vec![4000, 3000, 1000, 2000]), 2500.0);
        assert_eq!(Solution::average(vec![1000, 2000, 3000]), 2000.0);
    }
}

struct Solution {}

impl Solution {
    pub fn average(salary: Vec<i32>) -> f64 {
        let mut total = 0;
        let mut min = i32::MAX;
        let mut max = 0;
        let count = salary.len() - 2;

        for x in salary {
            total += x;
            max= std::cmp::max(x, max);
            min = std::cmp::min(x, min);
        }

        total -= max;
        total -= min;
        (total as f64)  / (count as f64)
    }
}