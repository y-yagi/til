#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

#[derive(Debug)]
struct ParkingSystem {
    capacities: Vec<i32>,
}

impl ParkingSystem {
    fn new(big: i32, medium: i32, small: i32) -> Self {
        Self {
            capacities: vec![big, medium, small],
        }
    }

    fn add_car(&mut self, car_type: i32) -> bool {
        if self.capacities[(car_type - 1) as usize] < 1 {
            return false;
        }
        self.capacities[(car_type - 1) as usize] -= 1;
        true
    }
}
