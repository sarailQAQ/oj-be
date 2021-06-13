
/**
 * @author: sarail
 * @time: 2021/6/13 15:57
**/

package api

import (
	"github.com/gin-gonic/gin"
	"log"
)


func Init() {
	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
