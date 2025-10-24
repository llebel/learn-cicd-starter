package main

import (
	"testing"
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

func TestDatabaseUserToUser(t *testing.T) {
	tests := []struct {
		name      string
		dbUser    database.User
		wantError bool
	}{
		{
			name: "Valid user conversion",
			dbUser: database.User{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				CreatedAt: "2023-10-24T12:00:00Z",
				UpdatedAt: "2023-10-24T12:00:00Z",
				Name:      "John Doe",
				ApiKey:    "test-api-key",
			},
			wantError: false,
		},
		{
			name: "Invalid createdAt format",
			dbUser: database.User{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				CreatedAt: "invalid-date",
				UpdatedAt: "2023-10-24T12:00:00Z",
				Name:      "John Doe",
				ApiKey:    "test-api-key",
			},
			wantError: true,
		},
		{
			name: "Invalid updatedAt format",
			dbUser: database.User{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				CreatedAt: "2023-10-24T12:00:00Z",
				UpdatedAt: "invalid-date",
				Name:      "John Doe",
				ApiKey:    "test-api-key",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := databaseUserToUser(tt.dbUser)
			if (err != nil) != tt.wantError {
				t.Errorf("databaseUserToUser() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError {
				if user.ID != tt.dbUser.ID {
					t.Errorf("Expected ID %v, got %v", tt.dbUser.ID, user.ID)
				}
				if user.Name != tt.dbUser.Name {
					t.Errorf("Expected Name %v, got %v", tt.dbUser.Name, user.Name)
				}
				if user.ApiKey != tt.dbUser.ApiKey {
					t.Errorf("Expected ApiKey %v, got %v", tt.dbUser.ApiKey, user.ApiKey)
				}
			}
		})
	}
}

func TestDatabaseNoteToNote(t *testing.T) {
	tests := []struct {
		name      string
		dbNote    database.Note
		wantError bool
	}{
		{
			name: "Valid note conversion",
			dbNote: database.Note{
				ID:        "note-123",
				CreatedAt: "2023-10-24T12:00:00Z",
				UpdatedAt: "2023-10-24T12:00:00Z",
				Note:      "Test note content",
				UserID:    "user-456",
			},
			wantError: false,
		},
		{
			name: "Invalid createdAt format",
			dbNote: database.Note{
				ID:        "note-123",
				CreatedAt: "invalid-date",
				UpdatedAt: "2023-10-24T12:00:00Z",
				Note:      "Test note content",
				UserID:    "user-456",
			},
			wantError: true,
		},
		{
			name: "Invalid updatedAt format",
			dbNote: database.Note{
				ID:        "note-123",
				CreatedAt: "2023-10-24T12:00:00Z",
				UpdatedAt: "invalid-date",
				Note:      "Test note content",
				UserID:    "user-456",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := databaseNoteToNote(tt.dbNote)
			if (err != nil) != tt.wantError {
				t.Errorf("databaseNoteToNote() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError {
				if note.ID != tt.dbNote.ID {
					t.Errorf("Expected ID %v, got %v", tt.dbNote.ID, note.ID)
				}
				if note.Note != tt.dbNote.Note {
					t.Errorf("Expected Note %v, got %v", tt.dbNote.Note, note.Note)
				}
				if note.UserID != tt.dbNote.UserID {
					t.Errorf("Expected UserID %v, got %v", tt.dbNote.UserID, note.UserID)
				}
			}
		})
	}
}

func TestDatabasePostsToPosts(t *testing.T) {
	validTime := time.Now().UTC().Format(time.RFC3339)

	tests := []struct {
		name      string
		dbNotes   []database.Note
		wantLen   int
		wantError bool
	}{
		{
			name: "Convert multiple notes",
			dbNotes: []database.Note{
				{
					ID:        "note-1",
					CreatedAt: validTime,
					UpdatedAt: validTime,
					Note:      "First note",
					UserID:    "user-1",
				},
				{
					ID:        "note-2",
					CreatedAt: validTime,
					UpdatedAt: validTime,
					Note:      "Second note",
					UserID:    "user-1",
				},
			},
			wantLen:   2,
			wantError: false,
		},
		{
			name:      "Empty slice",
			dbNotes:   []database.Note{},
			wantLen:   0,
			wantError: false,
		},
		{
			name: "Invalid note in slice",
			dbNotes: []database.Note{
				{
					ID:        "note-1",
					CreatedAt: validTime,
					UpdatedAt: validTime,
					Note:      "Valid note",
					UserID:    "user-1",
				},
				{
					ID:        "note-2",
					CreatedAt: "invalid-date",
					UpdatedAt: validTime,
					Note:      "Invalid note",
					UserID:    "user-1",
				},
			},
			wantLen:   0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notes, err := databasePostsToPosts(tt.dbNotes)
			if (err != nil) != tt.wantError {
				t.Errorf("databasePostsToPosts() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError && len(notes) != tt.wantLen {
				t.Errorf("Expected %d notes, got %d", tt.wantLen, len(notes))
			}
		})
	}
}
