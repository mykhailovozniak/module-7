# go app on Heroku

database in mongodb at mongo atlas

## Link to /hello
https://young-springs-45765.herokuapp.com/hello

## Link to /materials
https://young-springs-45765.herokuapp.com/materials

## Link to /post
#### 1,2,3-cached
https://young-springs-45765.herokuapp.com/post?postId=1

# how to check code coverage

```
$ cd cmd/web
$ go tool cover -html=cp.out

$ cd pkg/models/mongodb
$ go tool cover -html=cp.out
```
