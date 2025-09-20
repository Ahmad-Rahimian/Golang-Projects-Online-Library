package auth

import (
	"net/http"
	"time"

	"online-library/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	svc       user.Service
	jwtSecret string
}

func NewHandler(svc user.Service, jwtSecret string) *Handler {
	return &Handler{svc: svc, jwtSecret: jwtSecret}
}

// -------------------- Send OTP --------------------

// SendOTP godoc
// @Summary      Send OTP
// @Description  Send OTP code to user phone
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body SendOTPRequest true "Send OTP request"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /auth/send-otp [post]
func (h *Handler) SendOTP(c *gin.Context) {
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.SendOTP(c.Request.Context(), req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// -------------------- Verify OTP --------------------

// VerifyOTP godoc
// @Summary      Verify OTP
// @Description  Verify OTP code and login/register user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body VerifyOTPRequest true "Verify OTP request"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/verify-otp [post]
func (h *Handler) VerifyOTP(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// first check OTP and get user
	u, err := h.svc.VerifyOTP(c.Request.Context(), req.Phone, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// JWT claims
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"role":    u.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // one day expire
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP verified successfully",
		"token":   signedToken,
	})
}
