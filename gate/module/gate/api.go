package gate




func echo(rpcIndex uint32, refer uint32, session *link.Session, msg proto.Message, gmt uint32) (proto.Message, uint32, error) {
	fmt.Println(rpcIndex)
	fmt.Println(refer)
	fmt.Println(session)
	fmt.Println(msg)
	fmt.Println(gmt)
	
	return nil, 0, nil
}
