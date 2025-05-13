package gin

import "github.com/gin-gonic/gin"

func OpenHost() {

	router := gin.Default()

	router.Run(":8080")

}
