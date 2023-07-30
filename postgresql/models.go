package postgresql

var onboardedEntitiesMap = map[string]interface{}{
	"Account": Account{},
}

// func getEntityObj(obj interface{}) {

// 	switch v := obj.(type) {
// 	case Account:
// 		p := v
// 		fmt.Println("Created Account object:", p)

// 	default:
// 		fmt.Println("Unknown object type")
// 	}
// }

type Account struct {
	ID       int    `json:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
