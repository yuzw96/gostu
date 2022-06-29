package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
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

		doc, err := html.Parse(bytes.NewReader(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1:%v\n", err)
			os.Exit(1)
		}

		//for _,link :=range visit(nil,doc){
		//	fmt.Println(link)
		//}

		//outline(nil,doc)

		eleCount(doc)
		fmt.Println(count)

	}
}

func visit(links []string, n *html.Node) []string {
	//if n.Type == html.ElementNode && n.Data == "a"{
	//	for _,a :=range n.Attr{
	//		if a.Key == "href"{
	//			links = append(links, a.Val)
	//		}
	//	}
	//}
	//
	//for c:=n.FirstChild;c!=nil;c=c.NextSibling {
	//	links = visit(links,c)
	//}

	//递归替换循环
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		//outline有入栈操作，但是没有相应的出栈操作
		//当outline调用自身时，被调用者接收的是stack的拷贝。被调用者的入栈操作，修改的是stack的拷贝，而不是调用者的stack，因此当函数返回时，调用者的stack并未被修改
		stack = append(stack, n.Data)
	}
	fmt.Println(stack)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

var count = make(map[string]int)

func eleCount(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		eleCount(c)
	}
}
