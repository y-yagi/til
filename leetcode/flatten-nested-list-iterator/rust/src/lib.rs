#[derive(Debug, PartialEq, Eq)]
pub enum NestedInteger {
  Int(i32),
  List(Vec<NestedInteger>)
}

struct NestedIterator {
    stack: Vec<NestedInteger>,
}

impl NestedIterator {

    fn new(mut nestedList: Vec<NestedInteger>) -> Self {
        nestedList.reverse();
        Self {stack: nestedList}
    }

    fn next(&mut self) -> i32 {
        self.acquire().unwrap_or(-1)
    }

    fn has_next(&mut self) -> bool {
        if let Some(v) = self.acquire() {
            self.stack.push((NestedInteger::Int(v)));
            return true;
        }

        false
    }

    fn acquire(&mut self) -> Option<i32> {
        if let Some(i) = self.stack.pop() {
            match i {
                NestedInteger::Int(a) => return Some(a),
                NestedInteger::List(list) => {
                    for v in list.into_iter().rev() {
                        self.stack.push(v);
                    }
                    return self.acquire();
                }
            }
        }
        None
    }
}