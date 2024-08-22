package models

type Matter struct {
	Id          int `json:"id"`
	Description int `json:"description"`
}

type Topic struct {
	Id          int `json:"id"`
	Description int `json:"description"`
	MatterId    int `json:"matter_id"`
}

type Note struct {
	Id          int `json:"id"`
	Description int `json:"description"`
}

type Text struct {
	Id          int `json:"id"`
	Description int `json:"description"`
}

type Response struct {
	Id          int `json:"id"`
	Description int `json:"description"`
}

type Test struct {
	Id          int    `json:"id"`
	Description int    `json:"description"`
	Number      int    `json:"number"`
	Question    string `json:"question"`
	Alternative string `json:"alternative"`
	Response    int    `json:"response"`
	TextId      int    `json:"text_id"`
	ImageId     int    `json:"image_id"`
	TopicId     int    `json:"topic_id"`
	NoteId      int    `json:"note_id"`
}
