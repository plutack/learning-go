wpackage types
import "fmt"


type animal string
// type noise string

func (name animal) makeNoise()  {
	fmt.Printf("%s makes a sound", name)
}

func Run(){
	doga := animal("dog")
	doga.makeNoise()
}
 
type User struct {
	id   int
	name string
	age  int
}

// Use a pointer receiver to modify the original User instance
func (u *User) addAge() {
	u.age++
}

func Main() {
	user := User{id: 1, name: "John", age: 25}
	fmt.Printf("Before: %+v\n", user)

	user.addAge() // This will increment the age

	fmt.Printf("After: %+v\n", user) // Age will be incremented
}