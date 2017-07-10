# go-idec

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->
**Table of Contents**

- [go-idec](#go-idec)
    - [Usage](#usage)
- [License](#license)
- [Authors](#authors)

<!-- markdown-toc end -->


IDEC protocol in golang

## Usage

```go
go get github.com/Difrex/go-idec

cat > test.go <<EOF
package main

import (
	"fmt"
	idec "github.com/Difrex/go-idec"
)

func main() {
    txt := "aWkvb2svcmVwdG8vNFJnUkh4NUtJUHhwOWN1QUNKZ1gKZGV2ZWxvcC4xNgoxNDYzNjU5NDMzCmJ0aW1vZmVldgpzdGF0aW9uMTMsIDEzCkFsbApSZTog0J3QtdGB0LXRgtC10LLRi9C1INC/0YDQvtC10LrRgtGLCgoK0KMg0LzQtdC90Y8g0LjQtyDQv9GA0L7QtdC60YLQvtCyINC90LAg0LPQuNGC0YXQsNCx0LUg0YHQsNC80L7QtSDQuNC90YLQtdGA0LXRgdC90L7QtSBodHRwczovL2dpdGh1Yi5jb20vYnRpbW9mZWV2L2VtdWNoaXAKCtCt0YLQviDQsdGL0LvQsCDQvNC+0Y8g0L/QvtC/0YvRgtC60LAg0L3QsNC/0LjRgdCw0YLRjCDRjdC80YPQu9GP0YLQvtGAINC/0YDQvtGB0YLQtdC50YjQtdCz0L4g0LrQvtC80L/RjNGO0YLQtdGA0LAsINC90LAg0L/RgNC40LzQtdGA0LUgQ2hpcDgg0LggU3VwZXJDaGlwLiDQktGB0LUg0LjQs9GA0Ysg0LTQu9GPINGN0YLQuNGFINGB0LjRgdGC0LXQvCDRgNCw0LHQvtGC0LDRjtGCICjQt9CwINC40YHQutC70Y7Rh9C10L3QuNC10Lwg0L7QtNC90L7QuSwg0LrQsNC20LXRgtGB0Y8pLiDQndC+INCyINGN0LzRg9C70Y/RgtC+0YDQtSDQtdGB0YLRjCDQv9Cw0YDQsCDQvtGI0LjQsdC+0LosINC60L7RgtC+0YDRi9C1INGPINGC0LDQuiDQuCDQvdC1INGA0LXRiNC40LsgKNC+0LTQvdCwINGB0LLRj9C30LDQvdCwINGBINGA0LDQt9C80LXRgNC+0Lwg0L7QutC90LAgUXQsINCy0YLQvtGA0LDRjyDRgSDRg9C60LDQt9Cw0YLQtdC70Y/QvNC4IGMrKykuCgovLyDQsCDQstC+0L7QsdGJ0LUg0Y8g0LLRgdC10LPQtNCwINGF0L7RgtC10Lsg0L3QsNC/0LjRgdCw0YLRjCDRjdC80YPQu9GP0YLQvtGAIFNlZ2EgTWVnYSBEcml2ZSwg0L3QviDQtNGD0LzQsNGOINC90LUg0LTQvtGA0L7RgSDQtdGJ0LUpKQ=="
	
    message, err := idec.ParseMessage(txt)
	if err != nil {
		panic(err.Error())
	}
    
	fmt.Println(message.Body)
}
EOF

go build test.go
./test
Re: Несетевые проекты
```

# License

GNU GPL v3

# Authors

Denis Zheleztsov <difrex@lessmore.pw>
