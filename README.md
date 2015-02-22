# hashtag
Package hashtag implements the logic of [Twitter's official Java hashtag
extractor](https://github.com/twitter/twitter-text) in Go Language. 
It implements all the logic used by twitter including correctly 
handling both # and \uFF03 as hashtag markers and correctly 
omitting hashtags that are follwed immediately by hash or url markers.

As an additional feature, this package also extracts hashtags correctly
from multiline strings.

Installation
------------
```
go get github.com/srinathh/hashtag
```

Usage
-----
First import the package
```
import "github.com/srinathh/hashtag"
```

Then call the exported function `ExtractHashtags(s string) []string`to extract
the hashtag texts. ExtractHashtags extracts hashtag texts from the input parameter
without hash markers and returns them as a slice of strings. For example
```
hashtag.ExtractHashTags(" #user1 mention #user2 here #user3 ")
```
will return thestring slice
```
[]string{"user1","user2","user3"}
```



