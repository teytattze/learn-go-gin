package models

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var postsData = []Post{
	{Id: 1, Title: "Post Title 1", Content: "Content 1", Author: "Tey"},
	{Id: 2, Title: "Post Title 2", Content: "Content 2", Author: "Tey"},
	{Id: 3, Title: "Post Title 3", Content: "Content 3", Author: "Tey"},
}

func GetAllPosts() []Post {
	return postsData
}

func GetPostById(id int) Post {
	for _, v := range postsData {
		if id == v.Id {
			return v
		}
	}
	return Post{}
}
