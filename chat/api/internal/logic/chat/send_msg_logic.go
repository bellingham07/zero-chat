package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bellingham07/go-tool/errorx"
	"log"
	"strconv"
	"time"
	"zero-chat/chat/api/internal/common/imserver"
	"zero-chat/chat/api/internal/model"

	"zero-chat/chat/api/internal/svc"
	"zero-chat/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) error {
	userId := l.ctx.Value("userID").(string)

	r := imserver.SendMsgRequest{
		FromUid: userId,  // qq
		ToUid:   req.Uid, // 163
		// todo : lack of nickname field. i will add this field when i integrating RPC service
		Body:      req.Msg,
		TimeStamp: time.Now().Unix(),
	}
	rJson, err := json.Marshal(r)
	if err != nil {
		log.Printf("json marshal err:%s", err)
		return err
	}
	if err = l.svcCtx.Redis.Publish(l.ctx, "ws", rJson).Err(); err != nil {
		log.Printf("publish msg err:%s", err)
		return err
	}
	fmt.Println("msg sent")
	// store msg to db
	sId, _ := strconv.ParseInt(userId, 10, 64)
	rId, _ := strconv.ParseInt(req.Uid, 10, 64)
	message := &model.Message{
		SendId:    sId,
		ReceiveId: rId,
		Msg:       req.Msg,
	}
	if err = l.svcCtx.MessageModel.Insert(l.ctx, l.svcCtx.DB, message); err != nil {
		return errorx.Internal(err, "store msg error").Show()
	}
	return nil
}
