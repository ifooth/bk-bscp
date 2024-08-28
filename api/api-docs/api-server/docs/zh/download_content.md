### 描述

该接口提供版本：v1.0.0+

获取文件内容元数据

### 输入参数

| 参数名称 | 参数类型 | 必选 | 描述   |
| -------- | -------- | ---- | ------ |
| biz_id   | uint32   | 是   | 业务ID |

#### HEADER设置

| 参数名称                 | 参数类型 | 必选 | 描述                                 |
| ------------------------ | -------- | ---- | ------------------------------------ |
| X-Bscp-App-Id            | uint32   | 否   | 如果是应用配置项，则设置该应用ID     |
| X-Bscp-Template-Space-Id | uint32   | 否   | 如果是模版配置项，则设置该模版空间ID |
| X-Bkapi-File-Content-Id  | string   | 是   | 上传文件内容的SHA256值               |

**说明**：X-Bscp-App-Id和X-Bscp-Template-Space-Id有且只能设置其中一个，设置两个或都不设置将报错

### 调用示例

```

```

### 响应示例

```json
{
  "data": {
    "byte_size": 23,
    "sha256": "1e7f8c518aed1f83cd5cabac8b48785063c5ad9e9af6c368cf0e0c2ba6bc67d6"
  }
}
```

### 响应参数说明

| 参数名称 | 参数类型 | 描述     |
| -------- | -------- | -------- |
| data     | object   | 响应数据 |

#### data

| 参数名称  | 参数类型 | 描述             |
| --------- | -------- | ---------------- |
| byte_size | uint32   | 文件内容的字节数 |
| sha256    | array    | 文件内容的sha256 |
