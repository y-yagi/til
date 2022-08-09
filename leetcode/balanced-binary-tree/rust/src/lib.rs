// Definition for a binary tree node.
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
    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        Self::depth(root) != -1
    }

    fn depth(node: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        match node {
            Some(node) => {
                let left = Self::depth(node.borrow().left.clone());
                let right = Self::depth(node.borrow().right.clone());
                if (left - right).abs() > 1 || left == -1 || right == -1 {
                    return -1;
                }
                left.max(right) + 1
            }
            None => 0,
        }
    }
}
