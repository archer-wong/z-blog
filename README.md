
# Introduction

Z-blog is a lightweight blogging system based on golang (macaron framework), which uses markdown editing.

# Features

- Support markdown editing
- Support batch import of markdown files, batch release articles
- Support markdown pictures directly uploaded to the qiniu
- Clear structure and easy secondary development

# Technology

- [semantic-ui](https://semantic-ui.com/)
- [jquery-file-upload](https://github.com/blueimp/jQuery-File-Upload)
- [editor.md](https://pandao.github.io/editor.md/)
- [simplePagination.js](http://flaviusmatis.github.io/simplePagination.js/)
- [canvas-nest.js](https://github.com/hustcc/canvas-nest.js)
- [macaron](https://go-macaron.com/)
- [qiniu](https://developer.qiniu.com/kodo/sdk/1238/go)


# Show

Demo：[ArcherWong博客](https://www.archerwong.cn/)


Home Page:
![](http://markdown.archerwong.cn/2019-01-04-21-01-49_1.png)


Detail Page:
![](http://markdown.archerwong.cn/2019-01-04-21-10-07_2.png)

Login Page:
![](http://markdown.archerwong.cn/2019-01-04-21-03-08_3.png)


Manage Page:
![](http://markdown.archerwong.cn/2019-01-04-21-03-27_4.png)

# How to use

- Golang environment
- Mysql database
- The configuration file is located  at [conf/app.ini]
- Compiling the source code , then start the program


If you are not familiar with how golang is deployed, please refer to the following article

https://beego.me/docs/deploy/
