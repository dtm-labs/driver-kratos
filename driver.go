package driver_kratos

import (
	"github.com/dtm-labs/dtmdriver"
	_ "github.com/dtm-labs/dtmdriver-kratos"
	driver "github.com/dtm-labs/dtmdriver-kratos"
)

func init() {
	if err := dtmdriver.Use(driver.DriverName); err != nil {
		panic(err)
	}
}
