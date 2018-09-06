package str

import (
	"testing"
)

func TestGetFirstLastName(t *testing.T) {
	firstName, lastName := GetFirstLastName("")
	t.Logf("\n firstName: %s \nlastName: %s\n", firstName, lastName)

	firstName, lastName = GetFirstLastName("John")
	t.Logf("\n firstName: %s \nlastName: %s\n", firstName, lastName)

	firstName, lastName = GetFirstLastName("John Snow")
	t.Logf("\n firstName: %s \nlastName: %s\n", firstName, lastName)

	firstName, lastName = GetFirstLastName("   John       Snow     ")
	t.Logf("\n firstName: %s \nlastName: %s\n", firstName, lastName)
}
