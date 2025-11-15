package repository

import (
	"context"
	"educationproject/src/internal/file/model"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CSVRepository struct {
	File *os.File
}

func NewOperatorRepository(file *os.File) model.OperatorRepository {
	return &CSVRepository{File: file}
}

func (csvReader *CSVRepository) Get(ctx context.Context) ([]model.Operator, error) {
	reader := csv.NewReader(csvReader.File)

	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	var listModel []model.Operator
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		age, err := strconv.Atoi(record[1])

		if err != nil {
			return nil, err
		}
		operatorModel := model.Operator{
			Age:   age,
			Name:  record[0],
			Email: record[2],
		}
		listModel = append(listModel, operatorModel)
	}
	return listModel, nil
}
