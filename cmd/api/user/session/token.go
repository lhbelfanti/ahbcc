package session

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// CreateToken creates a new session token for the user login action
type CreateToken func(ctx context.Context, userID int) (string, time.Time, error)

// MakeCreateToken creates a new CreateToken function
func MakeCreateToken(insertUserSession Insert) CreateToken {
	return func(ctx context.Context, userID int) (string, time.Time, error) {
		expiresAt := time.Now().Add(30 * 24 * time.Hour) // 30 days session expiry
		payload := fmt.Sprintf("%d:%d", userID, expiresAt.Unix())

		mac := hmac.New(sha256.New, []byte(os.Getenv("SESSION_SECRET_KEY")))
		mac.Write([]byte(payload))
		signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))
		token := base64.URLEncoding.EncodeToString([]byte(payload)) + "." + signature

		session := DAO{
			UserID:    userID,
			Token:     token,
			ExpiresAt: expiresAt,
		}

		err := insertUserSession(ctx, session)
		if err != nil {
			log.Error(ctx, err.Error())
			return "", time.Time{}, FailedToCreatUserSessionToken
		}

		return token, expiresAt, nil
	}
}
