# ContractTerms


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Deadline`                                                          | [time.Time](https://pkg.go.dev/time#Time)                           | :heavy_check_mark:                                                  | The deadline for the contract.                                      |
| `Deliver`                                                           | [][ContractDeliverGood](../../models/shared/contractdelivergood.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Payment`                                                           | [ContractPayment](../../models/shared/contractpayment.md)           | :heavy_check_mark:                                                  | N/A                                                                 |