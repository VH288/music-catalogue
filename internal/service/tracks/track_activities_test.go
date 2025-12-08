package tracks

import (
	"context"
	"fmt"
	"testing"

	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_UpsertTrackActivites(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vtrue := true
	vfalse := false

	mockTrackActivityRepo := NewMocktrackActivitiesRepo(mockCtrl)
	type args struct {
		userID  uint
		request trackactivities.TrackActivityRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success: create",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &vtrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).
					Return(nil, gorm.ErrRecordNotFound)

				mockTrackActivityRepo.EXPECT().Create(gomock.Any(), trackactivities.TrackActivity{
					UserID:    args.userID,
					SpotifyID: args.request.SpotifyID,
					IsLiked:   args.request.IsLiked,
					CreatedBy: fmt.Sprintf("%d", args.userID),
					UpdatedBy: fmt.Sprintf("%d", args.userID),
				}).Return(nil)
			},
		},
		{
			name: "success: update",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &vtrue,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).
					Return(&trackactivities.TrackActivity{
						IsLiked: &vfalse,
					}, nil)

				mockTrackActivityRepo.EXPECT().Update(gomock.Any(), trackactivities.TrackActivity{
					IsLiked: args.request.IsLiked,
				}).Return(nil)
			},
		},
		{
			name: "error",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &vtrue,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).
					Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				trackActivitiesRepo: mockTrackActivityRepo,
			}
			if err := s.UpsertTrackActivites(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.UpsertTrackActivites() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
