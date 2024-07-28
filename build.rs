#[cfg(target_os = "macos")]
fn main() {
    use std::path::PathBuf;

    let protos = &["address/ecode/address_ecode.proto",
        "address/v1/address_v1.proto",
        "base/types.proto",
        "tran/ecode/tran_ecode.proto",
        "tran/v1/tran_v1.proto",];
    let includes = &[
        "./",
        "/Users/test/go/pkg/mod/github.com/go-bamboo/pkg@v0.0.39-0.20231128085640-d77909357dc6",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0/api",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0/third_party",
        "/Users/test/Documents/GitHub/metax/api/../third_party"];
    let out_dir = PathBuf::from("./src");
    tonic_build::configure()
        .out_dir(out_dir)
        // .compile_well_known_types(true)
        .compile(protos,includes)
        .unwrap();
}

#[cfg(not(target_os = "macos"))]
fn main() {
}