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
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        let mut node = head;
        let mut ans = 0;

        while node.is_some() {
            ans = ans * 2 + node.to_owned().unwrap().val;
            node = node.unwrap().next;
        }

        ans
    }
}