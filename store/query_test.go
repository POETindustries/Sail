package store

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sail/conf"
	"testing"
	"time"
)

const iconData = "89504e470d0a1a0a0000000d4948445200000020000000200802000000fc18eda300000337494441" +
	"5478daed96594f135114c7fb197cf14df808bef9e08bf1dd18138c11a2a2b209252c760345a48226" +
	"106d40239b5d289402a5ed944ee93e5d8012a1539829d0d24e2b4596820410b15512a2d316a1b425" +
	"e196f8a493939bb9ff4cefef9ef33fbdb9945f7ff9a1fc1b80edb5a0ca64e11bad1cb591a3465ab5" +
	"268969c4e3c4cf04d80cae924bac2e2dd5429adb7a3bc3b7c9742ed2a7fc6450ed442164c8921ae8" +
	"034a93c5c235589c3333600064fc639e15cf55db6e8a9412ab2d30ef76381c2ccd08cde12bd58dfb" +
	"bddec8477b3fd9b0a1d0307983d39167c6fc1e0f00a0727038e3d2e55b1dbdc2d1092e8f17138381" +
	"8512d40f59c70a0a0a628adfed2ac302576935cf9743529d0100d02414658be1a7be0d91c14ca150" +
	"a6312ca667ab460552282333331c0e93d3cf84a7d4e12fd28de736b78576be8279c0e814d2a77cb4" +
	"a1a37d09b4866b8defb2d88dfb3f22ab87b7b76a54086909c3b56a1c190336b9b44f414309261e28" +
	"52189b11db63859edd27bc70fedcf52b178b65aafc41cd7dfd047dfa13e90a035fe4ebccc080bc1e" +
	"39f9639a83888d9573419b80390135799136b3b89e3ebb12d5234162dec03a60407e04401c3150a2" +
	"0c4266709cf0f9a9f80a393d0490d1a0d00003a8bd507c06e45832eaa44a8699722d1d5b885f9d86" +
	"faeae4c3c0806a097c6c9be89f9114d1f86944a995ab8101af146ab2435200e2950391a8850001bb" +
	"e1bd6abe98116d92c415d124989da8e81e585c0a9e16e076bb1c3276bfa687115febf8bd27248412" +
	"0d4af992f935799c9c0ab0beb129e13595b7f2539728c180e8cb3d5ebf6690bbf32d04e0411da43e" +
	"71cb4966e476cb804da60f24755142f5e36039422930a044ac38964162e71c5372bac00145222845" +
	"891ca992407d0f84e0252a8f1e76a97b34a94de97d1030e0a55c75ca3f1a030bb42ab5c000c9105c" +
	"e5593fd180b8541ed95c386a4fe75671a7bde72809d497d273d6dc32a35d90e6b50586a4d47e5995" +
	"3b9870381f06d3bdf6a4a3c58668d201b8e6bdf850fdeee4fb8a17f5773b450f1194e5fdc27207c9" +
	"60ceaf67cbf4851f446dfcae80f9ad8453bc9f06e07b28cc6d7966428c071788b95995522955c272" +
	"25ac57c15b6b07471b86610aa9e8ffddf42ccf6f7af985bd894290440000000049454e44ae426082"

type TestData struct {
	id             uint64
	name           string
	content        sql.NullString
	birthday       time.Time
	active         bool
	progress       float32
	progressDetail float64
	tries          uint8
	avatar         []byte
}

func (t *TestData) String() string {
	return fmt.Sprintf("{%d '%s' '%s' '%v' %t %f %f %d}",
		t.id, t.name, t.content.String, t.birthday, t.active, t.progress,
		t.progressDetail, t.tries)
}

func (t *TestData) Equals(other *TestData) bool {
	return t.name == other.name &&
		t.content.String == other.content.String &&
		t.birthday.Equal(other.birthday) &&
		t.active == other.active &&
		t.progress == other.progress &&
		t.progressDetail == other.progressDetail &&
		t.tries == other.tries &&
		bytes.Equal(t.avatar, other.avatar)
}

