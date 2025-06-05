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
		AddMemberToClass(ctx context.Context, member *dto.InitAddMemberRequest) ([]entities.Member, []error)
		GetAllMembersByClassID(ctx context.Context, classID uuid.UUID) ([]dto.GetMemberResponse, error)
		// GetMemberById(ctx context.Context, id string) (*entities.Member, error)
		// UpdateMember(ctx context.Context, id string, member *dto.UpdateMemberRequest) (*entities.Member, error)
		DeleteMember(ctx context.Context, user_id uuid.UUID, class_id uuid.UUID) error
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

func (service *memberService) AddMemberToClass(ctx context.Context, member *dto.InitAddMemberRequest) ([]entities.Member, []error) {
	// check if the class exists
	var errs []error
	kelas, err := service.kelasRepo.GetById(ctx, nil, member.Kelas_kelasID)
	if err != nil {
		errs = append(errs, fmt.Errorf("error retrieving class with ID %s: %v", member.Kelas_kelasID, err))
		return nil, errs
	}
	if kelas.ID == uuid.Nil {
		errs = append(errs, fmt.Errorf("class with ID %s not found", member.Kelas_kelasID))
		return nil, errs
	}
	
	// check if requested members
	for _, student := range member.Students {
		// Check if the member already exists in the class
		checkStudent, err := service.memberRepo.GetMemberByClassIDAndUserID(ctx, nil,member.Kelas_kelasID ,student.User_userID)
		if err != nil {
			errs = append(errs, fmt.Errorf("error checking member existence for user %s: %v", student.Username, err))
			continue
		}
		if checkStudent != nil {
			errs = append(errs, fmt.Errorf("member %s with user ID %s already exists in class %s", student.Username, student.User_userID, checkStudent.Kelas.Name))
			continue
		}
	}
	if len(errs) > 0 {
		return nil, errs
	}
	var newMembers []entities.Member
	for _, student := range member.Students {
		memberEntity := entities.Member{
			Username:      student.Username,
			User_userID:   student.User_userID,
			Kelas_kelasID: member.Kelas_kelasID,
		}
		newMember, err := service.memberRepo.AddMemberToClass(ctx, nil, &memberEntity)
		if err != nil {	
			errs = append(errs, fmt.Errorf("error adding member %s to class %s: %v", student.Username, member.Kelas_kelasID, err))
			continue
		}
		newMembers = append(newMembers, *newMember)
	}
	return newMembers, errs
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
	url := os.Getenv("GATEWAY_URL")
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

func (service *memberService) DeleteMember(ctx context.Context, user_id uuid.UUID, class_id uuid.UUID) error {
	member, err := service.memberRepo.GetMemberByClassIDAndUserID(ctx, nil, class_id, user_id)
	if err != nil {
		return err
	}
	if member == nil {
		return fmt.Errorf("member with ID %s not found in class %s", user_id,class_id)
	}
	err = service.memberRepo.DeleteMember(ctx, nil, user_id, class_id)
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