use serde::{Deserialize, Deserializer, Serialize, Serializer};
use serde::ser::SerializeStruct;
use crate::errors::Status;

pub(crate) struct SpringResponse<T> {
    pub success: bool,
    pub code: String,
    pub message: String,
    pub data: T,
}

impl<T> SpringResponse<T> {
    fn new(success: bool, code: String, message: String, data: T) -> Self {
        Self {
            success,
            code,
            message,
            data,
        }
    }
}

impl<T> From<Status> for SpringResponse<T>
    where T: Default
{
    fn from(value: Status) -> Self {
        Self {
            success: false,
            code: value.reason,
            message: value.message,
            data: T::default(),
        }
    }
}

impl<T> Serialize for SpringResponse<T>
    where T: Serialize
{
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error> where S: Serializer {
        let mut state = serializer.serialize_struct("SpringResponse", 4)?;
        state.serialize_field("success", &self.success)?;
        state.serialize_field("code", &self.code)?;
        state.serialize_field("message", &self.message)?;
        state.serialize_field("data", &self.data)?;
        state.end()
    }
}

impl<'de, T> Deserialize<'de> for SpringResponse<T>
    where T: Deserialize<'de>
{
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error> where D: Deserializer<'de> {
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_seri() {
        let a = SpringResponse::new(false, "OK".to_string(), "OK".to_string(), "OK");
        let b = serde_json::to_string(&a).unwrap();
        println!(" {}", b);
    }
}