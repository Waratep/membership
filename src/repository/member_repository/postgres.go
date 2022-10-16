package member_repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/Waratep/membership/src/entity/member"
	"github.com/Waratep/membership/src/use_case"
	"github.com/gin-gonic/gin"
)

type postgresDB struct {
	db *sql.Conn
}

type memberPostgres struct {
	ID        int       `bson:"id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Phone     string    `bson:"phone"`
	Email     string    `bson:"email"`
	Address   string    `bson:"address"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func NewPostgres(db *sql.Conn) use_case.MemberRepository {
	p := &postgresDB{db: db}

	return p
}

func (p postgresDB) HealthCheck(ctx gin.Context) error {
	err := p.db.PingContext(&ctx)
	if err != nil {
		log.Println("Error ping postgres database")

		return err
	}

	return nil
}

func (m memberPostgres) toMemberUseCase() use_case.MembershipMember {
	return use_case.MembershipMember{
		ID: int64(m.ID),
		Member: member.Member{
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Phone:     m.Phone,
			Email:     m.Email,
			Address:   m.Address,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}

func (p postgresDB) CreateMember(ctx gin.Context, m member.Member) (use_case.MembershipMember, error) {
	_, err := p.db.ExecContext(
		&ctx,
		`
		INSERT INTO members 
			(
				first_name, 
				last_name, 
				phone, 
				email, 
				address,
				created_at,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		m.FirstName, m.LastName, m.Phone, m.Email, m.Address, time.Now().UTC(), time.Now().UTC())
	if err != nil {
		log.Println("Error insert member", err)

		return use_case.MembershipMember{}, err
	}

	return use_case.MembershipMember{
		Member: m,
	}, nil
}

func (p postgresDB) GetMembers(ctx gin.Context) error {
	return nil
}

func (p postgresDB) GetMemberByID(ctx gin.Context) error {
	return nil
}

func (p postgresDB) GetMemberByPhone(ctx gin.Context, phone string) (use_case.MembershipMember, error) {
	rows, err := p.db.QueryContext(&ctx, "SELECT * FROM members WHERE phone=$1", phone)
	if err != nil {
		log.Println("Error query member by phone", err)

		return use_case.MembershipMember{}, err
	}
	defer rows.Close()

	var ms []memberPostgres
	for rows.Next() {
		var m memberPostgres
		err := rows.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Phone, &m.Email, &m.Phone, &m.CreatedAt, &m.UpdatedAt)
		if err != nil {
			log.Println("Error parse member struct", err)

			return use_case.MembershipMember{}, err
		}

		ms = append(ms, m)
	}

	if len(ms) == 0 {
		return use_case.MembershipMember{}, use_case.ErrorItemNotFound
	}

	return ms[0].toMemberUseCase(), nil
}

func (p postgresDB) GetMemberByEmail(ctx gin.Context, email string) (use_case.MembershipMember, error) {
	rows, err := p.db.QueryContext(&ctx, "SELECT * FROM members WHERE email=$1", email)
	if err != nil {
		log.Println("Error query member by phone", err)

		return use_case.MembershipMember{}, err
	}
	defer rows.Close()

	var ms []memberPostgres
	for rows.Next() {
		var m memberPostgres
		err := rows.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Phone, &m.Email, &m.Phone, &m.CreatedAt, &m.UpdatedAt)
		if err != nil {
			log.Println("Error parse member struct", err)

			return use_case.MembershipMember{}, err
		}

		ms = append(ms, m)
	}
	if len(ms) == 0 {
		return use_case.MembershipMember{}, use_case.ErrorItemNotFound
	}

	return ms[0].toMemberUseCase(), nil
}

func (p postgresDB) UpdateMember(ctx gin.Context) error {
	return nil
}

func (p postgresDB) DeleteMember(ctx gin.Context) error {
	return nil
}
