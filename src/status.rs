use std::fmt::{Display, Formatter};
use axum::{
    http::StatusCode,
    Json,
    response::{IntoResponse, Response},
};
use tonic::Code;
use crate::errors::Status;
use serde_json::json;
use serde_json::Value::String;
use tonic::metadata::MetadataMap;

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
        let mut mm = MetadataMap::new();
        mm.insert("reason", self.reason.parse().unwrap());
        tonic::Status::with_metadata(Code::Internal, self.message, mm)
    }
}

impl IntoResponse for Status {
    fn into_response(self) -> Response {
        let body = Json(json!(self));
        (StatusCode::from_u16(self.code as u16).unwrap(), body).into_response()
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
