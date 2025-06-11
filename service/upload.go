package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func uploadFileItemPembelajaran(file io.Reader, fileName string) (string, error) {
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

func uploadFileAssignment(file io.Reader, fileName string) (string, error) {
	type FileUploadResponse struct {
		Id string `json:"id"`
	}
	
	fmt.Printf("Processing file upload\n")
	
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
	url := os.Getenv("CONTENT_URL") + "/item-pembelajaran/assignment"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	
	// Set proper content type with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	fmt.Printf("Sending file upload request\n")
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload file, content service is unavailable")
	}
	defer resp.Body.Close()
	
	fmt.Printf("Received response with status: %d\n", resp.StatusCode)
	
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

func uploadFileAssignmentSubmission(file io.Reader, fileName string, userID string) (string, error) {
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
	params := url.Values{}
	params.Add("userID", userID)
	url := os.Getenv("CONTENT_URL") + "/student-assignment/?" + params.Encode()
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