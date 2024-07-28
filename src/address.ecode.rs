// This file is @generated by prost-build.
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ErrorReason {
    /// 为某个枚举单独设置错误码
    UserNotFound = 0,
    ContentMissing = 1,
}
impl ErrorReason {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            ErrorReason::UserNotFound => "USER_NOT_FOUND",
            ErrorReason::ContentMissing => "CONTENT_MISSING",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "USER_NOT_FOUND" => Some(Self::UserNotFound),
            "CONTENT_MISSING" => Some(Self::ContentMissing),
            _ => None,
        }
    }
}
