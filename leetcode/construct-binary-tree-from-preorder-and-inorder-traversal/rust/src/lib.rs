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
            right: None,
        }
    }
}

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::builder(&preorder, &inorder)
    }

    fn builder(preorder: &[i32], inorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(root) = preorder.first() {
            let pivot_idx = inorder
                .iter()
                .enumerate()
                .find(|(_, v)| v == &root)
                .unwrap()
                .0;
            return Some(Rc::new(RefCell::new(TreeNode {
                val: *root,
                left: Self::builder(&preorder[1..(1 + pivot_idx)], &inorder[0..pivot_idx]),
                right: Self::builder(&preorder[(1 + pivot_idx)..], &inorder[(pivot_idx + 1)..]),
            })));
        } else {
            return None;
        }
    }
}
