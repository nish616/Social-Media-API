package socialmediaapi

import (
	"github.com/gin-gonic/gin"

	"nishin/social-media-API/users"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/api")

	users.Register(v1.Group("/users"))
}
