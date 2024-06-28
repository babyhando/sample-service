package presenter

import (
	"bytes"
	"math"
	"time"
)

type Timestamp time.Time

func (d *Timestamp) MarshalJSON() ([]byte, error) {
	value := time.Time(*d).Format(time.DateTime)
	return []byte("\"" + value + "\""), nil
}

func (d *Timestamp) UnmarshalJSON(v []byte) error {
	v = bytes.ReplaceAll(v, []byte("\""), []byte(""))

	t, err := time.Parse(time.DateTime, string(v))
	if err != nil {
		return err
	}

	*d = Timestamp(t)
	return nil
}

type PaginationResponse[T any] struct {
	Page       uint `json:"page"`
	PageSize   uint `json:"pageSize"`
	TotalPages uint `json:"totalPages"`
	Data       []T  `json:"data"`
}

func NewPagination[T any](data []T, page, pageSize, total uint) *PaginationResponse[T] {
	totalPages := uint(0)
	if pageSize > 0 && total > 0 {
		totalPages = uint(math.Ceil(float64(total) / float64(pageSize)))
	}
	return &PaginationResponse[T]{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		Data:       data,
	}
}
