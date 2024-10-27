package models

import (
	"fmt"
	"github.com/lib/pq"
	"time"
)

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
	Id           int       `json:"id"`
	Description  *string   `json:"description"`
	Number       *int      `json:"number"`
	Question     *string   `json:"question"`
	Alternative  *string   `json:"alternative"`
	Response     *string   `json:"response"`
	ResponseId   *int      `json:"response_id"`
	TextId       *int      `json:"text_id"`
	ImageId      *int      `json:"image_id"`
	TopicId      *int      `json:"topic_id"`
	NoteId       *int      `json:"note_id"`
	Alternatives []string  `json:"alternatives"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FilterQuestions struct {
	Topic       int    `json:"topic"`
	University  int    `json:"university"`
	Matter      int    `json:"matter"`
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
			&test.Response, &test.TextId, &test.ImageId, &test.TopicId, &test.NoteId, &test.ResponseId,
			pq.Array(&test.Alternatives), &test.Date, &test.CreatedAt, &test.UpdatedAt)
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
	query := (`SELECT 
		t.id, t.description, t.number, t.statement AS question, 
		t.alternative, t.response, t.response_id, 
		t.text_id, t.image_id, t.topic_id, t.note_id, 
		t.alternatives, t.date, t.created_at, t.updated_at
FROM 
    public.test_uece t
LEFT JOIN 
    image img ON img.id = t.image_id
LEFT JOIN 
    text tx ON tx.id = t.text_id
LEFT JOIN 
    note nt ON nt.id = t.note_id
WHERE 
`)

	var params []interface{}
	var paramCount int

	if filters.University != 0 {
		paramCount++
		query += fmt.Sprintf(" AND t.university_id = $%d", paramCount)
		params = append(params, filters.Topic)
	}

	if filters.Topic != 0 {
		paramCount++
		query += fmt.Sprintf(" AND t.topic_id = $%d", paramCount)
		params = append(params, filters.Topic)
	}

	if filters.InitialDate != "" && filters.FinalDate != "" {
		initialYear := filters.InitialDate[:4]
		finalYear := filters.FinalDate[:4]

		query += fmt.Sprintf(" AND t.date BETWEEN '%s-01-01' AND '%s-12-31'", initialYear, finalYear)
	}

	if filters.Offset >= 0 {
		paramCount++
		query += fmt.Sprintf(" OFFSET $%d", paramCount)
		params = append(params, filters.Offset)
	}

	if filters.Limit > 0 {
		paramCount++
		query += fmt.Sprintf(" LIMIT $%d", paramCount)
		params = append(params, filters.Limit)
	}

	rows, err := DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []Test

	for rows.Next() {
		var test Test
		err := rows.Scan(
			&test.Id, &test.Description, &test.Number, &test.Question,
			&test.Alternative, &test.Response, &test.ResponseId,
			&test.TextId, &test.ImageId, &test.TopicId, &test.NoteId,
			pq.Array(&test.Alternatives), &test.Date, &test.CreatedAt, &test.UpdatedAt,
		)
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
