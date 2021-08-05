package i18n

import "testing"

func TestAdd(t *testing.T) {
	i18nClient, err := NewI18n("./locale", "lang")

	if err != nil {
		t.Error("err")
		return
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候？")

	t.Log(str)

	str1 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")

	t.Log(str1)

}
