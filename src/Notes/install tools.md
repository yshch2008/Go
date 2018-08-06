# VS Code tools
## *If your net is OK,maybe it is not necessary.*

## Steps
1. 在%GOPATH%\src\ 目录下，建立golang.org 文件夹，并再新建x文件夹。  目录为  %GOPATH\src\golang.org\x\ 
2.  完成目录切换后，开始下载插件包：
    git clone https://github.com/golang/tools.git tools
   > that is not always effective,so you may need to download every single package and put them into the right folder.
   Like this exemple:
    <pre>
    get .zip from **github.com/nsf/gocode**
    put it in **d:\GoWorks\src\github.com\nsf\gocode**
    <code>

3. try to run below codes in terminal:
    <pre>
    go install github.com/ramya-rao-a/go-outline
    go install github.com/acroca/go-symbols
    go install golang.org/x/tools/cmd/guru
    go install golang.org/x/tools/cmd/gorename
    go install github.com/josharian/impl
    go install github.com/rogpeppe/godef
    go install github.com/sqs/goreturns
    go install github.com/golang/lint/golint
    go install github.com/cweill/gotests/gotests
    go install github.com/ramya-rao-a/go-outline
    go install github.com/acroca/go-symbols
    go install golang.org/x/tools/cmd/guru
    go install golang.org/x/tools/cmd/gorename
    go install github.com/josharian/impl
    go install github.com/rogpeppe/godef
    go install github.com/sqs/goreturns
    go install github.com/golang/lint/golint
    go install github.com/cweill/gotests/gotests
    <code>

> that is not always fine to install all tools successfully,so maybe need to repeat 2&3 till the hint shows no error.
> you can follow their steps 
> [article_1](http://blog.csdn.net/langzi7758521/article/details/51313521)
> [article_2](http://www.pythonsite.com/?p=429)

4. Last you need install the Delve to debug your .go [Delve](go get github.com/derekparker/delve/cmd/dlv)
5. Markdown-PDF need to set "markdown-pdf.executablePath" as "C:\\Users\\Administrator\\AppData\\Local\\CentBrowser\\Application\\chrome.exe"