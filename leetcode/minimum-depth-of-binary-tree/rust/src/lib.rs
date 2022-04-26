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
    pub fn min_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        fn helper(node: Rc<RefCell<TreeNode>>, depth: &mut i32, ans: &mut i32) {
            if node.borrow().left.is_none() && node.borrow().right.is_none() {
                *ans = std::cmp::min(*ans, *depth);
                return;
            }

            if let Some(left) = node.borrow().left.clone() {
                helper(left, &mut (*depth + 1), ans);
            }
            if let Some(right) = node.borrow().right.clone() {
                helper(right, &mut (*depth + 1), ans);
            }
        }

        let mut ans: i32 = i32::max_value();

        match(root) {
            Some(node) => {
                helper(node, &mut 1, &mut ans);
                ans
            },
            None => 0,
        }
    }
}