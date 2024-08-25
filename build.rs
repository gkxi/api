#[cfg(target_os = "macos")]
fn main() {
    use std::path::PathBuf;

    let protos = &["address/ecode/address_ecode.proto",
        "address/v1/address_v1.proto",
        "base/types.proto",
        "tran/ecode/tran_ecode.proto",
        "tran/v1/tran_v1.proto", ];
    let includes = &[
        "./",
        "/Users/test/go/pkg/mod/github.com/go-bamboo/pkg@v0.0.39-0.20231128085640-d77909357dc6",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0/api",
        "/Users/test/go/pkg/mod/github.com/go-kratos/kratos/v2@v2.7.0/third_party",
        "/Users/test/Documents/GitHub/metax/api/../third_party"];
    let out_dir = PathBuf::from("./src");
    tonic_build::configure()
        .type_attribute(".address.v1.NewBip44Request", "#[derive(::serde::Serialize, ::serde::Deserialize, ::validator::Validate)]")
        .field_attribute(".address.v1.NewBip44Request.mnemonic", "#[validate(length(min = 1))]")
        .type_attribute(".address.v1.NewBip44Result", "#[derive(::serde::Serialize, ::serde::Deserialize, ::validator::Validate)]")
        .type_attribute(".address.v1.NewBip44Reply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".address.v1.NewBip441Result", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".address.v1.NewBip441Reply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.ChainListRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.ChainListResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.ChainListReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.IsMultiSigAddressRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.IsMultiSigAddressReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.CreateAssociatedAccountRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.CreateAssociatedAccountResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.CreateAssociatedAccountReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BalanceRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BalanceResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BalanceReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.MinerFee1Request", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.MinerFee1Reply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.MinerFeeRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.MinerFeeResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.MinerFeeReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.SendTranRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.SendTranResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.SendTranReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.HeightRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.HeightReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.GetBlockHashByHeightRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.TxResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.Utxo", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.TxUtxo", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BalanceChanges", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.TxBalance", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BlockResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.GetBlockHashByHeightReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.GetTxByHashRequest", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.GetTxByHashReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BtcBlockResult", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .type_attribute(".tran.v1.BtcGetBlockHashByHeightReply", "#[derive(serde::Serialize, serde::Deserialize, validator::Validate)]")
        .out_dir(out_dir)
        // .compile_well_known_types(true)
        .compile(protos, includes)
        .unwrap();
}

#[cfg(not(target_os = "macos"))]
fn main() {}