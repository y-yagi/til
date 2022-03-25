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

use std::i64;
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root == None {
            return true
        }

        return Self::validate(&root, i64::MIN, i64::MAX);
    }

    pub fn validate(node: &Option<Rc<RefCell<TreeNode>>>, min: i64, max: i64) -> bool {
        if let Some(n) = node {
            let v = n.borrow().val as i64;

            if v <= min || v >= max { return false; }

            Self::validate(&n.borrow().left, min, v) && Self::validate(&n.borrow().right, v, max)
        } else {
            true
        }
    }
}