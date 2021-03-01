package httpHost

import (
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

type Host struct {
	*gin.Engine
	id    string
	Port  int
	Peers map[string]p2pnetwork.PeerInfos
}

func NewHttpProvider(port int, id string) (*Host, error) {
	h := &Host{Engine: gin.Default(), id: id, Port: port, Peers: make(map[string]p2pnetwork.PeerInfos)}
	if err := h.applyRoutes(); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *Host) AddNewPeer(id, ip string, port int) {
	h.Peers[id] = p2pnetwork.PeerInfos{ID: id, IP: ip, Transport: "tcp", Port: port}
}

func (h *Host) AddPeer() func(c *gin.Context) {
	return func(c *gin.Context) {
		infos := p2pnetwork.PeerInfos{}
		err := c.BindJSON(&infos)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("id is empty"))
			return
		}
		h.Peers[infos.ID] = infos
		c.String(http.StatusOK, infos.ID)
	}
}

func (h *Host) ListPeers() func(c *gin.Context) {
	return func(c *gin.Context) {
		peers := make([]p2pnetwork.PeerInfos, len(h.Peers))
		i := 0
		for _, val := range h.Peers {
			peers[i] = val
			i += 1
		}
		c.JSON(http.StatusOK, peers)
	}
}

func (h *Host) RemovePeer() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("id is empty"))
			return
		}
		// TODO: Check ip Adddr
		delete(h.Peers, id)
		c.String(http.StatusOK, id)
	}
}

func (h *Host) applyRoutes() error {
	h.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	h.POST("/peers/add", h.AddPeer())
	h.POST("/peers/del/:id", h.RemovePeer())
	h.GET("/peers", h.ListPeers())
	return nil
}

func (h *Host) String() string {
	return h.id
}

// getClientIPByRequest tries to get directly from the Request.
// https://blog.golang.org/context/userip/userip.go
func getClientIPByRequestRemoteAddr(req *http.Request) (ip string, err error) {

	// Try via request
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Printf("debug: Getting req.RemoteAddr %v", err)
		return "", err
	} else {
		log.Printf("debug: With req.RemoteAddr found IP:%v; Port: %v", ip, port)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		message := fmt.Sprintf("debug: Parsing IP from Request.RemoteAddr got nothing.")
		log.Printf(message)
		return "", fmt.Errorf(message)

	}
	log.Printf("debug: Found IP: %v", userIP)
	return userIP.String(), nil

}
