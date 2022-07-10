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
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> Vec<Vec<i32>> {
        let mut result = vec![];
        let curr = vec![];
        Self::dfs(&root, target_sum, &mut result, curr);
        result
    }

    fn dfs(
        node: &Option<Rc<RefCell<TreeNode>>>,
        target_sum: i32,
        result: &mut Vec<Vec<i32>>,
        curr: Vec<i32>,
    ) {
        if node.is_none() {
            return;
        } else if let Some(n) = node {
            let node = n.borrow();
            let mut new_curr = curr.clone();
            new_curr.push(node.val);

            let v = target_sum - node.val;
            if v == 0 && node.left.is_none() && node.right.is_none() {
                result.push(new_curr);
                return;
            }

            Self::dfs(&node.left, v, result, new_curr.clone());
            Self::dfs(&node.right, v, result, new_curr);
        }
    }
}
