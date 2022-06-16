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
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut cur = &mut head;

        while cur.is_some() {
            if cur.as_ref().unwrap().val == val {
                *cur = cur.take().unwrap().next;
            } else {
                cur = &mut cur.as_mut().unwrap().next;
            }
        }

        head
    }
}