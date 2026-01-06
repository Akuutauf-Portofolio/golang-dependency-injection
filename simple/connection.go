package simple

import "fmt"

type Connection struct {
	*File
}

// membuat method pada connection (struct)
func (c *Connection) Close() {
	fmt.Println("Close connection ", c.File.Name)
}

// membuat provider
// yang mengimplementasikan closure juga
func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{
		File: file,
	}

	return connection, func ()  {
		connection.Close()
	}
}