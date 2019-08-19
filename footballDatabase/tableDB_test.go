package footballDatabase_test

import (
	"bot-git/footballDatabase"
	"bot-git/normalizedDate"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestListingReservation(t *testing.T) {
	const dbPath = "./test1.db"
	db := setup(dbPath)
	defer teardown(dbPath)

	res1 := testDate(14, 00)
	db.SetReservation("test", res1)

	res3 := testDate(14, 50)
	db.SetReservation("test", res3)

	res2 := testDate(14, 25)
	db.SetReservation("test", res2)

	reservations := db.GetAllReservationByStartTime(res1)

	expected := footballDatabase.TimeReservations{
		footballDatabase.NewTimeReservation("test", roundToMinute(res1)),
		footballDatabase.NewTimeReservation("test", roundToMinute(res2)),
		footballDatabase.NewTimeReservation("test", roundToMinute(res3)),
	}
	assert.Equal(t, expected, reservations)
}

func TestUserCanOnlyBookOneReservationADay(t *testing.T) {
	const dbPath = "./test2.db"
	db := setup(dbPath)
	defer teardown(dbPath)

	res1 := testDate(14, 00)
	db.SetReservation("test", res1)

	res2 := testDate(14, 25)

	result := db.IsBooked("test", res2)
	assert.True(t, result)
}

func TestIsFreeReturnsCorrectResult(t *testing.T) {
	const dbPath = "./test3.db"
	db := setup(dbPath)
	defer teardown(dbPath)

	res1 := testDate(14, 0)
	db.SetReservation("test", res1)

	res2 := testDate(14, 25)
	db.SetReservation("test", res2)

	res3 := testDate(15, 25)
	db.SetReservation("test", res3)

	check1 := testDate(14, 50)
	result1 := db.IsFree(check1)
	assert.True(t, result1)

	check2 := testDate(13, 34)
	result2 := db.IsFree(check2)
	assert.True(t, result2)

	check3 := testDate(15, 50)
	result3 := db.IsFree(check3)
	assert.True(t, result3)

	check4 := testDate(13, 50)
	result4 := db.IsFree(check4)
	assert.False(t, result4)

	check5 := testDate(14, 5)
	result5 := db.IsFree(check5)
	assert.False(t, result5)

	result6 := db.IsFree(res1)
	assert.False(t, result6)

	check7 := testDate(14, 50)
	result7 := db.IsFree(check7)
	assert.True(t, result7)
}

func TestFirstFreeTimeForReturnsTime(t *testing.T) {
	const dbPath = "./test4.db"
	db := setup(dbPath)
	defer teardown(dbPath)

	assertIsFree(t, db, 14, 00, 14, 00)

	res1 := testDate(14, 0)
	db.SetReservation("test", res1)

	assertIsFree(t, db, 14, 5, 14, 25)

	res2 := testDate(14, 25)
	db.SetReservation("test", res2)

	res3 := testDate(15, 25)
	db.SetReservation("test", res3)

	assertIsFree(t, db, 14, 10, 14, 50)

	assertIsFree(t, db, 14, 50, 14, 50)

	assertIsFree(t, db, 14, 55, 15, 50)
}

func assertIsFree(t *testing.T, db *footballDatabase.FootballDb, hCh, mCh, hExp, minExp int) {
	check := testDate(hCh, mCh)
	res := db.FirstFreeTimeFor(check)
	exp := testDate(hExp, minExp)
	assert.Equal(t, exp.Raw(), res)
	resFree := db.IsFree(normalizedDate.NewNormalizeDate(res))
	assert.True(t, resFree)
}

func roundToMinute(t normalizedDate.NormalizeDate) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

func setup(dbPath string) *footballDatabase.FootballDb {
	return footballDatabase.NewFootballDb(dbPath)
}
func teardown(dbPath string) {
	err := os.Remove(dbPath)
	if err != nil {
		panic(err)
	}
}

func testDate(hour, min int) normalizedDate.NormalizeDate {
	t := time.Date(2019, 8, 16, hour, min, 34, 134, time.UTC)
	return normalizedDate.NewNormalizeDate(t)
}
