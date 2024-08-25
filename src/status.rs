use std::fmt::{Display, Formatter};
use tonic::Code;
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

impl From<tonic::Status> for Status {
    fn from(value: tonic::Status) -> Self {
        Status {
            code: 500,
            reason: value.code().to_string(),
            message: value.message().to_string(),
            metadata: Default::default(),
        }
    }
}

impl Into<tonic::Status> for Status {
    fn into(self) -> tonic::Status {
        tonic::Status::new(Code::Internal, self.message)
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
