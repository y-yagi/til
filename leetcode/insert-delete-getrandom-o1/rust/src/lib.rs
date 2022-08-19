use std::{collections::HashMap, convert::TryInto};
use rand::{thread_rng, Rng};

struct RandomizedSet {
    map : HashMap <i32, i32>,
    values: Vec<i32>,
}

impl RandomizedSet {
    fn new() -> Self {
        Self {
            values: vec![],
            map: HashMap::new()
        }
    }

    fn insert(&mut self, val: i32) -> bool {
        match self.map.get(&val) {
            Some(_) => false,
            None => {
                self.values.push(val);
                self.map.insert(val, (self.values.len()-1).try_into().unwrap());
                true
            }
        }
    }

    fn remove(&mut self, val: i32) -> bool {
        return match self.map.get(&val){
            None => false,
            Some(index) => {
                self.values[*index as usize] = self.values[self.values.len()-1];
                self.map.insert(self.values[self.values.len()-1], *index);
                self.values.pop();
                self.map.remove(&val);
                true
            }
        }
    }

    fn get_random(&self) -> i32 {
        self.values[thread_rng().gen_range(0..self.values.len())]
    }
}