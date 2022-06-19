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
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut ans = vec![];

        if root.is_none() {
            return ans;
        }

        Self::dig(root, &mut ans);
        ans
    }

    fn dig(node: Option<Rc<RefCell<TreeNode>>>, ans: &mut Vec<i32>) {
        if let Some(n) = node {
            if !n.borrow().left.is_none() {
                Self::dig(n.borrow().left.clone(), ans);
            }
            if !n.borrow().right.is_none() {
                Self::dig(n.borrow().right.clone(), ans);
            }
            ans.push(n.borrow().val);
        }
    }
}