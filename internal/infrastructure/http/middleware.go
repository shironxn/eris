package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/util"
)

type Middleware struct {
	JWT *util.JWT
}

func NewMiddleware(jwt *util.JWT) *Middleware {
	return &Middleware{
		JWT: jwt,
	}
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh-token")
		if err != nil {
			view.JSON(ctx, http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}

		claims, err := m.JWT.ValidateToken(refreshToken, m.JWT.Refresh)
		if err != nil {
			view.JSON(ctx, http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}

		accessToken, err := ctx.Cookie("access-token")
		if err != nil || accessToken == "" {
			accessToken, err := m.JWT.GenerateAccessToken(claims.UserID)
			if err != nil {
				view.JSON(ctx, http.StatusUnauthorized, nil)
				ctx.Abort()
				return
			}

			_, err = m.JWT.ValidateToken(accessToken, m.JWT.Access)
			if err != nil {
				view.JSON(ctx, http.StatusUnauthorized, nil)
				ctx.Abort()
				return
			}

			ctx.SetCookie("access-token", accessToken, int(time.Now().Add(10*time.Minute).Unix()), "/", "localhost", false, true)
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
