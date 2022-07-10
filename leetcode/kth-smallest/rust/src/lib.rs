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
    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {
        let mut ans = 0;
        let mut k = k;

        Solution::dfs(&root, &mut k, &mut ans);
        ans
    }

    fn dfs(node: &Option<Rc<RefCell<TreeNode>>>, k: &mut i32, ans: &mut i32) {
        if None == *node || *k == 0 {
            return;
        }

        if let Some(node) = node {
            Solution::dfs(&node.borrow().left, k, ans);
            *k -= 1;
            if *k == 0 {
                *ans = node.borrow().val;
            }
            Solution::dfs(&node.borrow().right, k, ans);
        }
    }
}
