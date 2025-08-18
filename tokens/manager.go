package tokens

type TokenManager interface {
	GenerateAccessToken(userID int64) (string, error)
	GenerateRefreshToken(userID int64) (string, error)
	VerifyAccessToken(token string) (int64, error)
	VerifyRefreshToken(token string) (int64, error)
}
