package entities

import (
	"time"

	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
	uuid "github.com/satori/go.uuid"
)

type PersonStatus string

const (
	PersontCreated  PersonStatus = "created"
	PersontAccepted PersonStatus = "accepted"
	PersontVerified PersonStatus = "rejected"
	PersontRejected PersonStatus = "failed"
)

type PersonEntity struct {
	ID                 uuid.UUID
	Status             PersonStatus
	Name               string
	Cpf                personValueObject.Cpf
	Birthdate          time.Time
	LegalName          *string
	DefaultProbability *int8
	DefaultScore       *int8
	DefaultDebits      *int8
	DefaultProtests    *int8
	IsActive           bool
	CreatedAt          time.Time
	AcceptedAt         *time.Time
	RejectedAt         *time.Time
	FailedAt           *time.Time
	UpdatedAt          time.Time
	UserID             uuid.UUID
}

func NewPersonEntity(name string, cpf string, birthdate time.Time, userID uuid.UUID) (*PersonEntity, error) {
	newCPF, cpfErr := personValueObject.NewCpf(cpf)
	if cpfErr != nil {
		return nil, cpfErr
	}
	person := PersonEntity{
		ID:        uuid.NewV4(),
		Status:    PersontCreated,
		Name:      name,
		Cpf:       *newCPF,
		Birthdate: birthdate,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
	}
	return &person, nil
}

func (p *PersonEntity) GetID() uuid.UUID {
	return p.ID
}

func (p *PersonEntity) GetStatus() PersonStatus {
	return p.Status
}

func (p *PersonEntity) GetName() string {
	return p.Name
}

func (p *PersonEntity) GetCpf() personValueObject.Cpf {
	return p.Cpf
}

func (p *PersonEntity) GetBirthdate() time.Time {
	return p.Birthdate
}

func (p *PersonEntity) GetLegalName() *string {
	return p.LegalName
}

func (p *PersonEntity) GetDefaultProbability() *int8 {
	return p.DefaultProbability
}

func (p *PersonEntity) GetDefaultScore() *int8 {
	return p.DefaultScore
}

func (p *PersonEntity) GetDefaultDebits() *int8 {
	return p.DefaultDebits
}

func (p *PersonEntity) GetDefaultProtests() *int8 {
	return p.DefaultProtests
}

func (p *PersonEntity) GetIsActive() bool {
	return p.IsActive
}

func (p *PersonEntity) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *PersonEntity) GetAcceptedAt() *time.Time {
	return &p.CreatedAt
}

func (p *PersonEntity) GetRejectedAt() *time.Time {
	return p.RejectedAt
}

func (p *PersonEntity) GetFailedAt() *time.Time {
	return p.FailedAt
}

func (p *PersonEntity) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *PersonEntity) GetUserID() uuid.UUID {
	return p.UserID
}
