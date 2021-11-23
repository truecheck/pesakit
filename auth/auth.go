package auth

type (
	Key struct {
        ID       string `json:"id"`
        Key      string `json:"key"`
        Password string `json:"password"`
    }
)
