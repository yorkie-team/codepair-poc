package dto

import "time"

// FindUserResponse represents the user details.
type FindUserResponse struct {
	ID                string    `json:"id"`                          // ID of user (required)
	Nickname          string    `json:"nickname,omitempty"`          // Nickname of user
	LastWorkspaceSlug string    `json:"lastWorkspaceSlug,omitempty"` // Last workspace slug of user
	CreatedAt         time.Time `json:"createdAt"`                   // Created date of user (required)
	UpdatedAt         time.Time `json:"updatedAt"`                   // Updated date of user (required)
}

// ChangeNicknameDto holds the nickname update information.
type ChangeNicknameDto struct {
	Nickname string `json:"nickname"` // Nickname of user to update, minLength: 2 (required)
}

// LoginResponse provides tokens after login.
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`  // Access token for CodePair (required)
	RefreshToken string `json:"refreshToken"` // Refresh token to get a new access token (required)
}

// RefreshTokenRequestDto holds the refresh token for requesting new access token.
type RefreshTokenRequestDto struct {
	RefreshToken string `json:"refreshToken"` // The refresh token to request a new access token (required)
}

// RefreshTokenResponseDto carries the new access token information.
type RefreshTokenResponseDto struct {
	NewAccessToken string `json:"newAccessToken"` // The new access token (required)
}

// CreateWorkspaceDto for creating a new workspace.
type CreateWorkspaceDto struct {
	Title string `json:"title"` // Title of project to create, minLength: 2 (required)
}

// CreateWorkspaceResponse contains details of a newly created workspace.
type CreateWorkspaceResponse struct {
	ID        string    `json:"id"`        // ID of the workspace (required)
	Title     string    `json:"title"`     // Title of the workspace (required)
	Slug      string    `json:"slug"`      // Slug of the workspace (required)
	CreatedAt time.Time `json:"createdAt"` // Created date of the workspace (required)
	UpdatedAt time.Time `json:"updatedAt"` // Updated date of the workspace (required)
}

// FindWorkspaceResponse contains details of a found workspace.
type FindWorkspaceResponse struct {
	ID        string    `json:"id"`        // ID of the workspace (required)
	Title     string    `json:"title"`     // Title of the workspace (required)
	Slug      string    `json:"slug"`      // Slug of the workspace (required)
	CreatedAt time.Time `json:"createdAt"` // Created date of the workspace (required)
	UpdatedAt time.Time `json:"updatedAt"` // Updated date of the workspace (required)
}

// HttpExceptionResponse represents a standardized HTTP error.
type HttpExceptionResponse struct {
	StatusCode int    `json:"statusCode"` // Status Code of HTTP Response (required)
	Message    string `json:"message"`    // Description about the error (required)
}

// WorkspaceDomain represents the workspace domain model.
type WorkspaceDomain struct {
	ID        string    `json:"id"`        // ID of the workspace (required)
	Title     string    `json:"title"`     // Title of the workspace (required)
	Slug      string    `json:"slug"`      // Slug of the workspace (required)
	CreatedAt time.Time `json:"createdAt"` // Created date of the workspace (required)
	UpdatedAt time.Time `json:"updatedAt"` // Updated date of the workspace (required)
}

// FindWorkspacesResponse for listing workspaces with pagination cursor.
type FindWorkspacesResponse struct {
	Workspaces []WorkspaceDomain `json:"workspaces"` // List of found workspaces (required)
	Cursor     string            `json:"cursor"`     // The ID of last workspace (required)
}

// CreateInvitationTokenDto for creating an invitation token.
type CreateInvitationTokenDto struct {
	ExpiredAt time.Time `json:"expiredAt"` // Expiration date of invitation token (required)
}

// CreateInvitationTokenResponse provides the invitation token.
type CreateInvitationTokenResponse struct {
	InvitationToken string `json:"invitationToken"` // Token for invitation (required)
}

// JoinWorkspaceDto for joining a workspace using an invitation token.
type JoinWorkspaceDto struct {
	InvitationToken string `json:"invitationToken"` // Invitation token of workspace to join (required)
}

// JoinWorkspaceResponse contains details of the workspace after joining.
type JoinWorkspaceResponse struct {
	ID        string    `json:"id"`        // ID of the workspace (required)
	Title     string    `json:"title"`     // Title of the workspace (required)
	Slug      string    `json:"slug"`      // Slug of the workspace (required)
	CreatedAt time.Time `json:"createdAt"` // Created date of the workspace (required)
	UpdatedAt time.Time `json:"updatedAt"` // Updated date of the workspace (required)
}

// WorkspaceUserDomain represents a user within a workspace.
type WorkspaceUserDomain struct {
	ID        string    `json:"id"`        // ID of the user (required)
	Nickname  string    `json:"nickname"`  // Nickname of the user (required)
	CreatedAt time.Time `json:"createdAt"` // Created date of the user (required)
	UpdatedAt time.Time `json:"updatedAt"` // Updated date of the user (required)
}

// FindWorkspaceUsersResponse for listing users of a workspace with a pagination cursor.
type FindWorkspaceUsersResponse struct {
	WorkspaceUsers []WorkspaceUserDomain `json:"workspaceUsers"` // List of found workspace users (required)
	Cursor         string                `json:"cursor"`         // The ID of last workspace user (required)
}

// CreateWorkspaceDocumentDto for creating a new document in a workspace.
type CreateWorkspaceDocumentDto struct {
	Title string `json:"title"` // Title of document to create (required)
}

// CreateWorkspaceDocumentResponse provides details of the newly created document.
type CreateWorkspaceDocumentResponse struct {
	ID               string    `json:"id"`                // ID of the document (required)
	YorkieDocumentID string    `json:"yorkieDocumentId"`  // Yorkie Document ID of the document (required)
	Title            string    `json:"title"`             // Title of the document (required)
	Content          string    `json:"content,omitempty"` // Content of the document
	CreatedAt        time.Time `json:"createdAt"`         // Created date of the document (required)
	UpdatedAt        time.Time `json:"updatedAt"`         // Updated date of the document (required)
	WorkspaceID      string    `json:"workspaceId"`       // ID of the workspace that includes the document (required)
}

// CreateWorkspaceDocumentShareTokenDto for sharing a document.
type CreateWorkspaceDocumentShareTokenDto struct {
	Role      string    `json:"role"`      // Role to share (required, Enum: [READ, EDIT])
	ExpiredAt time.Time `json:"expiredAt"` // Share link expiration date (required)
}

// CreateWorkspaceDocumentShareTokenResponse returns a sharing token.
type CreateWorkspaceDocumentShareTokenResponse struct {
	SharingToken string `json:"sharingToken"` // Sharing token (required)
}

// FindDocumentFromSharingTokenResponse returns document details from a sharing token.
type FindDocumentFromSharingTokenResponse struct {
	ID               string    `json:"id"`                // ID of the document (required)
	YorkieDocumentID string    `json:"yorkieDocumentId"`  // Yorkie Document ID of the document (required)
	Title            string    `json:"title"`             // Title of the document (required)
	Content          string    `json:"content,omitempty"` // Content of the document
	CreatedAt        time.Time `json:"createdAt"`         // Created date of the document (required)
	UpdatedAt        time.Time `json:"updatedAt"`         // Updated date of the document (required)
	WorkspaceID      string    `json:"workspaceId"`       // ID of the workspace that includes the document (required)
	Role             string    `json:"role"`              // Role of share token (required, Enum: [READ, EDIT])
}

// CheckNameConflictDto used to check name conflicts.
type CheckNameConflictDto struct {
	Name string `json:"name"` // Name to check conflict (required)
}

// CheckNameConflictResponse provides result of name conflict check.
type CheckNameConflictResponse struct {
	Conflict bool `json:"conflict"` // Whether the name is conflict (required)
}

// CheckYorkieDto to check Yorkie auth.
type CheckYorkieDto struct {
	Token      string   `json:"token"`      // Token from client (required)
	Method     string   `json:"method"`     // Method of Yorkie to invoke (required, Enum: [ActivateClient, DeactivateClient, AttachDocument, DetachDocument, WatchDocuments, PushPull])
	Attributes []string `json:"attributes"` // Attribute to check auth
}

// CheckYorkieResponse provides the result of Yorkie auth check.
type CheckYorkieResponse struct {
	Allowed bool   `json:"allowed"` // Whether the given token is authorized for this document (required)
	Reason  string `json:"reason"`  // Reason for this response (required)
}

// RunFollowUpDto for running a follow-up feature.
type RunFollowUpDto struct {
	DocumentID string `json:"documentId"` // ID of document (required)
	MemoryKey  string `json:"memoryKey"`  // Key of chat history (required)
	Content    string `json:"content"`    // Content to run feature (required)
}

// RunFeatureDto for running a feature.
type RunFeatureDto struct {
	DocumentID string `json:"documentId"` // ID of document (required)
	Content    string `json:"content"`    // Content to run feature (required)
}

// CreateUploadPresignedUrlDto contains file upload request details.
type CreateUploadPresignedUrlDto struct {
	WorkspaceID   string `json:"workspaceId"`   // ID of workspace to create file (required)
	ContentLength int64  `json:"contentLength"` // Length of content to upload (required)
	ContentType   string `json:"contentType"`   // Type of file (required)
}

// CreateUploadPresignedUrlResponse gives presigned URL info.
type CreateUploadPresignedUrlResponse struct {
	URL     string `json:"url"`     // Presigned URL for upload (required)
	FileKey string `json:"fileKey"` // Key of file (required)
}

// ExportFileRequestBody contains file export request details.
type ExportFileRequestBody struct {
	ExportType string `json:"exportType"` // Export file type (required)
	Content    string `json:"content"`    // Markdown content (required)
	FileName   string `json:"fileName"`   // File name (required)
}

// Buffer is an empty placeholder.
type Buffer struct {
	// Intentionally left blank
}

// ExportFileResponse contains the file export response.
type ExportFileResponse struct {
	FileContent map[string]interface{} `json:"fileContent"` // File content (description: File content)
	MimeType    string                 `json:"mimeType"`    // File mime type (required)
	FileName    string                 `json:"fileName"`    // File name (required)
}

// YorkieIntelligenceConfig holds configuration for Yorkie Intelligence.
type YorkieIntelligenceConfig struct {
	Enable bool                   `json:"enable"` // Enable Yorkie Intelligence (required)
	Config map[string]interface{} `json:"config"` // Yorkie Intelligence Config (description: Yorkie Intelligence Config)
}

// FileUploadConfig holds configuration for file uploads.
type FileUploadConfig struct {
	Enable bool `json:"enable"` // Enable File Upload (required)
}

// GetSettingsResponse contains settings response details.
type GetSettingsResponse struct {
	YorkieIntelligence struct {
		Enable bool                   `json:"enable"` // Enable Yorkie Intelligence (required)
		Config map[string]interface{} `json:"config"` // Yorkie Intelligence Config (description: Yorkie Intelligence Config)
	} `json:"yorkieIntelligence"`
	FileUpload struct {
		Enable bool `json:"enable"` // Enable File Upload (required)
	} `json:"fileUpload"`
}
