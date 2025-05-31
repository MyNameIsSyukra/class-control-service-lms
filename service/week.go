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
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	WeekService interface {
		GetAllWeekByClassID(ctx context.Context, classID uuid.UUID) ([]dto.WeekResponse, error)
		GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponse, error)

		// teacher
		CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest,file io.Reader) (*dto.ItemPembelajaranResponse, error)
		DeleteWeeklySection(ctx context.Context, weekID int) error
		UpdateWeeklySection(ctx context.Context, req dto.UpdateItemPembelajaranRequest,file io.Reader) (*dto.ItemPembelajaranResponse, error)
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
		
		// Check if Assignment is empty
		hasAssignment := week.Assignment.Title != "" && 
						week.Assignment.Description != "" && 
						week.Assignment.ID != 0
		
		// Check if ItemPembelajaran is empty
		hasItemPembelajaran := week.ItemPembelajaran.HeadingPertemuan != "" && 
							  week.ItemPembelajaran.BodyPertemuan != "" && 
							  week.ItemPembelajaran.WeekID != 0
		
		// Set Assignment data if exists
		if hasAssignment {
			weekRes.Assignment = &dto.AssignmentResponse{
				AssignmentID: int(week.Assignment.ID),
				Title:        week.Assignment.Title,
				Description:  week.Assignment.Description,
				Deadline:     week.Assignment.Deadline,
				FileName:     week.Assignment.FileName,
				FileId:       week.Assignment.FileId,
				FileUrl:      os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + "?" + url.Values{"id": []string{week.Assignment.FileId}}.Encode(),
			}
		} else {
			weekRes.Assignment = nil
		}

		// Set ItemPembelajaran data if exists
		if hasItemPembelajaran {
			weekRes.ItemPembelajarans = &dto.ItemPembelajaranResponse{
				WeekID:           week.ItemPembelajaran.WeekID,
				HeadingPertemuan: week.ItemPembelajaran.HeadingPertemuan,
				BodyPertemuan:    week.ItemPembelajaran.BodyPertemuan,
				UrlVideo:         week.ItemPembelajaran.UrlVideo,
				FileName:         week.ItemPembelajaran.FileName,
				FileId:           week.ItemPembelajaran.FileId,
				FileUrl:          os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + "?" + url.Values{"id": []string{week.ItemPembelajaran.FileId}}.Encode(),
			}
		} else {
			weekRes.ItemPembelajarans = nil
		}

		// Set common week data
		weekRes.WeekID = week.ID
		weekRes.WeekNumber = week.WeekNumber
		weekRes.KelasID = week.Kelas_idKelas

		// Optional: Log if both are empty
		if !hasAssignment && !hasItemPembelajaran {
			fmt.Printf("Week %d has no assignment and no item pembelajaran\n", week.ID)
		}
		weekResponse = append(weekResponse, weekRes)
	}		


	return weekResponse, nil
}

func (service *weekService) GetWeekByID(ctx context.Context, weekID int) (dto.WeekResponse, error) {
	week, err := service.weekRepo.GetWeekByID(ctx, nil, weekID)
	if err != nil {
		return dto.WeekResponse{}, err
	}
	resp := dto.WeekResponse{
		WeekID:           week.ID,
		WeekNumber:       week.WeekNumber,
		KelasID:          week.Kelas_idKelas,
		ItemPembelajarans: &dto.ItemPembelajaranResponse{
			WeekID:           week.ItemPembelajaran.WeekID,
			HeadingPertemuan: week.ItemPembelajaran.HeadingPertemuan,
			BodyPertemuan:    week.ItemPembelajaran.BodyPertemuan,
			UrlVideo:         week.ItemPembelajaran.UrlVideo,
			FileName:         week.ItemPembelajaran.FileName,
			FileId:           week.ItemPembelajaran.FileId,
			FileUrl:          os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + week.ItemPembelajaran.FileId + "?" + url.Values{"id": []string{week.ItemPembelajaran.FileId}}.Encode(),
		},
		Assignment:       &dto.AssignmentResponse{
			AssignmentID:      int(week.Assignment.ID),
			Title:       week.Assignment.Title,
			Description: week.Assignment.Description,
			Deadline:    week.Assignment.Deadline,
			FileName:    week.Assignment.FileName,
			FileId:      week.Assignment.FileId,
			FileUrl:     os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + "?" + url.Values{"id": []string{week.Assignment.FileId}}.Encode(),
		},
	}
	if week.Assignment.Title == "" {
		resp.Assignment = nil
	}
	if week.ItemPembelajaran.HeadingPertemuan == "" {
		resp.ItemPembelajarans = nil
	}
	return resp, nil
}

