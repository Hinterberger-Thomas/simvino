package user

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/Hinterberger-Thomas/simvino/db"
)

var ctx = context.Background()

func InsertUserToken(token string, userName string) {
	var resToken []UserToken
	checksum := sha256.Sum224([]byte(token))
	if db.Client.Get(ctx, userName).Scan(&resToken); resToken != nil {
		resToken = append(resToken, UserToken{checksum: string(checksum[:]), Token: token})
		err := db.Client.Set(ctx, userName, resToken, -1).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	resToken = append(resToken, UserToken{checksum: string(checksum[:]), Token: token})
	db.Client.Set(ctx, userName, resToken, -1)
}

func TokenExist(userName string, tokenChecksum string) bool {
	var resToken []UserToken
	if db.Client.Get(ctx, userName).Scan(&resToken); resToken != nil {
		for _, value := range resToken {
			if tokenChecksum == value.checksum {
				return true
			}
		}
	}
	return false
}
