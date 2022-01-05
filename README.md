# mycoin-cli-geth
## Table of Contents

 * [專案描述](#專案描述)
 * [執行專案](#執行專案)

## 專案描述


## 執行專案

### 執行測試

```bash
$ make test
```

#### 執行應用程式

```bash
#到專案目錄下
$ cd path_to_dir/mycoin-cli-geth

# 下載第三方套件
$ go mod download

# 編譯專案(輸出到當前目錄下,檔案名為main)
$ go build -o main . 

# 執行應用程式
$ ./main 

```
#### 編譯合約
```bash
$ abigen --abi=data/abi/MyCoin.json --bin=data/bin/MyCoin.bin  --pkg=mycoin --out=contracts/mycoin/mycoin.go
```