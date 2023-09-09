package repositories

import (
	"context"
	"errors"

	"github.com/IsaacDSC/rinha-backend-dsc/config"
	"github.com/IsaacDSC/rinha-backend-dsc/internal/models"
)

type PersonRepository struct{}

type ResultPerson struct {
	ID, LastName, Name, Birthday, Stack string
}

func (*PersonRepository) Create(ctx context.Context, person models.Person) (err error) {
	conn := config.DbConn()
	transaction, err := conn.Begin()
	if err != nil {
		return err
	}
	const query = `INSERT INTO public.person(
		id, username, name, birthday, updated_at, created_at)
		VALUES ($1,$2,$3,$4,NOW(),NOW());`
	row, err := transaction.ExecContext(
		ctx,
		query,
		person.ID,
		person.LastName,
		person.Name,
		person.Birthday,
	)
	if err != nil {
		return
	}
	affected, _ := row.RowsAffected()
	if affected != 1 {
		err = errors.New("Not-Affected-Row")
		return
	}
	const query_stack = `insert into stack (stack_name, person_id) VALUES ($1, $2);`
	for index := range person.Stack {
		_, err := transaction.ExecContext(ctx, query_stack, person.Stack[index], person.ID)
		if err != nil {
			transaction.Rollback()
			break
		}
	}
	if err != nil {
		return
	}
	transaction.Commit()
	return
}

func (*PersonRepository) FindById(ctx context.Context, personID string) (
	output models.Person, err error,
) {
	conn := config.DbConn()
	const query = `SELECT person.id, person.username, person.name, person.birthday, stack.stack_name FROM stack
	join person on stack.person_id = person.id 
	WHERE id = $1`
	row, err := conn.QueryContext(ctx, query, personID)
	if err != nil {
		return
	}
	var results []ResultPerson
	for row.Next() {
		var r ResultPerson
		if err = row.Scan(
			&r.ID,
			&r.LastName,
			&r.Name,
			&r.Birthday,
			&r.Stack,
		); err != nil {
			return
		}
		results = append(results, r)
	}
	for index := range results {
		if index == 0 {
			output = models.Person{
				ID:       results[index].ID,
				Name:     results[index].Name,
				LastName: results[index].LastName,
				Birthday: results[index].Birthday,
			}
			output.Stack = append(output.Stack, results[index].Stack)
		} else {
			output.Stack = append(output.Stack, results[index].Stack)
		}
	}
	return
}

func (*PersonRepository) Search(ctx context.Context, person models.Person) (
	output []models.Person, err error,
) {
	query := `SELECT person.id, person.username, person.name, person.birthday, stack.stack_name FROM stack
	join person on stack.person_id = person.id 
	WHERE username = $1 or name = $2 or birthday = $3 or stack_name = $4
	order by person.id;
	`
	var stackSearch string
	if len(person.Stack) > 0 {
		stackSearch = person.Stack[0]
	}
	conn := config.DbConn()
	rows, err := conn.QueryContext(
		ctx,
		query,
		person.LastName,
		person.Name,
		person.Birthday,
		stackSearch,
	)
	if err != nil {
		return
	}
	var aux int = 0
	for rows.Next() {
		var p ResultPerson
		if err := rows.Scan(
			&p.ID,
			&p.LastName,
			&p.Name,
			&p.Birthday,
			&p.Stack,
		); err != nil {
			break
		}
		if aux == 0 && len(p.ID) > 0 {
			output = append(output, models.Person{
				ID:       p.ID,
				Name:     p.Name,
				LastName: p.LastName,
				Birthday: p.Birthday,
				Stack:    []string{p.Stack},
			})
			aux++
		} else if aux > 0 {
			if output[aux-1].ID == p.ID {
				output[aux-1].Stack = append(output[aux-1].Stack, p.Stack)
			} else {
				output = append(output, models.Person{
					ID:       p.ID,
					Name:     p.Name,
					LastName: p.LastName,
					Birthday: p.Birthday,
					Stack:    []string{p.Stack},
				})
				aux++
			}
		}
	}
	return
}

func (*PersonRepository) CounterPersons(ctx context.Context) (total_person int32, err error) {
	conn := config.DbConn()
	const query = `SELECT COUNT(*) AS total_person FROM person;`
	row := conn.QueryRowContext(ctx, query)
	if row.Err() != nil {
		return
	}
	row.Scan(
		&total_person,
	)
	return
}
