package network

import "syscall"
import "net"


const SO_ORIGINAL_DST = 80

//GetOriginalDst  获取 tcp nat前的地址
func GetOriginalDst(tcpConn *net.TCPConn) (addr net.TCPAddr, newTCPConn *net.TCPConn, err error) {
	newTCPConn = tcpConn
	// net.TCPConn.File() will cause the receiver's (clientConn) socket to be placed in blocking mode.
	// The workaround is to take the File returned by .File(), do getsockopt() to get the original
	// destination, then create a new *net.TCPConn by calling net.Conn.FileConn().  The new TCPConn
	// will be in non-blocking mode.  What a pain.
	connFile, err := tcpConn.File()
	if err != nil {
		return
	} else {
		tcpConn.Close()
	}

	// Get original destination
	// this is the only syscall in the Golang libs that I can find that returns 16 bytes
	// Example result: &{Multiaddr:[2 0 31 144 206 190 36 45 0 0 0 0 0 0 0 0] Interface:0}
	// port starts at the 3rd byte and is 2 bytes long (31 144 = port 8080)
	// IPv4 address starts at the 5th byte, 4 bytes long (206 190 36 45)
	mreq, err := syscall.GetsockoptIPv6Mreq(int(connFile.Fd()), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
	if err != nil {
		return
	}
	newConn, err := net.FileConn(connFile)
	if err != nil {
		return
	}

	newTCPConn = newConn.(*net.TCPConn)
	connFile.Close()

	addr.IP = mreq.Multiaddr[4:8]
	addr.Port = int(mreq.Multiaddr[2])<<8 + int(mreq.Multiaddr[3])

	return
}
