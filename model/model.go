// package model

// type User struct {
// 	Id         string `json:"id"`
// 	Name       string `json:"name"`
// 	Email      string `json:"email"`
// 	Age        string `json:"age"`
// 	Created_at string `json:"created_at"`
// }


package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
