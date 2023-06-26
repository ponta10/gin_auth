// ./app/middlewares/authMiddleware.go
package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"gin_auth/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		// Authorizationヘッダーが存在するかチェックします。存在しない場合は、エラーメッセージを返しリクエストを中断します。
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		// Authorizationヘッダーの値が"Bearer "で始まる形式であるかチェックします。
		if !strings.Contains(header, BearerSchema) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// ヘッダーから"Bearer "を取り除き、JWTトークン文字列を取得します
		tokenString := strings.TrimPrefix(header, BearerSchema)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims := &models.Claims{}

		// jwt.ParseWithClaims関数を用いて、JWTトークンのペイロード部を解析し、その情報をclaimsオブジェクトに格納します。このとき、JWTの署名が正当であることを検証します。
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})

		if err != nil || !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 解析が成功したら、ペイロード（クレーム）から取得したユーザーIDを、その後の処理（ハンドラー関数）で使えるようにGinのコンテキストにセットします
		c.Set("userId", claims.UserID)

		// 最後にc.Next()を呼び出しています。これにより次のハンドラー関数（ミドルウェアも含む）への制御が移されます。
		c.Next()
	}
}
