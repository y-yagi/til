#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::flood_fill(vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]], 1, 1, 2),
            vec![vec![2, 2, 2], vec![2, 2, 0], vec![2, 0, 1]]
        );
        assert_eq!(
            Solution::flood_fill(vec![vec![0, 0, 0], vec![0, 1, 1]], 1, 1, 1),
            vec![vec![0, 0, 0], vec![0, 1, 1]]
        );
    }
}

struct Solution {}

impl Solution {
    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32) -> Vec<Vec<i32>> {
        let old_color = image[sr as usize][sc as usize];
        if new_color == old_color { return image }

        let mut new_image = image.clone();
        Self::fill(&mut new_image, sr, sc, old_color, new_color);

        new_image
    }

    fn fill(image: &mut Vec<Vec<i32>>, i: i32, j: i32, old_color: i32, new_color: i32) {
        if i < 0 || j < 0 || i as usize >= image.len() || j as usize >= image[0].len() { return }
        if image[i as usize][j as usize] != old_color { return }

        image[i as usize][j as usize] = new_color;
        Self::fill(image, i+1, j, old_color, new_color);
        Self::fill(image, i-1, j, old_color, new_color);
        Self::fill(image, i, j+1, old_color, new_color);
        Self::fill(image, i, j-1, old_color, new_color);
    }
}