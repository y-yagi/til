struct Solution {}

impl Solution {
    pub fn max_area_of_island(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 0 { return 0 }

        let mut grid = grid;
        let mut ans = 0;

        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                ans = std::cmp::max(ans, Self::dfs(i as i32, j as i32, &mut grid));
            }
        }

        ans
    }

    fn dfs(i: i32, j: i32, grid: &mut Vec<Vec<i32>>) -> i32 {
        let ui = i as usize;
        let uj = j as usize;

        if i < 0 || j < 0 || ui > grid.len()-1 || uj > grid[0].len()-1 || grid[ui][uj] == 0 { return 0 }

        grid[ui][uj] = 0;
        return 1 + Self::dfs(i-1, j, grid) +  Self::dfs(i+1, j, grid) + Self::dfs(i, j-1, grid) + Self::dfs(i, j+1, grid);
    }
}
