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
    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        return builder(0, nums.len() as i32 - 1, &nums);

        fn builder(left: i32, right: i32, nums: &Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
            if left > right {
                return None;
            }

            let mid = (left + right) / 2;

            // preorder traversal: node -> left -> right
            let mut node = TreeNode::new(nums[mid as usize]);
            node.left = builder(left, mid - 1, nums);
            node.right = builder(mid + 1, right, nums);

            return Some(Rc::new(RefCell::new(node)))
        }

    }
}