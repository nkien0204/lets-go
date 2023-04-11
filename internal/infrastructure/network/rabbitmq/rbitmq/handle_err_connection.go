package rbitmq

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"google.golang.org/protobuf/proto"

	"github.com/streadway/amqp"

	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

const BreakMessChar string = "\n"

type RabbitBackupHandler struct {
	fileName     string
	file         *os.File
	index        int
	backupStatus bool
	cfg          configs.RabbitConfig
}

type BackupObj struct {
}

func NewRabbitBackupHandler(cfg *configs.Cfg) (handler *RabbitBackupHandler) {
	logger := rolling.New()
	handler = &RabbitBackupHandler{
		file:         nil,
		index:        0,     // total of messages in the backup file
		backupStatus: false, // true: need to send backup data to controller first, false: do nothing
		fileName:     fmt.Sprintf("%s/%s", cfg.Rabbit.BackupFolder, cfg.Rabbit.BackupFileName),
		cfg:          cfg.Rabbit,
	}
	os.Mkdir(cfg.Rabbit.BackupFolder, 0775)
	info, err := os.Stat(cfg.Rabbit.BackupFolder + "/" + cfg.Rabbit.BackupFileName)
	if err != nil {
		logger.Warn("Error while getting file size", zap.Error(err))
		return
	}
	if info.Size() <= 0 {
		logger.Warn("file is empty")
		return
	}
	handler.backupStatus = true
	return
}

func (handler *RabbitBackupHandler) writeToBackupFile(message amqp.Publishing) (err error) {
	logger := rolling.New()

	handler.file, err = os.OpenFile(filepath.Join(handler.cfg.BackupFolder, filepath.Base(handler.fileName)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("could not open file", zap.String("name", handler.fileName), zap.Error(err))
		return err
	}
	defer handler.file.Close()
	saveData, err := transformData(message.Body)
	if err != nil {
		logger.Error("error while transform data")
		return err
	}
	saveData += BreakMessChar
	if _, err = handler.file.WriteString(saveData); err != nil {
		logger.Error("could not write to file", zap.String("name", handler.fileName), zap.Error(err))
		return err
	}
	handler.index++

	return nil
}

func (handler *RabbitBackupHandler) readFromBackupFile(scanner *bufio.Scanner) (message amqp.Publishing, err error) {
	logger := rolling.New()

	message, err = pushToMessages(scanner.Text())
	if err != nil {
		logger.Error("error while pushing to messages", zap.Error(err))
		return
	}
	return
}

func pushToMessages(content string) (message amqp.Publishing, err error) {
	logger := rolling.New()
	var fileStruct BackupObj
	err = json.Unmarshal([]byte(content), &fileStruct)
	if err != nil {
		logger.Error("unmarshal error", zap.Error(err))
		return
	}
	protMess := events.InternalMessageEvent{
		// fill data here
	}

	body, err := proto.Marshal(&protMess)
	if err != nil {
		rolling.New().Info("Error while marshal protobuf")
		return
	}
	message = amqp.Publishing{
		Headers:         nil,
		ContentType:     "",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            body,
	}
	return message, nil
}

func transformData(rawData []byte) (fileStruct string, err error) {
	logger := rolling.New()
	var fileStructByte []byte
	protoMess, err := convRaw2ProtoMess(rawData)
	if err != nil {
		logger.Error("error while converting data", zap.Error(err))
		return string(fileStruct), err
	}
	fileStructByte, err = json.Marshal(convProto2JsonMess(&protoMess))
	if err != nil {
		logger.Error("error while marshaling", zap.Error(err))
		return
	}

	return string(fileStructByte), nil
}

func convProto2JsonMess(protoMess *events.InternalMessageEvent) (fileStruct *BackupObj) {
	// mess := protoMess.GetFileCompletedEv()
	fileStruct = &BackupObj{}
	return
}

func convRaw2ProtoMess(rawData []byte) (protoMess events.InternalMessageEvent, err error) {
	logger := rolling.New()
	if err = proto.Unmarshal(rawData, &protoMess); err != nil {
		logger.Error("unmarshal err", zap.Error(err))
		return protoMess, err
	}
	return
}
