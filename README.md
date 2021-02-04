# kconfig
解析yaml或者json的配置文件
> InitConfig("test.yaml",false)
>

获取字符串
> GetString("arr.0.a.bb") // 返回"nnn"
>

获取数组
> GetStringArray("arr2.arr3") // 返回["aa","bb","cc"]

获取集合
> GetStringMap("arr.0.a") // 返回{"zz":"zzz","bb":"nnn"}

yaml例子:
```yaml
arr:
  - a:
      zz: zzz
      bb: nnn
arr2:
  arr3:
    - aa
    - bb
    - cc
```