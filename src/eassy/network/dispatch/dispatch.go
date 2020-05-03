package dispatch

func Dispatch(protoId int,msg []byte)  {
	protoPrefix := protoId/100
	switch protoPrefix {
	case 10:
		//login
	default:

	}
}
