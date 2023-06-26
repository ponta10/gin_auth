// ./app/models/user.go
package models

import (
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

// JWTトークンの署名に使用されます。署名により、トークンが改ざんされていないことを確認できます
var JwtKey = []byte("your_secret_key")

// JWTトークンのペイロード部分に含まれるデータを表すものです
// jwt.StandardClaimsを埋め込んでおり、これによりJWTの標準的なクレーム（トークンの発行者、被験者、有効期限など）も持つことができます。
type Claims struct {
    Username string `json:"username"`
    UserID   uint   `json:"user_id"`
    jwt.StandardClaims
}

type User struct {
    gorm.Model
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Posts []Post `gorm:"foreignKey:UserID"`
}

func (user *User) HashPassword() error {
    // bcryptはハッシュ化のためのライブラリで、安全性の高いハッシュ値を生成します。
    // パスワードは通常、文字列として保管されますが、bcrypt.GenerateFromPassword関数はバイトスライスを引数に取ります。そのため、user.Password を []byte 型に変換しています
    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // ここでハッシュ化したパスワードをユーザーオブジェクトのパスワードフィールドに代入しています
    user.Password = string(hash)
    return nil
}

func (user *User) CheckPassword(password string) error {
    // ハッシュ化したパスワードとテーブルに保存されているパスワードをチェック
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// db.Create(user) の部分で新しいユーザーのレコードをデータベースに保存しています。.Error はこの操作中にエラーが発生した場合のエラー情報を取得するためのものです。
func CreateUser(db *gorm.DB, user *User) error {
    return db.Create(user).Error
}

// usernameが一致するユーザーを取得
func FindUserByUsername(db *gorm.DB, user *User) error {
    return db.Where("username = ?", user.Username).First(user).Error
}

func (user *User) GenerateToken() (string, error) {
    // トークンの有効期限を設定します。現在時刻から5分後に期限切れとなるようにします。
    expirationTime := time.Now().Add(5 * time.Minute)

    // クレームは、トークンの所有者についての情報（ここではユーザー名とユーザーID）と、このトークンが有効な期間（ここでは上記で設定したexpirationTime）を含むデータ構造
    claims := &Claims{
        Username: user.Username,
        UserID:   user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // これらのクレームを持つ新しいJWTを作成します。
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // token.SignedString(JwtKey)は、JwtKeyを用いてトークンに署名を施し、その署名付きトークンを文字列として返す
    return token.SignedString(JwtKey)
}
