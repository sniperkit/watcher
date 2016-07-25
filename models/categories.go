package models

type Categories struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Parent     string `json:"parent,omitempty"`
	MainParent string `json:"main_parent,omitempty"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
}

func Extract() {

}
