# intab-webserver

## 模板引擎 Here

``` shell
#Install
 go get github.com/shiyanhui/hero
 go get github.com/shiyanhui/hero/hero
 go get golang.org/x/tools/cmd/goimports

#Run
 sh maketemplate.sh
```

Hero总共有九种语句，他们分别是：

- 函数定义语句 `<%: func define %>`
  - 该语句定义了该模板所对应的函数，如果一个模板中没有函数定义语句，那么最终结果不会生成对应的函数。
  - 该函数最后一个参数必须为`*bytes.Buffer`或者`io.Writer`, hero会自动识别该参数的名字，并把把结果写到该参数里。
  - 例:
    - `<%: func UserList(userList []string, buffer *bytes.Buffer) %>`
    - `<%: func UserList(userList []string, w io.Writer) %>`
    - `<%: func UserList(userList []string, w io.Writer) (int, error) %>`

- 模板继承语句 `<%~ "parent template" %>`
  - 该语句声明要继承的模板。
  - 例: `<%~ "index.html" >`

- 模板include语句 `<%+ "sub template" %>`
  - 该语句把要include的模板加载进该模板，工作原理和`C++`中的`#include`有点类似。
  - 例: `<%+ "user.html" >`

- 包导入语句 `<%! go code %>`
  - 该语句用来声明所有在函数外的代码，包括依赖包导入、全局变量、const等。

  - 该语句不会被子模板所继承

  - 例:

    ```go
    <%!
    	import (
          	"fmt"
        	"strings"
        )

    	var a int

    	const b = "hello, world"

    	func Add(a, b int) int {
        	return a + b
    	}

    	type S struct {
        	Name string
    	}

    	func (s S) String() string {
        	return s.Name
    	}
    %>
    ```

- 块语句 `<%@ blockName { %> <% } %>`

  - 块语句是用来在子模板中重写父模中的同名块，进而实现模板的继承。

  - 例:

    ```html
    <!DOCTYPE html>
    <html>
        <head>
            <meta charset="utf-8">
        </head>

        <body>
            <%@ body { %>
            <% } %>
        </body>
    </html>
    ```

- Go代码语句 `<% go code %>`

  - 该语句定义了函数内部的代码部分。

  - 例:

    ```go
    <% for _, user := userList { %>
        <% if user != "Alice" { %>
        	<%= user %>
        <% } %>
    <% } %>

    <%
    	a, b := 1, 2
    	c := Add(a, b)
    %>
    ```

- 原生值语句 `<%==[t] variable %>`

  - 该语句把变量转换为string。

  - `t`是变量的类型，hero会自动根据`t`来选择转换函数。`t`的待选值有:
    - `b`: bool
    - `i`: int, int8, int16, int32, int64
    - `u`: byte, uint, uint8, uint16, uint32, uint64
    - `f`: float32, float64
    - `s`: string
    - `bs`: []byte
    - `v`: interface

    注意：
    - 如果`t`没有设置，那么`t`默认为`s`.
    - 最好不要使用`v`，因为其对应的转换函数为`fmt.Sprintf("%v", variable)`，该函数很慢。

  - 例:

    ```go
    <%== "hello" %>
    <%==i 34  %>
    <%==u Add(a, b) %>
    <%==s user.Name %>
    ```

- 转义值语句 `<%= statement %>`

  - 该语句把变量转换为string后，又通过`html.EscapesString`记性转义。
  - `t`跟上面原生值语句中的`t`一样。
  - 例:

    ```go
    <%= a %>
    <%= a + b %>
    <%= Add(a, b) %>
    <%= user.Name %>
    ```

- 注释语句 `<%# note %>`

  - 该语句注释相关模板，注释不会被生成到go代码里边去。
  - 例: `<# 这是一个注释 >`.
