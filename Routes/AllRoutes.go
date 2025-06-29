package Routes

import (
	"github.com/gin-gonic/gin"
)

func AllRouteInit(r *gin.Engine) {
	InitAnnRoute(r)
	InitUserRoute(r)
	//AdminRouteInit(r)
	InitOrderRoute(r)
	InitVenueRoutes(r)
}
