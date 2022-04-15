struct Solution {}

impl Solution {
    pub fn duplicate_zeros(arr: &mut Vec<i32>) {
        let mut index = 0;

        while index < arr.len() {
            if arr[index] == 0 && index + 1 <= arr.len() {
                arr.insert(index + 1, 0) ;
                arr.remove(arr.len() - 1);
                index += 2;
                continue;
            }
            index += 1;
        }
    }
}