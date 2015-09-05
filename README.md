gin (forked from [gin](https://github.com/codegangsta/gin) )[![wercker status](https://app.wercker.com/status/f413ccbd85cfc4a58a37f03dd7aaa87e "wercker status")](https://app.wercker.com/project/bykey/f413ccbd85cfc4a58a37f03dd7aaa87e)
========

## Motivation of Fork

 - my main go file is under src/main/main.go. Normally I use `go build -o bin/main main` which is not workable under the origianl gin.

 - Monitor all files under `src` folder.

 - Monitor not only `.go` file, but also `.tmpl` file.

 Thus I modified the command input a bit, in order to be able to pass the folder name for building and the src for watching any changes.

========

`gin` is a simple command line utility for live-reloading Go web applications. 
Just run `gin` in your app directory and your web app will be served with 
`gin` as a proxy. `gin` will automatically recompile your code when it 
detects a change. Your app will be restarted the next time it receives an 
HTTP request.

`gin` adheres to the "silence is golden" principle, so it will only complain 
if there was a compiler error or if you succesfully compile after an error.

## Installation

Assuming you have a working Go environment and `GOPATH/bin` is in your 
`PATH`, `gin` is a breeze to install:

```shell
go get github.com/stefanzan/gin
```

Then verify that `gin` was installed correctly:

```shell
gin -h
```

How to use it ? (2015/09/05)

You can modify the `run.sh` to feed your own needs. Basically,
```shell
gin -a "8000" -p "3000" -m "main"
```
`main` is a folder under `src` folder, in which there is a main file named `main.go` . 



## Supporting Gin in Your Web app
`gin` assumes that your web app binds itself to the `PORT` environment 
variable so it can properly proxy requests to your app. Web frameworks 
like [Martini](http://github.com/codegangsta/martini) do this out of 
the box.

## Using flags?
When you normally start your server with [flags](https://godoc.org/flag)
if you want to override any of them when running `gin` we suggest you 
instead use [github.com/namsral/flag](https://github.com/namsral/flag)
as explained in [this post](http://stackoverflow.com/questions/24873883/organizing-environment-variables-golang/28160665#28160665)
