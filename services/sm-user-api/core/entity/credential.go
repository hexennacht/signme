package entity

type Credential struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

type UserCredential struct {
	Private *Credential `json:"private,omitempty"`
	Public  *Credential `json:"public,omitempty"`
}
