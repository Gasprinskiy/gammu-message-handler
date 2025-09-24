package rest_api

import (
	"context"
	"net/http"
	"tgsms/config"
	"tgsms/internal/entity/global"
	"tgsms/internal/entity/incoming_messages"
	"tgsms/internal/transaction"
	"tgsms/tools/gin_gen"
	"tgsms/tools/logger"
	"tgsms/uimport"

	"github.com/gin-gonic/gin"
)

type IncomingMessagesHandler struct {
	ui     *uimport.Usecase
	router *gin.Engine
	config *config.Config
	log    *logger.Logger
	sm     transaction.SessionManager
}

func NewIncomingMessagesHandler(
	ui *uimport.Usecase,
	router *gin.Engine,
	config *config.Config,
	log *logger.Logger,
	sm transaction.SessionManager,
) {
	handler := &IncomingMessagesHandler{
		ui,
		router,
		config,
		log,
		sm,
	}

	routerGroup := router.Group("/inbox")

	{
		routerGroup.POST(
			"/on_recieve",
			handler.OnReceive,
		)
	}
}

func (h *IncomingMessagesHandler) OnReceive(gctx *gin.Context) {
	var params incoming_messages.OnMessageReceiveParams

	if err := gctx.ShouldBindJSON(&params); err != nil {
		gin_gen.HandleError(gctx, global.ErrInvalidParam)
		return
	}

	err := transaction.RunInTxExec(
		gctx,
		h.log,
		h.sm,
		func(ctx context.Context) error {
			return h.ui.IncomingMessages.OnMessageReceive(ctx, params)
		},
	)

	if err != nil {
		gin_gen.HandleError(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"success": true})
}
