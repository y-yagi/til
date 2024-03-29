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
    pub fn sum_of_left_leaves(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::helper(&root, false)
    }

    fn helper(root: &Option<Rc<RefCell<TreeNode>>>, is_left: bool) -> i32 {
        match root {
            None => 0,
            Some(node) => {
                let node = node.borrow();
                let (left, right) = (&node.left, &node.right);

                if left.is_none() && right.is_none() {
                    if is_left {
                        node.val
                    } else {
                        0
                    }
                } else {
                    Self::helper(left, true) + Self::helper(right, false)
                }
            }
        }
    }
}
