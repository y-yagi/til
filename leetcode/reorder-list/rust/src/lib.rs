#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution {}

use std::collections::VecDeque;

impl Solution {
    pub fn reorder_list(mut head: &mut Option<Box<ListNode>>) {
        let mut queue = VecDeque::new();
        let mut first_node = head.as_mut().unwrap().next.take();

        while let Some(mut node) = first_node {
            first_node = node.next.take();
            queue.push_back(Some(node));
        }

        while !queue.is_empty() {
            if let Some(last) = queue.pop_back() {
                head.as_mut().unwrap().next = last;
                head = &mut head.as_mut().unwrap().next
            }

            if let Some(next) = queue.pop_front() {
                head.as_mut().unwrap().next = next;
                head = &mut head.as_mut().unwrap().next
            }
        }
    }
}
