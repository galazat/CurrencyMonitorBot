package currency

import (
	"fmt"
	"time"
)

func getUrl() string {
	tm := time.Now()

	return fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%v", tm.Format("02/01/2006"))
}
