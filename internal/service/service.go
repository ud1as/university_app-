package service

import (
	"context"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/storage"
	"go.uber.org/zap"
)

type IService interface {
	AllStudentsService(ctx context.Context) (student []model.Student, err error)
	StudentByID(ctx context.Context, id int) (student model.Student, err error)
	DeleteStudentById(ctx context.Context, id int) (err error)
	UpdateStudent(ctx context.Context, student model.Student, id int) (err error)
	AddNewStudent(ctx context.Context, student model.Student) (id int, err error)
}

type Service struct {
	conf   *config.Config
	logger *zap.Logger
	urepo  *storage.Repo
}

func NewService(conf *config.Config, logger *zap.Logger, urepo *storage.Repo) *Service {
	return &Service{conf: conf, logger: logger, urepo: urepo}
}

func (s *Service) AllStudentsService(ctx context.Context) (student []model.Student, err error) {
	student, err = s.urepo.AllStudents(ctx)

	return student, err
}

func (s *Service) StudentByID(ctx context.Context, id int) (student model.Student, err error) {
	student, err = s.urepo.StudentByID(ctx, id)

	return student, err
}

func (s *Service) DeleteStudentById(ctx context.Context, id int) (err error) {
	_, err = s.urepo.DeleteStudentById(ctx, id)

	return err
}

func (s *Service) UpdateStudent(ctx context.Context, student model.Student, id int) (err error) {
	err = s.urepo.UpdateStudent(ctx, student, id)

	return err
}

func (s *Service) AddNewStudent(ctx context.Context, student model.Student) (id int, err error) {
	id, err = s.urepo.AddNewStudent(ctx, student)

	return id, err
}
