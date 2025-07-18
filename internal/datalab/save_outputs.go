package datalab

import (
	"aavev3-raw-events-decoder/internal/decoder"
	"encoding/json"
	"strings"

	"github.com/minio/minio-go"
)

func SaveRecords[T any](endpoint string, accessKeyID string, secretAccessKey string, rec []T, bucket, key string) error {
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return err
	}
	// fmt.Printf("%v\n", minioClient)

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(data))

	_, err = minioClient.PutObject(bucket, key, reader, reader.Size(), minio.PutObjectOptions{ContentType: "text/plain"})
	if err != nil {
		return err
	}
	// fmt.Printf("%v\n", info)

	return nil
}

func SaveOutputs(endpoint, bucket, output_path, accessKeyID, secretAccessKey string, allDecodedEvents decoder.DecodedEventsCollection) {

	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.ActiveUsers, bucket, output_path+"all_active_users.json")

	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Burn, bucket, output_path+"decoded_Burn.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.IsolationModeTotalDebtUpdated, bucket, output_path+"decoded_IsolationModeTotalDebtUpdated.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.ReserveDataUpdated, bucket, output_path+"decoded_ReserveDataUpdated.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.MintedToTreasury, bucket, output_path+"decoded_MintedToTreasury.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Transfer, bucket, output_path+"decoded_Transfer.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.UserEModeSet, bucket, output_path+"decoded_UserEModeSet.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Borrow, bucket, output_path+"decoded_Borrow.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.BorrowAllowanceDelegated, bucket, output_path+"decoded_BorrowAllowanceDelegated.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.BackUnbacked, bucket, output_path+"decoded_BackUnbacked.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.ReserveUsedAsCollateralDisabled, bucket, output_path+"decoded_ReserveUsedAsCollateralDisabled.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.LiquidationCall, bucket, output_path+"decoded_LiquidationCall.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.DeficitCovered, bucket, output_path+"decoded_DeficitCovered.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.ReserveUsedAsCollateralEnabled, bucket, output_path+"decoded_ReserveUsedAsCollateralEnabled.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.FlashLoan, bucket, output_path+"decoded_FlashLoan.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Repay, bucket, output_path+"decoded_Repay.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Approval, bucket, output_path+"decoded_Approval.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.OwnershipTransferred, bucket, output_path+"decoded_OwnershipTransferred.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Mint, bucket, output_path+"decoded_Mint.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Supply, bucket, output_path+"decoded_Supply.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.BalanceTransfer, bucket, output_path+"decoded_BalanceTransfer.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.DeficitCreated, bucket, output_path+"decoded_DeficitCreated.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.Withdraw, bucket, output_path+"decoded_Withdraw.json")
	SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.MintUnbacked, bucket, output_path+"decoded_MintUnbacked.json")
}
