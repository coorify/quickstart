package plugin

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func Database(o interface{}) gin.HandlerFunc {
	init := sync.Once{}
	return func(c *gin.Context) {
		init.Do(func() {
			// db := c.MustGet(field.SYS_DB).(*gorm.DB)
			// db.AutoMigrate()
		})
	}
}
