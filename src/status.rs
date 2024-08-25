use std::fmt::{Display, Formatter};
use crate::errors::Status;

impl Status {
    fn new(reason: &str, message: &str) -> Status {
        Status {
            code: 500,
            reason: reason.to_string(),
            message: "".to_string(),
            metadata: Default::default(),
        }
    }
}

// impl PartialEq for Status {
//     fn eq(&self, other: &Self) -> bool {
//         self.code == other.code && self.reason == other.reason
//     }
// }

impl Display for Status {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.write_str(self.reason.as_str())
    }
}


#[cfg(test)]
mod tests {
    use crate::address;

    use super::*;

    #[test]
    fn it_works() {
        let s = Status::new(address::ecode::ErrorReason::UserNotFound.as_str_name(), "");
        let s1 = Status::new(address::ecode::ErrorReason::UserNotFound.as_str_name(), "");
        assert_eq!(s, s1);
    }
}
