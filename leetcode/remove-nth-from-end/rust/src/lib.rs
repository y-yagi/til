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
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode { val: 0, next: head, });
        let mut fast = dummy.clone();
        let mut slow = dummy.as_mut();

        for _ in 0..n {
            fast = fast.next.unwrap();
        }

        while let Some(node) = fast.next {
            fast = node;
            slow = slow.next.as_mut().unwrap();
        }

        slow.next = slow.next.as_mut().unwrap().next.clone();
        dummy.next
    }
}