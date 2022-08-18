use std::collections::HashMap;

struct AuthenticationManager {
    life: i32,
    tokens: HashMap<String,i32>,
}

impl AuthenticationManager {
    fn new(timeToLive: i32) -> Self {
        AuthenticationManager {
            life: timeToLive,
            tokens: HashMap::new(),
        }
    }

    fn generate(&mut self, token_id: String, current_time: i32) {
        self.tokens.insert(token_id, current_time+self.life);
    }

    fn renew(&mut self, token_id: String, current_time: i32) {
        if let Some(life) = self.tokens.get_mut(&token_id) {
            if current_time < *life {
                *life = current_time + self.life;
            } else {
                self.tokens.remove(&token_id);
            }
        }
    }

    fn count_unexpired_tokens(&mut self, current_time: i32) -> i32 {
        let mut targets = vec![];

        for (key, life) in self.tokens.iter_mut() {
            if current_time >= *life {
                targets.push(key.clone());
            }
        }
        for name in targets {
            self.tokens.remove(&name);
        }

        self.tokens.keys().len() as i32
    }
}