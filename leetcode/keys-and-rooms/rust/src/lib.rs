#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::can_visit_all_rooms(vec![vec![1],vec![2], vec![3], vec![]]), true);
        assert_eq!(Solution::can_visit_all_rooms(vec![vec![1, 3],vec![3, 0, 1], vec![2], vec![0]]), false);
    }
}

struct Solution {}

impl Solution {
    pub fn can_visit_all_rooms(rooms: Vec<Vec<i32>>) -> bool {
        let mut seen = vec![false; rooms.len()];
        seen[0] = true;
        let mut stack = vec![];
        stack.push(0);

        while stack.len() != 0 {
            if let Some(node) = stack.pop(){
                for nei in &rooms[node] {
                    if !seen[*nei as usize] {
                        seen[*nei as usize] = true;
                        stack.push(*nei as usize);
                    }
                }
            }
        }

        for v in seen {
            if !v {
                return false
            }
        }

        true
    }
}