package pass

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var defaultArgon2Config = argon2Config{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func Argon2Hash(password string) (string, error) {
	salt := make([]byte, defaultArgon2Config.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultArgon2Config.Iterations,
		defaultArgon2Config.Memory,
		defaultArgon2Config.Parallelism,
		defaultArgon2Config.KeyLength,
	)
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultArgon2Config.Memory,
		defaultArgon2Config.Iterations,
		defaultArgon2Config.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return encodedHash, nil
}

const passPartsCount = 6
const (
	_ = iota
	passArgonVariant
	passVersion
	passParameters
	passSalt
	passHash
)

func Argon2Verify(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != passPartsCount {
		return false, errors.New("invalid hash format")
	}
	if parts[passArgonVariant] != "argon2id" {
		return false, errors.New("unsupported argon variant")
	}
	var version int
	if _, err := fmt.Sscanf(parts[passVersion], "v=%d", &version); err != nil {
		return false, fmt.Errorf("failed to parse argon version: %w", err)
	}
	if version != argon2.Version {
		return false, errors.New("unsupported argon version")
	}
	var memory, iterations uint32
	var parallelism uint8
	if _, err := fmt.Sscanf(parts[passParameters], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism); err != nil {
		return false, fmt.Errorf("failed to parse argon parameters: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[passSalt])
	if err != nil {
		return false, fmt.Errorf("failed to parse salt: %w", err)
	}
	storedHash, err := base64.RawStdEncoding.DecodeString(parts[passHash])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultArgon2Config.Iterations,
		defaultArgon2Config.Memory,
		defaultArgon2Config.Parallelism,
		defaultArgon2Config.KeyLength,
	)
	if subtle.ConstantTimeCompare(storedHash, computedHash) != 1 {
		return false, nil
	}

	return true, nil
}
