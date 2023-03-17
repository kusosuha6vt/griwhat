package main

import "fmt"

func main() {
	tagSystem := NewTagSystem()
	if err := tagSystem.AddTag("admin", 2); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	_ = tagSystem.AddTag("admin", -1)
	_ = tagSystem.AddTag("programmer", int(tagSystem.getTagByName("admin").Id))
	_ = tagSystem.AddTag("lawyer", int(tagSystem.getTagByName("programmer").Id))
	fmt.Printf("admin is a parent of lawer: %v\n",
		tagSystem.isParentOf(
			tagSystem.getTagByName("admin"),
			tagSystem.getTagByName("lawyer")))
	fmt.Printf("lawer is a parent of admin: %v\n",
		tagSystem.isParentOf(
			tagSystem.getTagByName("lawyer"),
			tagSystem.getTagByName("admin")))

}
