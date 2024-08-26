
#[path = ""]
pub mod address {
    #[path = "address.ecode.rs"]
    pub mod ecode;
    #[path = "address.v1.rs"]
    pub mod v1;
}

pub mod base;

pub mod errors;
pub mod status;

#[path = ""]
pub mod tran {
    #[path = "tran.ecode.rs"]
    pub mod ecode;
    #[path = "tran.v1.rs"]
    pub mod v1;
}

