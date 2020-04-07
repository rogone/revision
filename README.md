# tags
## 1.根据struct tag的设置修正结构体成员值
目前支持：
* min：最小值，变量值小于此值时生效
* max：最大值，变量值大于此值时生效
* default：默认值，变量为零值时生效
* env：使用环境变量值
* flag: 使用`flag` 库解析命令行参数，支持默认值，不设置时使用零值

## 2.使用方法
定义结构体时添加tag
```golang
type T struct{
    I int `env:"T_I" min:"1" max:"0" default:"2"`
    F float64 `flag:"F,1"` // 等于 flag.FloatVar(&t.F,"F",1,"F")
}

t := &T{}
```
`tags.Revise(t)`函数将按照tag声明的顺序处理。

假设环境变量`T_I`没有设置，那么`T.I`的值随`tag`处理的过程为：
1. `env`设置为0
2. 因为0小于1，`min`设置为1
3. 因为1>0,`max`设置为0，
4. 因为是零值，所以`default`处理后为2。

实际操作中应该避免多个tag处理时相互干扰。

## 3. 类型支持

### flag: flag标准库支持的8种类型`bool`, `int`, `uint`, `int64`, `uint64`, `float64`, `string`, `time.Duration`