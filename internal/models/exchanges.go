package models

type Exchanges struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"name"`
	ApiKey        []byte `json:"api_key"`
	APISecret     []byte `json:"api_secret"`
	APIPassPhrase []byte `json:"api_pass"`
}
