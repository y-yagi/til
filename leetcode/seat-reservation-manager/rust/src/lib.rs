use std::cmp::Reverse;
use std::collections::BinaryHeap;

struct SeatManager {
    seats: BinaryHeap<Reverse<i32>>,
}

impl SeatManager {
    fn new(n: i32) -> Self {
        Self{seats: BinaryHeap::from((1..=n).into_iter().map(|x| Reverse(x)).collect::<Vec<_>>())}
    }

    fn reserve(&mut self) -> i32 {
        self.seats.pop().unwrap().0
    }

    fn unreserve(&mut self, seat_number: i32) {
        self.seats.push(Reverse(seat_number));
    }
}