package test

import (
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductC struct {
	gorm.Model
	Code   string
	Price  uint
	Labels types.Strs
	Hash   []byte
}

func TestSQL(t *testing.T) {
	dsn := "host=localhost user=postgres password=dimogateway dbname=dimo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ProductC{})
	// Create
	db.Create(&ProductC{Code: "D42", Price: 100})

	var product ProductC
	db.Last(&product, "code = ?", "D42")

	db.Model(&product).Updates(ProductC{
		Labels: types.Strs{"abs", "cde"},
		Hash:   []byte("aaa"),
	})

	t.Log(product.ID, product.Code, product.Price, product.Labels)

	product.Labels = append(product.Labels, "edhg")

	db.Model(&product).Update("labels", product.Labels)

	t.Log(product.ID, product.Code, product.Price, product.Labels, product.Hash)

	t.Fatal()
}

func TestGORM(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("/tmp/test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&ProductC{})

	// Create
	db.Create(&ProductC{Code: "D42", Price: 100})

	//datatypes.JSONSet("labels").Set("age", 20)

	// Read
	var product ProductC
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	t.Log(product.ID, product.Code, product.Price)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200).UpdateColumn("labels", datatypes.JSONSet("labels").Set("age", 20))
	db.Model(&product).UpdateColumn("labels", datatypes.JSONSet("labels").Set("sex", "male"))

	t.Log(product.ID, product.Code, product.Price, product.Labels)
	t.Log(product)

	res := datatypes.JSONQuery("labels").Extract("age")
	t.Log(res)
	// Update - 更新多个字段
	db.Model(&product).Updates(ProductC{Price: 300, Code: "F42"}) // 仅更新非零值字段
	t.Log(product.ID, product.Code, product.Price)
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "E42"})
	t.Log(product.ID, product.Code, product.Price)

	// Delete - 删除 product
	db.Delete(&product, 1)

	t.Log(product.ID, product.Code, product.Price)

	t.Fatal("")
}
