package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"online-library/pkg/redis"
)

type Service interface {
	SendOTP(ctx context.Context, phone string) error
	VerifyOTP(ctx context.Context, phone, otp string) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}

type service struct {
	repo       Repository
	rdb        *redis.Client
	apiKey     string
	lineNumber string
}

// new service for user service
func NewService(repo Repository, rdb *redis.Client) Service {
	return &service{
		repo:       repo,
		rdb:        rdb,
		apiKey:     "dv2MaMViWGZbQuYe5Io0XYVbpyXHlA4aceoKcVQGF5vWnw3f",
		lineNumber: "1",
	}
}

// generateOTP
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// generateOTP code and save in redis and send sms with sms.ir
func (s *service) SendOTP(ctx context.Context, phone string) error {
	otp := generateOTP()

	// save in redis (1 minute)
	if err := s.rdb.Set(ctx, phone, otp, time.Minute); err != nil {
		return err
	}

	// send sms with sms.ir
	url := "https://api.sms.ir/v1/send/verify"
	payload := map[string]interface{}{
		"mobile":     phone,
		"templateId": 530770,
		"parameters": []map[string]string{
			{
				"name":  "Code",
				"value": otp,
			},
		},
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send sms, status: %d", resp.StatusCode)
	}

	return nil
}

// VerifyOTP check code and create/return user if needed
func (s *service) VerifyOTP(ctx context.Context, phone, otp string) (*User, error) {
	stored, err := s.rdb.Get(ctx, phone)
	if err != nil {
		return nil, errors.New("otp expired or not found")
	}

	if stored != otp {
		return nil, errors.New("invalid otp")
	}

	// if user not found, create user
	u, err := s.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if u == nil {
		if err := s.repo.CreateUser(ctx, phone); err != nil {
			return nil, err
		}
		u, err = s.repo.GetUserByPhone(ctx, phone)
		if err != nil {
			return nil, err
		}
	}

	// (optional) we can delete otp from redis if we have Delete method; add delete method if needed
	return u, nil
}

// GetUserByPhone wrapper on repository
func (s *service) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	return s.repo.GetUserByPhone(ctx, phone)
}
