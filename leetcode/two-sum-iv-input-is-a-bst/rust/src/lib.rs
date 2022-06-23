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
use std::collections::HashSet;

impl Solution {
    pub fn find_target(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> bool {
        let mut queue= vec![root];
        let mut set  = HashSet::new();

        while queue.len() > 0 {
            if let Some(data) = queue.pop() {
                if let Some(node) = data {
                  let v = node.borrow().val;
                  if set.contains(&v) {
                      return true
                  }
                  set.insert(k - v);
                  queue.push(node.borrow().left.clone());
                  queue.push(node.borrow().right.clone());
                }
            }
        }

        false
    }
}