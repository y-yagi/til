#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut result: Vec<Vec<i32>> = vec![];
        Self::travel(&root, &mut result, 0);
        Self::reverse(&mut result);

        result
    }

    fn travel(node: &Option<Rc<RefCell<TreeNode>>>, arr: &mut Vec<Vec<i32>>, depth: usize) {
        let node = match node {
            Some(x) => Rc::clone(x),
            None => return,
        };

        if arr.len() <= depth {
            arr.push(vec![]);
        }

        arr[depth].push(node.borrow().val);
        Self::travel(&node.borrow().left, arr, depth + 1);
        Self::travel(&node.borrow().right, arr, depth + 1);
    }

    fn reverse(arr: &mut Vec<Vec<i32>>) {
        for ind in (1..arr.len()).step_by(2) {
            arr[ind].reverse();
        }
    }
}
