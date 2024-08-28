### 描述

该接口提供版本：v1.0.0+

添加模版到模版套餐

### 输入参数

| 参数名称          | 参数类型 | 必选 | 描述                      |
| ----------------- | -------- | ---- | ------------------------- |
| biz_id            | uint32   | 是   | 业务ID                    |
| template_space_id | uint32   | 是   | 模版空间ID                |
| template_ids      | uint32   | 是   | 模版ID列表，最多200个     |
| template_set_ids  | []uint32 | 是   | 模版套餐ID列表，最多200个 |

### 调用示例

```json
{
  "template_ids": [
    1,
    2
  ],
  "template_set_ids": [
    1,
    2
  ]
}
```

### 响应示例

```json
{
  "data": {}
}
```

### 响应参数说明

| 参数名称 | 参数类型 | 描述     |
| -------- | -------- | -------- |
| data     | object   | 响应数据 |
