struct Solution {}

impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        if grid.len() == 0 { return 0 }

        let mut grid = grid;
        let mut islands = 0;

        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                islands += Self::dfs(i as i32, j as i32, &mut grid);
            }
        }

        islands
    }

    pub fn dfs(i: i32, j: i32, grid: &mut Vec<Vec<char>>) -> i32 {
        let ui = i as usize;
        let uj = j as usize;

        if i < 0 || j < 0 || ui >= grid.len() || uj >= grid[0].len() || grid[ui][uj] == '0' { return 0 }

        if grid[ui][uj] == '1' {
            grid[ui][uj] = '0';
            Self::dfs(i-1, j, grid);
            Self::dfs(i+1, j, grid);
            Self::dfs(i, j-1, grid);
            Self::dfs(i, j+1, grid);

            return 1
        }

        0
    }
}