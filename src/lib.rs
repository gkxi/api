#[path =""]
pub mod address {
    #[path ="address.ecode.rs"]
    pub mod ecode;
    #[path ="address.v1.rs"]
    pub mod v1;
}

pub mod base;

#[path =""]
pub mod tran {
    #[path ="tran.ecode.rs"]
    pub mod ecode;
    #[path ="tran.v1.rs"]
    pub mod v1;
}


pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
