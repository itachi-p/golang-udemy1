package main

import (
	"fmt"
	"net/url"
)

//net/url

func main() {
	//URLを解析
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

	//URLに含まれたクエリーをmapで取得
	fmt.Println(u.Query())

	//URLを生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	//クエリをkey:valueのmapとして設定
	q := url.Query()
	q.Set("q", "Go言語")

	//URLに日本語（2バイト文字）が使われていた場合にエンコードされる
	url.RawQuery = q.Encode()
	fmt.Println(url)
}
