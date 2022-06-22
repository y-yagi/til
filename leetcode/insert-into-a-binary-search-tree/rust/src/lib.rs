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
    pub fn insert_into_bst(root: Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return Some(Rc::new(RefCell::new(TreeNode::new(val))));
        }

        if let Some(node) = &root {
            if node.borrow().val < val  {
                let newnode = Solution::insert_into_bst(node.borrow().right.clone(), val);
                node.borrow_mut().right = newnode;
            } else {
                let newnode = Solution::insert_into_bst(node.borrow().left.clone(), val);
                node.borrow_mut().left = newnode;
            }
        }

        root
    }
}