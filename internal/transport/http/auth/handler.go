/*
 * Copyright (c) 2025-11-20 shinoda4
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	authv1 "github.com/shinoda4/sd-grpc-proto/proto/auth/v1"
	"google.golang.org/grpc/metadata"
)

type Handler struct {
	client authv1.AuthServiceClient
}

func NewHandler(client authv1.AuthServiceClient) *Handler {
	return &Handler{client: client}
}

func ctxWithToken(c *gin.Context) context.Context {
	token := c.GetHeader("Authorization")
	if token != "" {
		md := metadata.Pairs("authorization", token)
		return metadata.NewOutgoingContext(c.Request.Context(), md)
	}
	return c.Request.Context()
}

func (h *Handler) Register(c *gin.Context) {
	var req authv1.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var req authv1.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	sendEmail := c.Query("send_email") == "true"

	req := &authv1.VerifyEmailRequest{
		Token:     token,
		SendEmail: sendEmail,
	}

	res, err := h.client.VerifyEmail(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) ValidateToken(c *gin.Context) {
	req := &authv1.ValidateTokenRequest{}
	res, err := h.client.ValidateToken(ctxWithToken(c), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req authv1.RefreshTokenRequest
	res, err := h.client.RefreshToken(ctxWithToken(c), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	var req authv1.LogoutRequest
	res, err := h.client.Logout(ctxWithToken(c), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Me(c *gin.Context) {
	var req authv1.MeRequest
	res, err := h.client.Me(ctxWithToken(c), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	var req authv1.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.ForgotPassword(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var req authv1.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.ResetPassword(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
