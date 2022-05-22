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

        let mut new_head = head;
        let mut prev = new_head;
        let mut curr = prev.unwrap().next;

        while let Some(next) = Some(curr).unwrap().next.as_mut() {
            if curr.val == next.val {
                let t = curr.val;
                while curr.val == t {
                    curr = curr.next.as_mut().unwrap();
                }
                prev.next = curr.as_mut();
            } else {
                prev = curr;
                curr = prev.next.as_mut().unwrap();
            }
        }

        new_head.unwrap().next;
    }
}
