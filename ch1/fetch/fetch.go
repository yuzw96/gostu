package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//缺省
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		//http.Get是用于创建http Get请求的函数，resp为请求结果
		resp, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			//当请求出现错误时，我们用Exit函数来进行关闭，并返回响应的错误码，如1
			os.Exit(1)
		}
		//方式1
		//resp中Body中包含一个可读的服务器响应流，ReadAll方法会从response中读取全部内容
		b, err := ioutil.ReadAll(resp.Body)

		//方式2 直接输出结果
		//_,err := io.Copy(os.Stdout,resp.Body)

		//既然是流，毫无疑问在用完之后要将其关闭，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s",b)

		//输出状态码，如200 OK
		fmt.Println(resp.Status, b)
	}
}
