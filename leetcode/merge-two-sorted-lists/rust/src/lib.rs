// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

struct Solution {}

use std::option::Option;

impl Solution {
    pub fn merge_two_lists(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode::new(0));
        let mut cur = &mut head;
        let mut l1_cur = list1;
        let mut l2_cur = list2;

        while l1_cur.is_some() && l2_cur.is_some() {
            let mut l1_t = l1_cur.take();
            let mut l2_t = l2_cur.take();

            if let (Some(mut l1_head), Some(mut l2_head)) = (l1_t, l2_t) {
                if l1_head.val <= l2_head.val {
                    l1_cur = l1_head.next.take();
                    l2_cur = Some(l2_head);
                    cur = cur.next.get_or_insert(l1_head);
                } else {
                    l2_cur = l2_head.next.take();
                    l1_cur = Some(l1_head);
                    cur = cur.next.get_or_insert(l2_head);
                }
            }
        }

        if l1_cur.is_some() {
            cur.next = l1_cur;
        }
        if l2_cur.is_some() {
            cur.next = l2_cur;
        }

        head.next
    }
}