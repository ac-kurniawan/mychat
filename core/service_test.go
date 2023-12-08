package core

import (
	"context"
	"errors"
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Service", func() {
	var (
		mockCtrl       *gomock.Controller
		repositoryMock *MockIRepository
		utilMock       *MockIUtil
		timeNow        = time.Now()
		service        IMychatService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repositoryMock = NewMockIRepository(mockCtrl)
		utilMock = NewMockIUtil(mockCtrl)
		utilMock.EXPECT().StartTrace(gomock.Any(), gomock.Any())
		utilMock.EXPECT().EndTrace(gomock.Any())
		service = NewMychatService(MychatService{
			Repository: repositoryMock,
			Util:       utilMock,
		})
	})

	AfterEach(func() {
		repositoryMock = nil
		utilMock = nil
		mockCtrl.Finish()
	})

	Context("GetRoomChatsByParticipant", func() {
		When("session status is exist and doesnt match with registered list", func() {
			It("should return error ErrSessionStatusInvalid", func() {
				utilMock.EXPECT().LogError(gomock.Any(), gomock.Any())
				utilMock.EXPECT().TraceError(gomock.Any(), gomock.Any())
				randomSessionId := "asd"
				result, err := service.GetRoomChatsByParticipant(context.Background(), "", &randomSessionId)
				Expect(result).To(BeNil())
				Expect(err).Error()
				Expect(err).To(MatchError(ErrSessionStatusInvalid))
			})
		})
		When("error in Repository.GetParticipantGroupsByParticipant", func() {
			It("should return error ErrFailedGetParticipantGroup", func() {
				utilMock.EXPECT().LogError(gomock.Any(), gomock.Any())
				utilMock.EXPECT().TraceError(gomock.Any(), gomock.Any())
				sessionId := "ACTIVE"
				repositoryMock.EXPECT().GetParticipantGroupsByParticipant(gomock.Any(), "participant").Return(nil, errors.New("error"))

				result, err := service.GetRoomChatsByParticipant(context.Background(), "participant", &sessionId)
				Expect(result).To(BeNil())
				Expect(err).Error()
				Expect(err).To(MatchError(ErrFailedGetParticipantGroup))
			})
		})
		When("error in Repository.GetRoomChatByParticipantGroups", func() {
			It("should return error ErrFailedGetRoomChat", func() {
				utilMock.EXPECT().LogError(gomock.Any(), gomock.Any())
				utilMock.EXPECT().TraceError(gomock.Any(), gomock.Any())

				participantList := []model.ParticipantGroupModel{
					{
						ParticipantGroup: "group1",
						Participant:      "1",
					},
					{
						ParticipantGroup: "group2",
						Participant:      "1",
					},
				}
				repositoryMock.EXPECT().GetParticipantGroupsByParticipant(gomock.Any(), "participant").Return(participantList, nil)
				repositoryMock.EXPECT().GetRoomChatByParticipantGroups(gomock.Any(), []string{"group1", "group2"}, nil, uint(10)).Return(nil, errors.New("error"))
				result, err := service.GetRoomChatsByParticipant(context.Background(), "participant", nil)

				Expect(result).To(BeNil())
				Expect(err).Error()
				Expect(err).To(MatchError(ErrFailedGetRoomChat))
			})
		})
		When("success get room chat", func() {
			It("should return list of room chat", func() {
				participantList := []model.ParticipantGroupModel{
					{
						ParticipantGroup: "group1",
						Participant:      "1",
					},
					{
						ParticipantGroup: "group2",
						Participant:      "1",
					},
				}
				repositoryMock.EXPECT().GetParticipantGroupsByParticipant(gomock.Any(), "participant").Return(participantList, nil)
				roomChatList := []model.RoomChatModel{
					{
						ParticipantGroup: "group1",
						CreatedAt:        timeNow,
					},
					{
						ParticipantGroup: "group2",
						CreatedAt:        timeNow,
					},
				}
				repositoryMock.EXPECT().GetRoomChatByParticipantGroups(gomock.Any(), []string{"group1", "group2"}, nil, uint(10)).Return(roomChatList, nil)
				result, err := service.GetRoomChatsByParticipant(context.Background(), "participant", nil)

				Expect(result).To(Equal(roomChatList))
				Expect(err).To(BeNil())
			})
		})
	})

	Context("GetRoomChatBySessionId", func() {
		When("session status is exist and doesnt match with registered list", func() {
			It("should return error ErrSessionStatusInvalid", func() {
				utilMock.EXPECT().LogError(gomock.Any(), gomock.Any())
				utilMock.EXPECT().TraceError(gomock.Any(), gomock.Any())
				randomSessionId := "asd"
				result, err := service.GetRoomChatBySessionId(context.Background(), "", &randomSessionId)
				Expect(result).To(BeNil())
				Expect(err).Error()
				Expect(err).To(MatchError(ErrSessionStatusInvalid))
			})
		})
		When("Repository.GetRoomChatBySessionId return error", func() {
			It("should return error ErrFailedGetRoomChat", func() {
				utilMock.EXPECT().LogError(gomock.Any(), gomock.Any())
				utilMock.EXPECT().TraceError(gomock.Any(), gomock.Any())
				repositoryMock.EXPECT().GetRoomChatBySessionId(gomock.Any(), "id", nil, uint(10)).Return(nil, errors.New("error"))
				result, err := service.GetRoomChatBySessionId(context.Background(), "id", nil)
				Expect(result).To(BeNil())
				Expect(err).Error()
				Expect(err).To(MatchError(ErrFailedGetRoomChat))
			})
		})
		When("success get room chat", func() {
			It("should return list of room chat", func() {
				repositoryMock.EXPECT().GetRoomChatBySessionId(gomock.Any(), "id", nil, uint(10)).Return(&model.RoomChatModel{
					ParticipantGroup: "group1",
					CreatedAt:        timeNow,
				}, nil)
				result, err := service.GetRoomChatBySessionId(context.Background(), "id", nil)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(&model.RoomChatModel{
					ParticipantGroup: "group1",
					CreatedAt:        timeNow,
				}))
			})
		})
	})
})
