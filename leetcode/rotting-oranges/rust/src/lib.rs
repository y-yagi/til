#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::oranges_rotting(vec![vec![2, 1, 1], vec![1, 1, 0], vec![0, 1, 1]]),
            4
        );
    }
}

struct Solution {}

impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        let gridded = &mut { grid };
        let mut rottens = Vec::new();

        for (i, row) in gridded.iter().enumerate() {
            for (j, cell) in row.iter().enumerate() {
                if *cell == 2 {
                    rottens.push((i, j));
                }
            }
        }

        let mut timer = 0;
        while !rottens.is_empty() {
            let l = rottens.len();

            for _ in 0..l {
                let rotten = rottens.remove(0);
                let (i, j) = rotten;

                if i > 0 && gridded[i - 1][j] == 1 {
                    gridded[i - 1][j] = 2;
                    rottens.push((i - 1, j));
                }
                if j > 0 && gridded[j][j - 1] == 1 {
                    gridded[i][j - 1] = 2;
                    rottens.push((i, j - 1));
                }
                if i < gridded.len() - 1 && gridded[i + 1][j] == 1 {
                    gridded[i + 1][j] = 2;
                    rottens.push((i + 1, j));
                }
                if j < gridded[0].len() - 1 && gridded[i][j + 1] == 1 {
                    gridded[i][j + 1] = 2;
                    rottens.push((i, j + 1));
                }
            }
            timer += 1;
        }

        if timer != 0 {
            timer -= 1
        }

        for row in gridded {
            for cell in row {
                if *cell == 1 {
                    return -1;
                }
            }
        }
        timer
    }
}
