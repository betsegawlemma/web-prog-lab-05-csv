package service

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/betsegawlemma/webprogcsv/entity"
)

// CategoryService represents csv implementation of menu.CategoryService
type CategoryService struct {
	FileName string
}

// NewCategoryService returns new Category Service
func NewCategoryService(fileName string) *CategoryService {
	return &CategoryService{FileName: fileName}
}

// Categories returns all categories read from csv file
func (cs CategoryService) Categories() ([]entity.Category, error) {
	file, err := os.Open(cs.FileName)
	if err != nil {
		return nil, errors.New("File could not be open")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("File could not be open")
	}
	var ctgs []entity.Category
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		c := entity.Category{ID: int(id), Name: item[1],
			Description: item[2], Image: item[3]}
		ctgs = append(ctgs, c)
	}
	return ctgs, nil
}

// StoreCategories stores a batch of categories data to the a csv file
func (cs CategoryService) StoreCategories(ctgs []entity.Category) error {
	csvFile, err := os.Create(cs.FileName)
	if err != nil {
		return errors.New("File could not be created")
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, c := range ctgs {
		line := []string{strconv.Itoa(c.ID), c.Name, c.Description, c.Image}
		writer.Write(line)
	}
	writer.Flush()
	return nil
}
