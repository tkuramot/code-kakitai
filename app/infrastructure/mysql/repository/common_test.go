package repository

import (
	"testing"

	"gopkg.in/testfixtures.v2"

	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

var (
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error

	// DBの立ち上げ
	resource, pool := db_test.CreateContainer()
	defer db_test.CloseContainer(resource, pool)

	// DBへ接続する
	dbCon := db_test.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBをセットアップ
	db_test.SetupTestDB()

	// テストデータの準備
	fixturePath := "../fixtures"
	fixtures, err = testfixtures.NewFolder(dbCon, &testfixtures.MySQL{}, fixturePath)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetQuery(q)
	db.SetDB(dbCon)

	// テスト実行
	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
