package middleware

import (
	"net/http"
	"ustoj-master/common"
	"ustoj-master/service"
	"ustoj-master/vo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var logger = common.LogInstance()

//AuthorizeJWT valudates the token user given,return 401 if not valid
func AuthorizenJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var req vo.LoginRequest
		var loginobject vo.LoginResponse
		code := vo.OK
		defer func() {
			//resp := vo.LoginResponse{
			//	Code: code,
			//	Data: loginobject.Data,
			//}
			//c.JSON(http.StatusOK, resp)
			//utils.LogReqRespBody(req, resp, "XXXXXXXXXXX")
		}()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			//response := vo.UnknownError
			code = vo.UnknownError
			resp := vo.LoginResponse{
				Code: code,
				Data: loginobject.Data,
			}
			c.AbortWithStatusJSON(http.StatusForbidden, resp)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			//logger.Println("Claim[Username]:", claims["Username"])
			logger.Println("Claim[issuer]:", claims["issuer"])
		} else {
			logger.Errorln(err)
			code = vo.UnknownError
			resp := vo.LoginResponse{
				Code: code,
				Data: loginobject.Data,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
	}
}
