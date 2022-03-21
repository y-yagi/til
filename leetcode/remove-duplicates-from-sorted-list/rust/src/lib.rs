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

impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        let mut h = head;
        let mut curr = h.as_mut().unwrap();

        while let Some(next) = curr.next.as_mut() {
            if curr.val == next.val {
                curr.next = next.next.take()
            } else {
                curr = curr.next.as_mut().unwrap();
            }
        }

        h
    }
}