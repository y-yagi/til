#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn two_sum_test1() {
        assert_eq!(two_sum(vec![2, 7, 11, 15], 9), [1, 0]);
    }
}

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    use std::collections::HashMap;

    let mut m: HashMap<i32, i32> = HashMap::new();
    for (index, num) in nums.iter().enumerate() {
        match m.get(&(target - *num)) {
            Some(&index2) => return vec![index as i32, index2],
            None => m.insert(*num, index as i32),
        };
    }

    vec![]
}