func (service *weekService) CreateWeeklySection(ctx context.Context, request dto.CreateItemPembelajaranRequest, file io.Reader) (*dto.ItemPembelajaranResponse, error) {
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
		fileId, err := service.uploadFile(file, request.FileName)
		if err != nil {
			return nil, err
		}
		request.FileId = fileId
	} else {
		fmt.Printf("No file provided\n")
		request.FileId = ""
		request.FileName = ""
	}
	newItem := &entities.ItemPembelajaran{
		WeekID: newWeek.ID,
		HeadingPertemuan: request.HeadingPertemuan,
		BodyPertemuan: request.BodyPertemuan,
		UrlVideo: request.UrlVideo,
		FileName: request.FileName,
		FileId: request.FileId,
	}
	newItem, err = service.weekRepo.CreateItemPembelajaran(ctx, nil, newItem)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("id", newItem.FileId)
	fileUrl := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + "?" + params.Encode()
	return &dto.ItemPembelajaranResponse{
		WeekID:           newItem.WeekID,
		HeadingPertemuan: newItem.HeadingPertemuan,
		BodyPertemuan:    newItem.BodyPertemuan,
		UrlVideo:         newItem.UrlVideo,
		FileName:         newItem.FileName,
		FileId:           newItem.FileId,
		FileUrl:          fileUrl,
	}, nil
}

// Update WeeklySection is not implemented in the original code, so we will not implement it here.
func (service *weekService) UpdateWeeklySection(ctx context.Context,req dto.UpdateItemPembelajaranRequest, file io.Reader) (*dto.ItemPembelajaranResponse, error) {
	oldItem, err := service.weekRepo.GetItemPembelajaran(ctx, nil, req.WeekID)
	if err != nil {
		return nil, fmt.Errorf("failed to get old item pembelajaran: %w", err)
	}
	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	if file != nil {
		id := oldItem.FileId
		if (id != "") {
			params := url.Values{}
			params.Add("id", id)
			deleteURL := os.Getenv("CONTENT_URL") + "/item-pembelajaran/" + "?" + params.Encode()
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
		fileId, err := service.uploadFile(file, req.FileName)
		if err != nil {
			return nil, err
		}
		req.FileId = fileId
	} else {
		fmt.Printf("No file provided\n")
		req.FileId = ""
		req.FileName = ""
	}
	if req.HeadingPertemuan == "" {
		req.HeadingPertemuan = oldItem.HeadingPertemuan
	}
	if req.BodyPertemuan == "" {
		req.BodyPertemuan = oldItem.BodyPertemuan
	}
	if req.UrlVideo == "" {
		req.UrlVideo = oldItem.UrlVideo
	}
	// fmt.Printf("Updating item pembelajaran with WeekID: %d\n", req.WeekID)
	if req.WeekID == 0 {
		return nil, fmt.Errorf("WeekID cannot be zero")
	}
	if req.WeekID != oldItem.WeekID {
		return nil, fmt.Errorf("WeekID mismatch: expected %d, got %d", oldItem.WeekID, req.WeekID)
	}
	if req.FileName == "" {
		req.FileName = oldItem.FileName
	}
	// fmt.Printf("Creating new item pembelajaran with WeekID: %d\n", req.WeekID)
	if req.FileId == "" {
		req.FileId = oldItem.FileId
	}
	item := &entities.ItemPembelajaran{
		WeekID: req.WeekID,
		HeadingPertemuan: req.HeadingPertemuan,
		BodyPertemuan: req.BodyPertemuan,
		UrlVideo: req.UrlVideo,
		FileName: req.FileName,
		FileId: req.FileId,
	}
	
	item, err = service.weekRepo.UpdateItemPembelajaran(ctx, nil, item)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("id", item.FileId)
	fileUrl := os.Getenv("GATEWAY_URL") + "/item-pembelajaran/" + "?" + params.Encode() 
	return &dto.ItemPembelajaranResponse{
		WeekID:           item.WeekID,
		HeadingPertemuan: item.HeadingPertemuan,
		BodyPertemuan:    item.BodyPertemuan,
		UrlVideo:         item.UrlVideo,
		FileName:         item.FileName,
		FileId:           item.FileId,
		FileUrl:          fileUrl,
	}, nil
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
		Id string `json:"id"`
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
	
	fmt.Printf("File uploaded successfully: %s\n", uploadResp.Id)
	return uploadResp.Id, nil
}