var testdata = TestData{
	1, "Voldemort", sql.NullString{String: "I am the Greatest!"}, time.Date(1926, 12, 31, 0, 0, 0, 0, time.Local), true,
	8990.0, 8990.00000112, 5, []byte{}}

func TestQuery(t *testing.T) {
	conf.Instance().DBName = "sl_test"
	pic, _ := hex.DecodeString(iconData)
	testdata.avatar = pic

	for _, d := range []string{"sqlite3", "mysql", "postgres"} {
		conf.Instance().DBDriver = d
		if err := DB().init(); err != nil {
			t.Errorf("%s init failed: %s", d, err.Error())
			continue
		}

		testTableCreate(t)
		testQueryAdd(t)
		testQueryGet(t)
		//testQueryUpdate(t)
		testQueryDelete(t)

	}
}

func testTableCreate(t *testing.T) {
	data := []*SetupData{
		{Name: "test_id", Value: 1, IsPrimary: true},
		{Name: "test_name", Value: "", Size: Small},
		{Name: "test_content", Value: "", Size: All},
		{Name: "test_birthday", Value: time.Now()},
		{Name: "test_active", Value: true},
		{Name: "test_progress", Value: float32(0.0)},
		{Name: "test_progress_detailed", Value: float64(0.0)},
		{Name: "test_num_tries", Value: uint8(0)},
		{Name: "test_avatar", Value: []byte{}, Size: All}}
	DB().Setup("test", data)
}

func testQueryAdd(t *testing.T) {
	valsGood := map[string]interface{}{
		"test_name":              testdata.name,
		"test_content":           testdata.content.String,
		"test_birthday":          testdata.birthday.Unix(),
		"test_active":            testdata.active,
		"test_progress":          testdata.progress,
		"test_progress_detailed": testdata.progressDetail,
		"test_num_tries":         testdata.tries,
		"test_avatar":            testdata.avatar}
	valsBad := map[string]interface{}{
		"test_name":     12,
		"test_birthday": "hello"}

	if _, ok := Add().In("test").Values(valsGood).Exec(); !ok {
		t.Errorf("%s: error in Add with data: %+v", conf.Instance().DBDriver, valsGood)
	}
	if _, ok := Add().In("test").Values(valsBad).Exec(); ok {
		t.Errorf("%s: no error happened, but should have: %+v", conf.Instance().DBDriver, valsBad)
	}
}

func testQueryGet(t *testing.T) {
	if r, ok := Get().In("test").All().Exec(); ok {
		getHelper(r, t)
	}
	if r, ok := Get().In("test").Equals("test_name", "Voldemort").Exec(); ok {
		getHelper(r, t)
	}
}

func getHelper(r *sql.Rows, t *testing.T) {
	defer r.Close()
	for r.Next() {
		var data TestData
		var birthdayTimestamp int64

		if err := r.Scan(&data.id, &data.name, &data.content, &birthdayTimestamp,
			&data.active, &data.progress, &data.progressDetail, &data.tries, &data.avatar); err != nil {
			t.Error(conf.Instance().DBDriver + ": scan failed - " + err.Error())
		}
		data.birthday = time.Unix(birthdayTimestamp, 0)
		fmt.Printf("%s: %s\n", conf.Instance().DBDriver, &data)

		if !data.Equals(&testdata) {
			t.Errorf("%s: Data not equal after scan: %s %s", conf.Instance().DBDriver, data.String(), testdata.String())
		} else {
			fmt.Printf("%s: Storage and retreival successful.\n", conf.Instance().DBDriver)
		}
	}
}

func testQueryUpdate(t *testing.T) {
	Update().In("test").Exec()
}

func testQueryDelete(t *testing.T) {
	Delete().In("test").All().Exec()
}
