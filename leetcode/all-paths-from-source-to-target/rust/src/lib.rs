struct Solution {}

use std::collections::VecDeque;

impl Solution {
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut q = VecDeque::new();
        q.push_back((0i32, vec![]));

        let mut result = vec![];
        while let Some((node, mut path)) = q.pop_front() {
            path.push(node);

            if node as usize == graph.len() - 1 {
                result.push(path);
            } else {
                for i in 0..graph[node as usize].len().checked_sub(1).unwrap_or(0) {
                    let n = &graph[node as usize][i];
                    q.push_back((*n, path.clone()));
                }
                if let Some(n) = graph[node as usize].last() {
                    q.push_back((*n, path));
                }
            }
        }

        result
    }
}