package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)
	

type (
	MemberService interface {
		AddMemberToClass(ctx context.Context, member *dto.AddMemberRequest) (*entities.Member, error)
		GetAllMembersByClassID(ctx context.Context, classID uuid.UUID) ([]*entities.Member, error)
		// GetMemberById(ctx context.Context, id string) (*entities.Member, error)
		// UpdateMember(ctx context.Context, id string, member *dto.UpdateMemberRequest) (*entities.Member, error)
		DeleteMember(ctx context.Context, id uuid.UUID) error
		GetAllClassAndAssesmentByUserID(ctx context.Context, userID uuid.UUID) ([]dto.GetClassAndAssignmentResponse, error)
	}

	memberService struct {
		memberRepo database.StudentRepository
	}
)

func NewMemberService(memberRepo database.StudentRepository) MemberService {
	return &memberService{memberRepo}
}

func (service *memberService) AddMemberToClass(ctx context.Context, member *dto.AddMemberRequest) (*entities.Member, error) {
	// Check if the member already exists in the class
	existingMember, err := service.memberRepo.GetMemberByClassIDAndUserID(ctx, nil,member.Kelas_kelasID ,member.User_userID)
	if err != nil {
		return nil, err
	}	
	print("existingMember",existingMember)

	memberEntity := &entities.Member{
		Username:      member.Username,
		User_userID:   member.User_userID,
		Kelas_kelasID: member.Kelas_kelasID,
	}
	newMember, err := service.memberRepo.AddStudentToClass(ctx, nil, memberEntity)
	if err != nil {
		return nil, err
	}
	return newMember, nil
}

func (service *memberService) GetAllClassAndAssesmentByUserID(ctx context.Context, userID uuid.UUID) ([]dto.GetClassAndAssignmentResponse, error) {
	classes, err := service.memberRepo.GetAllClassAndAssesmentByUserID(ctx, nil, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (service *memberService) GetAllMembersByClassID(ctx context.Context, classID uuid.UUID) ([]*entities.Member, error) {
	members, err := service.memberRepo.GetAllMembersByClassID(ctx, nil, classID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

















































func (service *memberService) DeleteMember(ctx context.Context, id uuid.UUID) error {
	member, err := service.memberRepo.GetMemberById(ctx, nil, id)
	if err != nil {
		return err
	}
	if member == nil {
		return fmt.Errorf("member with ID %s not found", id)
	}
	err = service.memberRepo.DeleteMember(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}

