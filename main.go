package main

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Samuel struct {
	//UUID uuid.UUID `gorm:"default:uuid_generate_v4()`
	ID   string
	Name string
	Cars []Car
}

type Car struct {
	gorm.Model
	Plate    string
	SamuelID *string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Samuel) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4()
	s.ID = id.String()

	if s.Name == "ramon" {
		return errors.New("invalid role")
	}

	return
}

func main() {
	dsn := "host=localhost user=postgres password=test-db dbname=test port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Product{}, &User{}, &CreditCard{}, &Samuel{}, &Car{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "411111111111"},
	})

	db.Create(&Samuel{
		Name: "samueluco",
		Cars: nil,
	})

	Car1 := Car{
		Plate: "CagateLorito1",
	}
	Car2 := Car{
		Plate: "CagateLorito2",
	}

	db.Create(&Samuel{
		Name: "samueluquiño",
		Cars: []Car{Car1, Car2},
	})
}
