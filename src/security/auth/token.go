 package auth

 import (
	 "fmt"
	 "github.com/dgrijalva/jwt-go"
	 "github.com/twinj/uuid"
	 "net/http"
	 "os"
	 "strconv"
	 "strings"
	 "time"
	 "web-shop/domain/auth"
	 "web-shop/usecase"
 )



type Token struct {
	RedisUsecase                 usecase.RedisUsecase
}

func NewToken(redisUsecase usecase.RedisUsecase) *Token {
	return &Token{
		RedisUsecase : redisUsecase,
	}
}
func NewToken2() *Token {
	return &Token{nil}
}
type TokenInterface interface {
	CreateToken(userid uint64) (*auth.TokenDetails, error)
	ExtractTokenMetadata(r *http.Request) (*auth.AccessDetails, error)
}


func (t *Token) CreateToken(userid uint64) (*auth.TokenDetails, error) {
	td := &auth.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()

	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	td.RefreshUuid = td.TokenUuid + "++" + strconv.Itoa(int(userid))

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*auth.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &auth.AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
		}, nil
	}
	return nil, err
}



 func ExtractUserId(r *http.Request) (string, error) {
	 tokenString := ExtractToken(r)

	 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		 return []byte(os.Getenv("ACCESS_SECRET")), nil
	 })

	 claims, ok := token.Claims.(jwt.MapClaims)

	 if ok  {
		 userId, ok := claims["user_id"].(string)
		 if !ok {
			 return "", err
		 }

		 return userId, nil
	 }
	 return "", err
 }
