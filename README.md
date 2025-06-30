# vbs-sdk-go

## Installation

```bash
go get github.com/nccasia/vbs-sdk-go
```

### Chống supply chain attack 
  - Tool: https://github.com/google/osv-scanner
  - Run tool (OS: Linux)
  ```bash
    $ go install github.com/google/osv-scanner/v2/cmd/osv-scanner@latest
    $ cd vbs-sdk-go
    $ osv-scanner --recursive . > ./osv-scanner-result/result.text

  ```