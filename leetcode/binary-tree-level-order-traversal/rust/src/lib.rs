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
use std::collections::VecDeque;

impl Solution {
	pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
		let mut queue: VecDeque<Rc<RefCell<TreeNode>>> = VecDeque::new();
		let mut ans: Vec<Vec<i32>> = Vec::new();

		if let Some(node) = root {
			queue.push_back(node);
		} else {
			return ans;
		}

		while !queue.is_empty() {
			ans.push(queue.iter().map(|node| node.borrow().val).collect());
			for _ in 0..queue.len() {
				if let Some(node) = queue.pop_front() {
					if let Some(left) = node.borrow().left.clone() {
						queue.push_back(Rc::clone(&left));
					}
					if let Some(right) = node.borrow().right.clone() {
						queue.push_back(Rc::clone(&right));
					}
				}
			}
		}
		ans
	}
}