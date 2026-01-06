package simple

import "fmt"

type File struct {
	Name string
}

// membuat method pada file (struct) untuk menampilkan nama file yang ditutup
func (f *File) Close() {
	fmt.Println("Close file ", f.Name)
}

// membuat provider dengan menambahkan cleanup function (closure)
// closure dalam provider ini adalah return value berupa anonymous function
// yang nantinya digunakan sebagai penutup ketika file selesai dieksekusi
func NewFile(name string) (*File, func()) {
	file := &File{
		Name: name,
	}

	// penggunaan closure bisa seperti ini
	return file, func() {
		file.Close() // menutup file jika sudah selesai digunakan
	}
}