# 接口文档
1、新增货品

2、修改货品

3、删除货品

2、进货

3、拆箱

4、售货

5、查库存
## 一、新增货品
### 1、接口名称
```
/goods/createcategories
```
### 2、入参
``` json
{
    "userId": 1,    //用户id、操作人员 必传
    "projectId": 1, //项目id 必传
    "categories": [
        {
            "name": "白酒",     // 商品名称 string 必传
            "remark": "xxx",    // 商品备注 string
            "extra_info":{},    // 额外信息 json
        },
        {
            "name": "啤酒",     // 商品名称 string 必传
            "remark": "xxx",    // 商品备注 string
            "extra_info":{},    // 额外信息 json
        }
    ]
}
```
### 3、出参
``` json
{
    "status": 0,
    "message": "success",
    "data": [
        {
            "name": "白酒",     // 商品名称 string
            "code": "CATE001",  // 商品代码 string
            "remark": "xxx",    // 商品备注 string
            "extra_info":{},    // 额外信息 json
        },
        {
            "name": "啤酒",     // 商品名称 string
            "code": "CATE002",  // 商品代码 string
            "remark": "xxx",    // 商品备注 string
            "extra_info":{},    // 额外信息 json
        }
    ]
}
```
## 二、修改货品
### 1、接口名称
```
/goods/updatecategory
```
### 2、入参
``` json
{
    "userId": 1,    //用户id、操作人员 必传
    "projectId": 1, //项目id 必传
    "code": "CATE001",  // 商品代码 string 必传
    "name": "白酒",     // 商品名称 string
    "remark": "xxx",    // 商品备注 string
    "extra_info":{},    // 额外信息 json
}
```
### 3、出参
``` json
{
    "status": 0,
    "message": "success",
    "data": true
}
```
## 二、修改货品
### 1、接口名称
```
/goods/deletecategory
```
### 2、入参
``` json
{
    "userId": 1,    //用户id、操作人员 必传
    "projectId": 1, //项目id 必传
    "code": "CATE001",  // 商品代码 string 必传
}
```
### 3、出参
``` json
{
    "status": 0,
    "message": "success",
    "data": true
}
```
