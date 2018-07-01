package domain

type User struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Age       int                    `json:"age,omitempty"`
	Phones    []interface{}          `json:"phones,omitempty"`
	Relatives map[string]interface{} `json:"relatives,omitempty"`
}

func NewUser(id string, name string, age int, phones []interface{}, relatives map[string]interface{}) *User {
	return &User{id, name, age, phones, relatives}
}

func (u *User) GetPhones() []interface{} {
	return u.Phones
}

func (u *User) GetRelatives() map[string]interface{} {
	return u.Relatives
}

func (u *User) String() string {
	return u.Name
}

type UserService interface {
	Create(user *User) error
	Retrieve(userID string) (*User, error)
}
