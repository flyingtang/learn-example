
这是抄https://github.com/benbjohnson/genesis的项目

功能：
    将资源文件转化成二进制并存入go 文件，并提供接口给外部调用
例如：
    有一个config.json 文件 ，生成config.go 文件
    然后将可以通过http或者直接调用的方式，得到资源信息

    ```go

    	f := AssetFile("/config.json")
    	fmt.Println(string(f.data))
    ```
    或者

    ```go

	http.Handle("/", http.FileServer(FileSystem()))
	http.ListenAndServe(":8000", nil)
	```