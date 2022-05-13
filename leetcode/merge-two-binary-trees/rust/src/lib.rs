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
    pub fn merge_trees(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if root1.is_none() {
            return root2;
        }
        if root2.is_none() {
            return root1;
        }

        let root1 = root1.unwrap();
        let root2 = root2.unwrap();
        let mut ans;

        ans = TreeNode::new(root1.borrow().val + root2.borrow().val);
        ans.left = Self::merge_trees(root1.borrow().left.clone(), root2.borrow().left.clone());
        ans.right = Self::merge_trees(root1.borrow().right.clone(), root2.borrow().right.clone());

        return Some(Rc::new(RefCell::new(ans)));
    }
}
