package model

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gorilla/securecookie"
	"github.com/tinyci/ci-agents/errors"
)

func (m *Model) createToken(name string) ([]byte, string) {
	key := securecookie.GenerateRandomKey(64)
	encoded := fmt.Sprintf("%s %s", name, string(key))
	return key, base64.URLEncoding.EncodeToString([]byte(encoded))
}

func (m *Model) unpackToken(token string) (string, []byte, *errors.Error) {
	b, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", nil, errors.ErrInvalidAuth.Wrap(err)
	}

	parts := strings.SplitN(string(b), " ", 2)
	if len(parts) != 2 {
		return "", nil, errors.ErrInvalidAuth.Wrap("invalid token")
	}

	return parts[0], []byte(parts[1]), nil
}

// ValidateToken checks that a token is valid for a given user.
func (m *Model) ValidateToken(token string) (*User, *errors.Error) {
	name, authToken, err := m.unpackToken(token)
	if err != nil {
		return nil, err // except for in this case, where it's done already above.
	}

	u, err := m.FindUserByName(name)
	if err != nil {
		return nil, errors.ErrInvalidAuth
	}

	if len(u.LoginToken) == 0 {
		return nil, errors.ErrInvalidAuth
	}

	if bytes.Equal(u.LoginToken, authToken) {
		return u, nil
	}

	return nil, errors.ErrInvalidAuth
}

// GetToken retrieves a new token for logging in. If one exists, the
// DeleteToken method must be called first; otherwise this routine will throw
// an error.
func (m *Model) GetToken(name string) (string, *errors.Error) {
	u, err := m.FindUserByName(name)
	if err != nil {
		return "", err
	}

	if len(u.LoginToken) != 0 {
		return "", errors.New("Login token already exists, must delete the old one first")
	}

	key, token := m.createToken(name)
	u.LoginToken = key

	if err := m.Save(u).Error; err != nil {
		return "", errors.New(err)
	}

	return token, nil
}

// DeleteToken removes the existing token.
func (m *Model) DeleteToken(name string) *errors.Error {
	u, err := m.FindUserByName(name)
	if err != nil {
		return err
	}

	u.LoginToken = nil
	if err := m.Save(u).Error; err != nil {
		return errors.New(err)
	}

	return nil
}
