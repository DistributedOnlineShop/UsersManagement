package util

import (
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateText() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.Word(),
		Valid:  true,
	}
}

func GenerateDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateInt32() int32 {
	return rand.Int32N(1000) + 1
}

func GenerateBool() pgtype.Bool {
	return pgtype.Bool{
		Bool:  gofakeit.Bool(),
		Valid: true,
	}
}

func GenerateSessionStatus() string {
	statuses := []string{
		"LOGIN",
		"SIGNUP",
		"ADMIN",
	}
	return statuses[rand.IntN(len(statuses))]
}

func GenerateUserStatus() string {
	statuses := []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DISABLED",
	}
	return statuses[rand.IntN(len(statuses))]
}

func GenerateUserRole() string {
	statuses := []string{
		"USER",
		"ADMIN",
	}
	return statuses[rand.IntN(len(statuses))]
}
