## 使用

```
go get -u github.com/landryXu/i18n
```

```
package main

import (
	"github.com/landryXu/i18n"
	"log"
)

func main() {
	i18nClient, err := i18n.NewI18n("./locale", "lang")
	if err != nil {
		log.Fatalln(err)
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候？")
	log.Println(str)

	str1 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")
	log.Println(str1)

}
```