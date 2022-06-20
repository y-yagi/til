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

type Node = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    pub fn is_symmetric(root: Node) -> bool {
        Self::is_mirror(&root, &root)
    }

    fn is_mirror(node1: &Node, node2: &Node) -> bool {
        match (node1, node2) {
            (None, None) => true,
            (Some(node1), Some(node2)) => {
                let (node1, node2) = (node1.borrow(), node2.borrow());
                node1.val == node2.val
                    && Self::is_mirror(&node1.left, &node2.right)
                    && Self::is_mirror(&node1.right, &node2.left)
            }
            (_, _) => false,
        }
    }
}
