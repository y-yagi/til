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
      right: None
    }
  }
}

struct Solution {}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn has_path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {
        Self::helper(&root, target_sum)
    }

    fn helper(node: &Option<Rc<RefCell<TreeNode>>>, sum: i32) -> bool {
        if let Some(data) = node {
            let n = data.borrow_mut();
            let new_sum = sum - n.val;
            if new_sum == 0 && n.left.is_none() && n.right.is_none() {
                return true
            }
            return Self::helper(&n.left, new_sum) || Self::helper(&n.right, new_sum)
        }

        false
    }
}