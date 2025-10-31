package routes

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/bot"
	"EverythingSuckz/fsb/internal/types"
	"EverythingSuckz/fsb/internal/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var linkLog *zap.Logger

func (e *allRoutes) LoadLink(r *Route) {
	linkLog = e.log.Named("Link")
	defer linkLog.Info("Loaded link route")
	r.Engine.GET("/api/link/:messageID", getLinkRoute)
}

func getLinkRoute(ctx *gin.Context) {
	messageIDParam := ctx.Param("messageID")
	messageID, err := strconv.Atoi(messageIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.LinkResponse{
			Ok:    false,
			Error: "Invalid message ID: " + err.Error(),
		})
		return
	}

	worker := bot.GetNextWorker()

	file, err := utils.FileFromMessage(ctx, worker.Client, messageID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.LinkResponse{
			Ok:    false,
			Error: "Failed to get file information: " + err.Error(),
		})
		return
	}

	hash := utils.PackFile(
		file.FileName,
		file.FileSize,
		file.MimeType,
		file.ID,
	)
	shortHash := utils.GetShortHash(hash)

	downloadLink := fmt.Sprintf("%s/stream/%d?hash=%s", config.ValueOf.Host, messageID, shortHash)

	ctx.JSON(http.StatusOK, types.LinkResponse{
		Ok:           true,
		MessageID:    messageID,
		FileName:     file.FileName,
		FileSize:     file.FileSize,
		MimeType:     file.MimeType,
		DownloadLink: downloadLink,
	})
}
