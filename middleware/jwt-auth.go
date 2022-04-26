package middleware

import (
	"log"
	"net/http"
	"ustoj-master/service"
	"ustoj-master/utils"
	"ustoj-master/vo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT valudates the token user given,return 401 if not valid
func AuthorizenJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req vo.LoginRequest
		var loginobject vo.LoginResponse
		code := vo.OK
		defer func() {
			resp := vo.LoginResponse{
				Code: code,
				Data: loginobject.Data,
			}
			c.JSON(http.StatusOK, resp)
			utils.LogReqRespBody(req, resp, "XXXXXXXXXXX")
		}()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			//response := vo.UnknownError
			code = vo.UnknownError
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]:", claims["user_id"])
			log.Println("Claim[issuer]:", claims["issuer"])
		} else {
			log.Print(err)
			code = vo.UnknownError
			//c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}
	}
}
