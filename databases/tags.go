package databases

import (
	"fmt"
)

type Tag struct {
	Id        int64
	UniqueId  string
	Name     string
	UserId int64
	CreatedAt string
	UpdatedAt string
}

func AddTag(tag Tag) {
	fmt.Println("Adding Tag")
}

func GetTags() {

}