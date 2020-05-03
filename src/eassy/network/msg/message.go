package msg

func MsgPack(protoId int,msg []byte) []byte {
	idBytes := caculateMsgIdBytes(protoId)
	msgLen := len(msg) + idBytes
	buffer := make([]byte,msgLen)
	offset :=0
	offset = packInId(protoId,buffer,offset)
	copy(buffer[offset:],msg)
	return buffer
}

func MsgUnpack(buffer []byte) (protoId int,msg []byte) {
	length :=len(buffer)
	var offset = 0
	protoId,offset = packOutId(buffer,0)
	msg = make([]byte,length-offset)
	copy(msg,buffer[offset:])
	return
}

func caculateMsgIdBytes(id int) int {
	l := 0
	for ;id>0; {
		l += 1
		id >>= 7
	}
	return l
}

func packInId(id int, buffer []byte, offset int) int {
	tmp := id % 128
	next := id / 128
	if next != 0 {
		tmp += 128
	}
	buffer[offset] = byte(tmp)
	offset++
	id = next
	for ;id != 0; {
		tmp = id % 128
		next = id / 128
		if next != 0 {
			tmp += 128
		}
		buffer[offset] = byte(tmp)
		offset++
		id = next
	}
	return offset
}

func packOutId(buffer []byte,offset int) (id int,offset2 int) {
	var i uint32 = 0
	m := int(buffer[offset])
	id += (m & 0x7f) << (7*i)
	offset++
	i++
	for ;m >= 128; {
		m = int(buffer[offset])
		id += (m & 0x7f) << (7*i)
		offset++
		i++
	}
	offset2 = offset
	return
}
