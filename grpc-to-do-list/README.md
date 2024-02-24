# 專案架構

## 框架

- gin

## 接口風格

- restful api

## 檔案特性

``` text
controllers：用於處理進來的請求並返回響應的路由處理器。
models：應用程式將使用的數據模型。
middlewares：中介軟體函數，如身份驗證和日誌記錄。
services：業務邏輯和服務層。
repositories：資料庫互動。
routers：路由配置。
utils：工具函數和助手。

```

## 整理套件依賴關係

```text
go mod tidy 
```
