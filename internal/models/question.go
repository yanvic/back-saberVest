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
	Id          int     `json:"id"`
	Description *string `json:"description"`
	Number      int     `json:"number"`
	Question    string  `json:"question"`
	Alternative string  `json:"alternative"`
	Response    int     `json:"response"`
	ResponseId  string  `json:"response_id"`
	TextId      *int    `json:"text_id"`
	ImageId     *int    `json:"image_id"`
	TopicId     int     `json:"topic_id"`
	NoteId      *int    `json:"note_id"`
}

type FilterQuestions struct {
	Topic       string `json:"topic"`
	Matter      string `json:"matter"`
	InitialDate string `json:"date_initial"`
	FinalDate   string `json:"final_date"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

func (c *Test) AllQuestions() ([]Test, error) {
	rows, err := DB.Query("SELECT * FROM test_uece")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []Test

	for rows.Next() {
		var test Test
		err := rows.Scan(&test.Id, &test.Description, &test.Number, &test.Question, &test.Alternative,
			&test.Response, &test.TextId, &test.ImageId, &test.TopicId, &test.NoteId, &test.ResponseId)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}

func (c *Test) QuestionsParams(filters FilterQuestions) ([]Test, error) {
	rows, err := DB.Query(`SELECT t.id, t.description, t.number, t.statement, t.alternative, t.response_id, t.text_id, t.image_id, 
       t.topic_id, t.note_id, t.response, tx.description, img.image, nt.description
	   FROM public.test_uece
	   LEFT JOIN image img ON (img.id = t.image_id)
	   LEFT JOIN text tx ON (tx.id = t.text_id)
	   LEFT JOIN note nt ON (nt.id = t.note_id)
	   `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []Test

	for rows.Next() {
		var test Test
		err := rows.Scan(&test.Id, &test.Description, &test.Number, &test.Question, &test.Alternative,
			&test.Response, &test.TextId, &test.ImageId, &test.TopicId, &test.NoteId, &test.ResponseId)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}
