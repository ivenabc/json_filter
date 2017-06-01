```
go get github.com/Jeffail/gabs
go get github.com/iven1990/json_filter
```

### Useage

解析规则  
baseKeys 基本数据类型  
objectKeys 结构体类型  
arrayKeys 数组类型
```
{
  "baseKeys": ["hello"],
  "objectKeys":["another", "alsoInner"],
  "arrayKeys": ["inner"],
  "inner": {
    "baseKeys": [
      "value2"
    ],
    "objectKeys": ["tt"],
    "tt": {
      "baseKeys":["value5"]
    }
  },
  "alsoInner":{
    "baseKeys": ["value1"]
  },
  "another":{
    "baseKeys": ["value3"]
  }
}

```

example
```
{
  "inner":[{
    "value1":10,
    "value2":22,
    "tt":{
      "value4": 122,
      "value5": 112,
      "value6": 102
    },
    "value3":122
  }],
  "alsoInner":{
    "value1":20,
    "value2":42,
    "value3":92
  },
  "hello": 11,
  "world": 13,
  "another":{
    "value1":null,
    "value2":null,
    "value3":null
  }
}
```

result
```
{"alsoInner":{"value1":20},"another":{},"hello":11,"inner":[{"tt":{"value5":112},"value2":22}]}
```