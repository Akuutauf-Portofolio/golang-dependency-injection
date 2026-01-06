package simple

// membuat struct database
type Database struct {
	Name string
}

// untuk menangani multiple binding pada injector, bisa menggunakan alias
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

// membuat constructor untuk database postgree sql
// sehingga pada tipe data yang sama yaitu (*Database), nanti akan diganti menjadi alias
// seperti return value pada construct milik database postgree sql dan mongo db
func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	// return value akan memaksa agar dikonversi ke bentuk database
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreeSQL"})
}

// membuat constructor untuk database mongo db
func NewDatabaseMongoDB() *DatabaseMongoDB {
	// return value akan memaksa agar dikonversi ke bentuk database
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

// membuat repository database, tanpa alias (contoh)
// akan membingungkan
type DatabaseRepository struct {
	// DatabasePostgreSQL *Database # sebelum menggunakan alias
	// DatabaseMongoDB *Database # sebelum menggunakan alias
	DatabasePostgreSQL *DatabasePostgreSQL // setelah menggunakan alias
	DatabaseMongoDB *DatabaseMongoDB // setelah menggunakan alias
}

// membuat constructor untuk database repository
// kalau ada parameter dengan tipe data yang sama, maka injector tidak mendukung (mulitple binding)
func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMongoDB: mongoDB,
	}
}