package main

import (
	"fmt"
	"sort"
	"strings"
)

type FileSystem struct {
	Children    map[string]*FileSystem
	Content     string
	FileName    string
	IsDirectory bool
}

func Constructor() FileSystem {
	return *NewFileSystem("")
}

func NewFileSystem(fileName string) *FileSystem {
	return &FileSystem{make(map[string]*FileSystem), "", fileName, true}
}

func (self *FileSystem) Ls(path string) []string {
	fs := self.getOrCreateDeepFiler(path)
	children := make([]string, 0)
	if fs.IsDirectory {
		for k, _ := range fs.Children {
			children = append(children, k)
		}
	} else {
		children = append(children, fs.FileName)
	}

	sort.Sort(ByLex(children))
	return children
}

func (self *FileSystem) Mkdir(path string) {
	self.getOrCreateDeepFiler(path)
}

func (self *FileSystem) AddContentToFile(filePath string, content string) {
	fs := self.getOrCreateDeepFiler(filePath)
	fs.IsDirectory = false
	fs.Content = fs.Content + content
}

func (self *FileSystem) ReadContentFromFile(filePath string) string {
	fs := self.getOrCreateDeepFiler(filePath)
	return fs.Content
}

type ByLex []string

func (self ByLex) Len() int           { return len(self) }
func (self ByLex) Less(i, j int) bool { return self[i] < self[j] }
func (self ByLex) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func toFileNames(s string) []string {
	if s[len(s)-1] != '/' {
		s = fmt.Sprintf("%s/", s)
	}
	fileNames := strings.Split(s, "/")
	if len(fileNames) > 1 {
		fileNames = fileNames[:len(fileNames)-1]
	}
	return fileNames
}

func (self *FileSystem) getOrCreateDeepFiler(path string) *FileSystem {
	fileNames := toFileNames(path)
	current := self
	for i := 1; i < len(fileNames); i++ {
		current = current.getOrCreateChildFiler(fileNames[i])
	}
	return current
}

func (self *FileSystem) getOrCreateChildFiler(fileName string) *FileSystem {
	if v, ok := self.Children[fileName]; ok {
		return v
	}
	self.Children[fileName] = NewFileSystem(fileName)
	return self.Children[fileName]
}
