package storage

import (
	"context"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewRepository(conf *config.Config, logger *zap.Logger) (*Repo, error) {
	pgDB, err := ConnectDB(conf)
	if err != nil {
		logger.Sugar().Error("Unable to connect")
		return nil, err
	}
	return &Repo{DB: pgDB}, nil
}

type Repo struct {
	DB *pgx.Conn
}

type IRepository interface {
	AllStudents(ctx context.Context) (student []model.Student, err error)
	StudentByID(ctx context.Context, id int) (student model.Student, err error)
	DeleteStudentById(ctx context.Context, id int) (err error)
	UpdateStudent(ctx context.Context, student model.Student, id int) (err error)
	AddNewStudent(ctx context.Context, student model.Student) (id int, err error)
}

func (r *Repo) StudentByID(ctx context.Context, id int) (student model.Student, err error) {

	query := `select id, name, surname, gender from students where id = $1 `
	err = r.DB.QueryRow(ctx, query, id).Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)

	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return student, err
	}

	return student, err
}

func (r *Repo) AllStudents(ctx context.Context) (students []model.Student, err error) {

	students = make([]model.Student, 0)
	query := `select id, name, surname, gender from students`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return nil, err
	}
	var student model.Student

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
		if err != nil {
			//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
			return nil, err
		}

		students = append(students, student)
	}

	defer rows.Close()
	return students, nil
}

func (r *Repo) AddNewStudent(ctx context.Context, student model.Student) (id int, err error) {
	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3) RETURNING id`

	err = r.DB.QueryRow(ctx, query, student.Name, student.Surname, student.Gender).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *Repo) UpdateStudent(ctx context.Context, student model.Student, id int) (err error) {
	query := `UPDATE public.students
	SET name=$2, surname = $3, gender = $4 
	WHERE id = $1;`
	err = r.DB.QueryRow(ctx, query, id, student.Name, student.Surname, student.Gender).Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return err
	}

	return err
}

func (r *Repo) DeleteStudentById(ctx context.Context, id int) (int int, err error) {
	query := `DELETE FROM students_by_group WHERE student_id = $1`
	query2 := `DELETE FROM students WHERE id = $1`

	_, err = r.DB.Exec(ctx, query, id)
	_, err = r.DB.Exec(ctx, query2, id)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return -1, err
	}

	return id, err
}
