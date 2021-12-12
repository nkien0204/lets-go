package rbitmq

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"

	"os"
	"path/filepath"
	"github.com/nkien0204/projectTemplate/log"
	"time"

	"go.uber.org/zap"
)

const BreakMessChar string = "\n"

type RabbitBackupHandler struct {
	fileName     string
	file         *os.File
	index        int
	backupStatus bool
}

type BackupObj struct {
}

func NewRabbitBackupHandler() (handler *RabbitBackupHandler) {
	logger := log.Logger()
	handler = &RabbitBackupHandler{
		file:         nil,
		index:        0,     // total of messages in the backup file
		backupStatus: false, // true: need to send backup data to controller first, false: do nothing
		fileName:     fmt.Sprintf("%s/%s", BackupFolder, BackupFileName),
	}
	os.Mkdir(BackupFolder, 0775)
	info, err := os.Stat(BackupFolder + "/" + BackupFileName)
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

func (backupFile *RabbitBackupHandler) writeToBackupFile(message amqp.Publishing) (err error) {
	logger := log.Logger()

	backupFile.file, err = os.OpenFile(filepath.Join(BackupFolder, filepath.Base(backupFile.fileName)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("could not open file", zap.String("name", backupFile.fileName), zap.Error(err))
		return err
	}
	defer backupFile.file.Close()
	saveData, err := transformData(message.Body)
	if err != nil {
		logger.Error("error while transform data")
		return err
	}
	saveData += BreakMessChar
	if _, err = backupFile.file.WriteString(saveData); err != nil {
		logger.Error("could not write to file", zap.String("name", backupFile.fileName), zap.Error(err))
		return err
	}
	backupFile.index++

	return nil
}

func (backupFile *RabbitBackupHandler) readFromBackupFile(scanner *bufio.Scanner) (message amqp.Publishing, err error) {
	logger := log.Logger()

	message, err = pushToMessages(scanner.Text())
	if err != nil {
		logger.Error("error while pushing to messages", zap.Error(err))
		return
	}
	return
}

func pushToMessages(content string) (message amqp.Publishing, err error) {
	logger := log.Logger()
	var fileStruct BackupObj
	err = json.Unmarshal([]byte(content), &fileStruct)
	if err != nil {
		logger.Error("unmarshal error", zap.Error(err))
		return
	}
	protMess := vms.NVRInternalMessageEvent{
		// fill data here
	}

	body, err := proto.Marshal(&protMess)
	if err != nil {
		log.Logger().Info("Error while marshal protobuf")
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
	logger := log.Logger()
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

func convProto2JsonMess(protoMess *vms.NVRInternalMessageEvent) (fileStruct *BackupObj) {
	mess := protoMess.GetFileCompletedEv()
	fileStruct = &BackupObj{
		FileName:   mess.GetFileName(),
		FileSize:   mess.GetSize(),
		FileTime:   int64(mess.GetTimeLen()),
		Codec:      mess.GetCodec(),
		CamId:      mess.GetCameraId(),
		StartTime:  mess.GetStartTime(),
		EndTime:    mess.GetEndTime(),
		FilePath:   mess.GetPath(),
		Domain:     mess.GetDomain(),
		ZoneUuid:   mess.GetZoneUuid(),
		CameraUuid: mess.GetCameraUuid(),
		NvrId:      mess.GetNvrId(),
		FileType:   mess.GetFileType(),
		NginxHost:  mess.GetNginxConf(),
		Img:        mess.GetImg(),
		CameraName: mess.GetCameraName(),
		DiskId:     mess.GetDiskId(),
	}
	return
}

func convRaw2ProtoMess(rawData []byte) (protoMess vms.NVRInternalMessageEvent, err error) {
	logger := log.Logger()
	if err = proto.Unmarshal(rawData, &protoMess); err != nil {
		logger.Error("unmarshal err", zap.Error(err))
		return protoMess, err
	}
	return
}
