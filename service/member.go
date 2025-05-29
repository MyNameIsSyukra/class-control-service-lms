package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)
	

type (
	MemberService interface {
		AddMemberToClass(ctx context.Context, member *dto.AddMemberRequest) (*entities.Member, error)
		GetAllMembersByClassID(ctx context.Context, classID uuid.UUID) ([]dto.GetMemberResponse, error)
		// GetMemberById(ctx context.Context, id string) (*entities.Member, error)
		// UpdateMember(ctx context.Context, id string, member *dto.UpdateMemberRequest) (*entities.Member, error)
		DeleteMember(ctx context.Context, id uuid.UUID) error
		GetAllClassAndAssesmentByUserID(ctx context.Context, userID uuid.UUID) ([]dto.GetClassAndAssignmentResponse, error)
		GetAllClassByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Kelas, error)
		// Lintas Service
		GetMemberByClassIDAndUserID(ctx context.Context, classID, userID uuid.UUID) (*entities.Member, error)
	}

	memberService struct {
		memberRepo database.StudentRepository
		kelasRepo database.KelasRepository
	}
)

func NewMemberService(memberRepo database.StudentRepository, kelasRepo database.KelasRepository) MemberService {
	return &memberService{memberRepo, kelasRepo}
}

func (service *memberService) AddMemberToClass(ctx context.Context, member *dto.AddMemberRequest) (*entities.Member, error) {
	// check if the class exists
	kelas, err := service.kelasRepo.GetById(ctx, nil, member.Kelas_kelasID)
	if err != nil {
		return nil, err
	}
	if kelas.ID == uuid.Nil {
		return nil, fmt.Errorf("class with ID %s not found", member.Kelas_kelasID)
	}
	// Check if the member already exists in the class
	_, err = service.memberRepo.GetMemberByClassIDAndUserID(ctx, nil,member.Kelas_kelasID ,member.User_userID)
	if err != nil {
		return nil, err
	}	
	
	memberEntity := &entities.Member{
		Username:      member.Username,
		Role:          member.Role,
		User_userID:   member.User_userID,
		Kelas_kelasID: member.Kelas_kelasID,
	}
	newMember, err := service.memberRepo.AddMemberToClass(ctx, nil, memberEntity)
	if err != nil {
		return nil, err
	}
	if newMember.Role == entities.MemberRoleTeacher {
		_, err := service.kelasRepo.UpdateClassTeacherID(ctx, nil, member.Kelas_kelasID, member.User_userID, member.Username)
		if err != nil {
			return nil, err
		}
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

func (service *memberService) GetAllMembersByClassID(ctx context.Context, classID uuid.UUID) ([]dto.GetMemberResponse, error) {
	members, err := service.memberRepo.GetAllMembersByClassID(ctx, nil, classID)
	if err != nil {
		return nil, err
	}
	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	url := os.Getenv("GATEWAY_API_URL")
	var response []dto.GetMemberResponse
	for _, member := range members {
		response = append(response, dto.GetMemberResponse{
			Username:      member.Username,
			User_userID:   member.User_userID,
			Role:          member.Role,
			Kelas_kelasID: member.Kelas_kelasID,
			PhotoUrl: fmt.Sprintf("%s/storage/user_profile_pictures/%s.jpg", url, member.User_userID),
		})
	}
	return response, nil
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

func (service *memberService) GetAllClassByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Kelas, error) {
	classes, err := service.memberRepo.GetAllClassByUserID(ctx, nil, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (service *memberService) GetMemberByClassIDAndUserID(ctx context.Context, classID, userID uuid.UUID) (*entities.Member, error) {
	member, err := service.memberRepo.GetMemberByClassIDAndUserID(ctx, nil, classID, userID)
	if err != nil {
		return nil, err
	}
	if member == nil {
		return nil, fmt.Errorf("member with class ID %s and user ID %s not found", classID, userID)
	}
	if member.User_userID == uuid.Nil {
		return nil, fmt.Errorf("member with class ID %s and user ID %s not found", classID, userID)
	}
	return member, nil
}