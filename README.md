## 亮讯接口调用说明

## Usage

```
func main(){
    client := NewClient("123", "AAAAAAAAAA")
    resp, err := client.GetCardUsageInfo(&ReqCardUsageInfo{
    		Msisdns: []string{"89860412102090029371"},
    		Month:   "202107",
    	})
    if err != nil {
        log.Fatal(err)
    }
    log.Println(resp)
}
```

## API 文档

[API 文档](http://ec.upiot.net/api_doc/#_%E6%8E%A5%E5%8F%A3%E8%B0%83%E7%94%A8%E8%AF%B4%E6%98%8E) 是上海亮讯物联卡平台接口文档

## 申请API key / secret
### 申请说明
- 在新版本API接口中，将脱离原先获取token的认证模式，使用基于API_KEY和API_SECRET的API认证。用户首先需要在 [物联卡管理平台](http://ec.upiot.net/login/)上申请自己的API_KEY和API_SECRET.

## License

```
Copyright 2021 MingGuang

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```