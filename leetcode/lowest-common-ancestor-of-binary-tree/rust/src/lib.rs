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
    pub fn lowest_common_ancestor(root: Option<Rc<RefCell<TreeNode>>>, p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let p_val = p.unwrap().borrow().val;
        let q_val = q.unwrap().borrow().val;

        let mut curr_opt = root;
        while let Some(curr) = curr_opt {
            let mut curr_ref = curr.borrow_mut();
            let curr_val = curr_ref.val;
            if p_val > curr_val && q_val > curr_val {
                curr_opt = curr_ref.right.take();
            } else if p_val < curr_val && q_val < curr_val {
                curr_opt = curr_ref.left.take();
            } else {
                drop(curr_ref);
                return Some(curr);
            }
        }
        None
    }

}