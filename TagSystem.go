package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Tag struct {
	Id       uint
	Uuid     uuid.UUID
	Name     string
	ParentId int
}

type TagSystem struct {
	lastId     uint
	tagsById   map[uint]*Tag
	tagsByName map[string]*Tag
	tagsPool   []Tag
}

func NewTagSystem() TagSystem {
	return TagSystem{0, make(map[uint]*Tag), make(map[string]*Tag), make([]Tag, 0)}
}

func (tagSystem *TagSystem) AddTag(name string, parent int) (err error) {
	if parent < -1 || parent >= len(tagSystem.tagsPool) {
		return fmt.Errorf("parent is not valid: %v", parent)
	}
	// -1 => no parent
	id := tagSystem.lastId
	tagSystem.lastId += 1
	tagSystem.tagsPool = append(tagSystem.tagsPool, Tag{id, uuid.New(), name, parent})
	tagSystem.tagsByName[name] = &tagSystem.tagsPool[len(tagSystem.tagsPool)-1]
	tagSystem.tagsById[id] = &tagSystem.tagsPool[len(tagSystem.tagsPool)-1]
	return nil
}

func (tagSystem *TagSystem) getTagById(id uint) *Tag {
	return tagSystem.tagsById[id]
}

func (tagSystem *TagSystem) getTagByName(name string) *Tag {
	return tagSystem.tagsByName[name]
}

func (tagSystem *TagSystem) isParentOf(a, b *Tag) bool {
	if a == nil || b == nil {
		return false
	}
	// is a parent of b?
	for {
		if a.Id == b.Id {
			return true
		}
		if b.ParentId >= 0 {
			b = tagSystem.getTagById(uint(b.ParentId))
		} else {
			break
		}
	}
	return false
}

func (tagSystem *TagSystem) deleteTag(tag *Tag) {
	if tag == nil {
		return
	}
	newParent := -1
	if tag.ParentId != -1 {
		newParent = int(tagSystem.getTagById(uint(tag.ParentId)).Id)
	}
	for _, element := range tagSystem.tagsPool {
		if element.ParentId == int(tag.Id) {
			element.ParentId = newParent
		}
	}
	delete(tagSystem.tagsById, tag.Id)
	delete(tagSystem.tagsByName, tag.Name)
}
