package testtcp

import "encoding/binary"

// ReadEndian 读取消息长度
func ReadEndian(buf []byte, headerLen int, bigEndian bool) (msgLen uint) {
	if len(buf) < headerLen {
		return 0 // 消息长度不够
	}
	switch headerLen {
	case 1:
		msgLen = uint(buf[0])
	case 2:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint16(buf[:2]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint16(buf[:2]))
		}
	case 4:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint32(buf[:4]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint32(buf[:4]))
		}
	case 8:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint64(buf[:8]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint64(buf[:8]))
		}
	}
	return msgLen
}

// WriteEndian 消息长度写入头部字节
func WriteEndian(buf []byte, byteLen int, msgLen uint, bigEndian bool) {
	switch byteLen {
	case 1:
		buf[0] = byte(msgLen)
	case 2:
		if bigEndian {
			binary.BigEndian.PutUint16(buf[:2], uint16(msgLen))
		} else {
			binary.LittleEndian.PutUint16(buf[:2], uint16(msgLen))
		}
	case 4:
		if bigEndian {
			binary.BigEndian.PutUint32(buf[:4], uint32(msgLen))
		} else {
			binary.LittleEndian.PutUint32(buf[:4], uint32(msgLen))
		}
	case 8:
		if bigEndian {
			binary.BigEndian.PutUint64(buf[:8], uint64(msgLen))
		} else {
			binary.LittleEndian.PutUint64(buf[:8], uint64(msgLen))
		}
	}
}

// WrapHeader 包装消息头
func WrapHeader(msg []byte, headLen int) []byte {
	buf := make([]byte, len(msg)+headLen)
	WriteEndian(buf, headLen, uint(len(msg)), bigEndian)
	copy(buf[headLen:], msg)
	return buf
}

// WrapMsg 包装消息和消息编号
func WrapMsg(msg []byte, msgId uint, msgIdBytes int) []byte {
	buf := make([]byte, len(msg)+msgIdBytes)
	WriteEndian(buf, msgIdBytes, msgId, bigEndian)
	copy(buf[msgIdBytes:], msg)
	return buf
}
