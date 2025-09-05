# vbs-sdk-go

## Installation

```bash
go get github.com/nccasia/vbs-sdk-go
```

## Sử dụng Fabric Client để gọi Chaincode

### Khởi tạo Fabric Client

```go
package main

import (
    "fmt"
    "github.com/nccasia/vbs-sdk-go/pkg/client/fabric"
    "github.com/nccasia/vbs-sdk-go/pkg/core/config"
    "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/chaincode"
)

func main() {
    // Cấu hình client
    api := "http://localhost:8889"
    userCode := "UserCode1"
    appCode := "AppCode1"
    privK := "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgxONXM9QezTD7JvSs\ndfMuV64CD8b0jCa2qpc3qJDGjYagCgYIKoZIzj0DAQehRANCAATNAe5f9X2LLSCt\nFP2AFwzYL6dNRb6rckxSMfVd27mjYrKSPelRY/l5bIKLbAi1iXXcUoJie6mwnLdR\nWMl8wJYf\n-----END PRIVATE KEY-----\n"
    mspDir := "./"

    config, err := config.NewConfig(api, userCode, appCode, privK, mspDir)
    if err != nil {
        return err
    }

    fabricClient, err := fabric.InitFabricClient(config)
    if err != nil {
        return err
    }
}
```

### Query Chaincode (Đọc dữ liệu)

```go
// Query tất cả assets
args := []string{}
queryReq := chaincode.QueryChaincodeReq{
    UserID:        "tutest004",
    ChaincodeName: "contract-fabric-testing-2app1748855175839756",
    FunctionName:  "GetAllAssets",
    Args:          args,
}

res, err := fabricClient.QueryChaincode(queryReq, nil)
if err != nil {
    return err
}
fmt.Printf("Query result: %+v\n", res)

// Query một asset cụ thể
args = []string{"tu00011"}
queryReq = chaincode.QueryChaincodeReq{
    UserID:        "tutest003",
    ChaincodeName: "contract-fabric-testing-2app1748855175839756",
    FunctionName:  "ReadAsset",
    Args:          args,
}

res, err = fabricClient.QueryChaincode(queryReq, nil)
if err != nil {
    return err
}
fmt.Printf("Query result: %+v\n", res)
```

### Invoke Chaincode (Ghi dữ liệu)

```go
// Tạo asset mới
args := []string{"tu0009", "green", "31", "aba", "20000"}
invokeReq := chaincode.InvokeChaincodeReqBody{
    UserID:        "tutest004",
    ChaincodeName: "contract-fabric-testing-2app1748855175839756",
    FunctionName:  "CreateAsset",
    Args:          args,
}

res, err := fabricClient.InvokeChaincode(invokeReq, nil)
if err != nil {
    return err
}
fmt.Printf("Invoke result: %+v\n", res)
```

## Sử dụng FiscoBcos Client để gọi Contract

### Khởi tạo FiscoBcos Client

```go
package main

import (
    "fmt"
    "github.com/nccasia/vbs-sdk-go/pkg/client/fiscobcos"
    "github.com/nccasia/vbs-sdk-go/pkg/core/config"
    "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/contract"
)

func main() {
    // Cấu hình client
    api := "http://localhost:8889"
    userCode := "UserCode1"
    appCode := "AppCode1"
    privK := "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgxONXM9QezTD7JvSs\ndfMuV64CD8b0jCa2qpc3qJDGjYagCgYIKoZIzj0DAQehRANCAATNAe5f9X2LLSCt\nFP2AFwzYL6dNRb6rckxSMfVd27mjYrKSPelRY/l5bIKLbAi1iXXcUoJie6mwnLdR\nWMl8wJYf\n-----END PRIVATE KEY-----\n"
    mspDir := "./"

    config, err := config.NewConfig(api, userCode, appCode, privK, mspDir)
    if err != nil {
        return err
    }

    fiscobcosClient, err := fiscobcos.NewFiscoBcosClient(config)
    if err != nil {
        return err
    }
}
```

### Query Contract (Đọc dữ liệu)

```go
// Query contract
args := []string{}
queryReq := contract.QueryContractReqBody{
    UserID:          "tutest04",
    ContractAddress: "0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085",
    FunctionName:    "get",
    Args:            args,
}

res, err := fiscobcosClient.QueryContract(queryReq)
if err != nil {
    return err
}
fmt.Printf("Query result: %+v\n", res)
fmt.Printf("Payload: %s\n", res.Body.Payload)
```

### Invoke Contract (Ghi dữ liệu)

```go
// Invoke contract
args := []string{"Hello, World!!!!!!!!"}
invokeReq := contract.InvokeContractReqBody{
    UserID:          "tutest04",
    ContractAddress: "0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085",
    FunctionName:    "set",
    Args:            args,
}

res, err := fiscobcosClient.InvokeContract(invokeReq)
if err != nil {
    return err
}
fmt.Printf("Invoke result: %+v\n", res)
```

## Các tham số cấu hình

### Fabric Client
- `api`: URL của API gateway
- `userCode`: Mã người dùng
- `appCode`: Mã ứng dụng
- `privK`: Private key của người dùng
- `mspDir`: Đường dẫn thư mục MSP

### FiscoBcos Client
- `api`: URL của API gateway
- `userCode`: Mã người dùng
- `appCode`: Mã ứng dụng
- `privK`: Private key của người dùng
- `mspDir`: Đường dẫn thư mục MSP

## Note
- Đảm bảo chaincode/contract đã được deploy trước khi gọi
- UserID phải tồn tại trong hệ thống
- ContractAddress cho FiscoBcos phải chính xác
- Private key phải đúng định dạng PEM

## Hướng dẫn lấy thông tin cấu hình

### Lấy userCode, appCode và privateKey
1. Đăng nhập vào trang **VBSN**
2. Vào mục **Quản lý Node phân quyền** -> **Chứng chỉ của tôi**
3. Lấy `userCode` và `appCode` từ thông tin tài khoản
4. Thực hiện **tải file keypair** của dự án xuống để lấy `privateKey`

### Lấy địa chỉ API
1. Vào mục **Dự án đã phát hành**
2. Chọn **Chi tiết dự án**
3. Vào **Tham số cấu hình**
4. Lấy **Địa chỉ trung tâm dữ liệu** làm giá trị cho biến `api`

### Chống supply chain attack 
  - Tool: https://github.com/google/osv-scanner
  - Run tool (OS: Linux)
  ```bash
    $ go install github.com/google/osv-scanner/v2/cmd/osv-scanner@latest
    $ cd vbs-sdk-go
    $ osv-scanner --recursive . > ./osv-scanner-result/result.text

  ```