package test

import (
	"fmt"
	utils "tech_tubbies/middleware/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("Mon January 2, 2006 3:4:5 pm")
	fmt.Println(formattedTime)

	require.Equal(t, formattedTime, "")
}

func TestMonthYear(t *testing.T) {
	fmt.Println(utils.GetYears())
}
