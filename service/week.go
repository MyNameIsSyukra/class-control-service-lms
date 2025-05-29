package service

import (
	dto "LMSGo/dto"
	entities "LMSGo/entity"
	database "LMSGo/repository"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type (
	WeekService interface {
		GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) ([]dto.WeekResponse, error)
		GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponseByID, error)

		// teacher
		CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest,file io.Reader) (*entities.ItemPembelajaran, error)
		DeleteWeeklySection(ctx context.Context, weekID int) error
		UpdateWeeklySection(ctx context.Context, req dto.UpdateItemPembelajaranRequest,file io.Reader) (*entities.ItemPembelajaran, error)
		// CreateWeeklySection(ctx context.Context, weekReq dto.WeekRequest) (*entities.Week, error)
		// DeleteWeeklySection(ctx context.Context, weekID int) error
	}
	weekService struct {
		weekRepo database.WeekRepository
		kelasRepo database.KelasRepository
	}
)

func NewWeekService(weekRepo database.WeekRepository, kelasRepo database.KelasRepository) WeekService {
	return &weekService{weekRepo , kelasRepo}
}



func (service *weekService) GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) ([]dto.WeekResponse, error) {
	class, err := service.kelasRepo.GetById(ctx, nil, classID)
	if err != nil {
		return []dto.WeekResponse{}, err
	}
	if class.ID == uuid.Nil {
		return []dto.WeekResponse{}, fmt.Errorf("class with ID %s not found", classID)
	}
	weeks, err := service.weekRepo.GetAllWeekByClassID(ctx, nil, classID)
	if err != nil {
		return []dto.WeekResponse{}, err
	}
	if len(weeks) == 0 {
		return []dto.WeekResponse{}, nil
	}
		
	var weekResponse []dto.WeekResponse
	for _, week := range weeks {
		var weekRes dto.WeekResponse
		if week.Assignment.Title == "" {
			weekRes.Assignment = nil
			weekRes.ItemPembelajarans = &week.ItemPembelajaran
		}else if week.ItemPembelajaran.HeadingPertemuan == "" {
			weekRes.ItemPembelajarans = nil
			weekRes.Assignment = &week.Assignment
		}else {
			weekRes.ItemPembelajarans = &week.ItemPembelajaran
			weekRes.Assignment = &week.Assignment
		}
		weekRes.WeekID = week.ID
		weekRes.WeekNumber = week.WeekNumber
		weekRes.KelasID = week.Kelas_idKelas
		weekResponse = append(weekResponse, weekRes)
	}		


	return weekResponse, nil
}

func (service *weekService) GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponseByID, error) {
	week, err := service.weekRepo.GetWeekByID(ctx, nil, weekID)
	if err != nil {
		return dto.WeekResponseByID{}, err
	}
	resp := dto.WeekResponseByID{
		WeekID:           week.ID,
		WeekNumber:       week.WeekNumber,
		KelasID:          week.Kelas_idKelas,
		ItemPembelajarans: &week.ItemPembelajaran,
		Assignment:       &week.Assignment,
	}
	if week.Assignment.Title == "" {
		resp.Assignment = nil
	}
	if week.ItemPembelajaran.HeadingPertemuan == "" {
		resp.ItemPembelajarans = nil
	}
	return resp, nil
}

func (service *weekService) CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest, file io.Reader) (*entities.ItemPembelajaran, error) {
	class , err := service.kelasRepo.GetById(ctx, nil, request.KelasID)
	if err != nil {
		return nil, fmt.Errorf("class with ID %s not found", request.KelasID)
	}
	if class.ID == uuid.Nil {
		return nil, fmt.Errorf("class with ID %s not found", request.KelasID)
	}
	newWeekRequest := dto.WeekRequest{
		WeekNumber:    request.WeekNumber,
		Kelas_idKelas: request.KelasID,
	}
	newWeek, err := service.weekRepo.CreateWeeklySection(ctx, nil, newWeekRequest)
	if err != nil {
		return nil, err
	}
	if file != nil {
		fileURL, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return nil, err
		}
		request.FileLink = fileURL
	} else {
		fmt.Printf("No file provided\n")
		request.FileLink = ""
		request.FileName = ""
	}
	newItem := &entities.ItemPembelajaran{
		WeekID: newWeek.ID,
		HeadingPertemuan: request.HeadingPertemuan,
		BodyPertemuan: request.BodyPertemuan,
		UrlVideo: request.UrlVideo,
		FileName: request.FileName,
		FileLink: request.FileLink,
	}
	newItem, err = service.weekRepo.CreateItemPembelajaran(ctx, nil, newItem)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

// Update WeeklySection is not implemented in the original code, so we will not implement it here.
func (service *weekService) UpdateWeeklySection(ctx context.Context,req dto.UpdateItemPembelajaranRequest, file io.Reader) (*entities.ItemPembelajaran, error) {
	oldItem, err := service.weekRepo.GetItemPembelajaran(ctx, nil, req.WeekID)
	if err != nil {
		return nil, fmt.Errorf("failed to get old item pembelajaran: %w", err)
	}
	if file != nil {
		deleteURL := oldItem.FileLink
		if (deleteURL != "") {
			delReq, err := http.NewRequest(http.MethodDelete, deleteURL, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create delete request: %w", err)
			}
			client := &http.Client{}
			resp, err := client.Do(delReq)
			if err != nil {
				return nil, fmt.Errorf("failed to delete old file: %w", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusNoContent {
				respBody, _ := io.ReadAll(resp.Body)
				return nil, fmt.Errorf("failed to delete old file with status %d: %s", resp.StatusCode, string(respBody))
			}
		}
		fileURL, err := service.uploadFile(file, req.FileName)
		if err != nil {
			return nil, err
		}
		req.FileLink = fileURL
	} else {
		fmt.Printf("No file provided\n")
		req.FileLink = ""
		req.FileName = ""
	}
	item := &entities.ItemPembelajaran{
		WeekID: req.WeekID,
		HeadingPertemuan: req.HeadingPertemuan,
		BodyPertemuan: req.BodyPertemuan,
		UrlVideo: req.UrlVideo,
		FileName: req.FileName,
		FileLink: req.FileLink,
	}
	
	item, err = service.weekRepo.UpdateItemPembelajaran(ctx, nil, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
	

func (service *weekService) DeleteWeeklySection(ctx context.Context, weekID int) error {
	err := service.weekRepo.DeleteWeeklySection(ctx, nil, weekID)
	if err != nil {
		return err
	}
	return nil
}


func (service *weekService) uploadFile(file io.Reader, fileName string) (string, error) {
	type FileUploadResponse struct {
		Url string `json:"url"`
	}
	
	// fmt.Printf("Processing file upload\n")
	
	// Create multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	
	// Create form file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	
	// Copy file content to form
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}
	
	// Close writer to finalize multipart data
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %w", err)
	}
	
	// Prepare HTTP request
	url := os.Getenv("CONTENT_URL") + "/item-pembelajaran/"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	
	// Set proper content type with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	// fmt.Printf("Sending file upload request\n")
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	
	// fmt.Printf("Received response with status: %d\n", resp.StatusCode)
	
	// Check response status
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("file upload failed with status %d: %s", resp.StatusCode, string(respBody))
	}
	
	// Parse response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	
	var uploadResp FileUploadResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return "", fmt.Errorf("failed to parse upload response: %w", err)
	}
	
	fmt.Printf("File uploaded successfully: %s\n", uploadResp.Url)
	return uploadResp.Url, nil
}